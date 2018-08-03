package api

import (
	"github.com/golang/glog"
	"sub_account_service/app_server_v2/balance"
	"sub_account_service/app_server_v2/db"
	"sub_account_service/app_server_v2/lib"
	"github.com/gin-gonic/gin"
	"net/http"
	"sub_account_service/app_server_v2/model"
	"time"
	"sub_account_service/app_server_v2/config"
)

var jobs = make(chan db.NumberOrder, 100)

//SaveOrders 存取订单
func SaveOrders() {
	tk := time.NewTicker(60 * time.Second)

	get := func() {
		orders, err := balance.GetOrderList()
		if err != nil {
			glog.Errorln("SaveOrders***************get order err:", err)
			return
		}
		glog.Infoln("SaveOrders*****************orders:", orders)
		for _, order := range orders {
			if db.NotFindOrder(order.ThirdTradeNo, order.SubAccountNo) {
				//没有找到
				jobs <- order
			}
		}
	}
	go func() {
		get()
		for {
			select {
			case <-tk.C:
				get()
			default:
				time.Sleep(1*time.Second)
			}
		}
	}()

	go writeOrder()
	go getBook()

	//todo 测试删除
	//go getOrder()
}

func writeOrder() {
	defer func() {
		if jobs != nil {
			close(jobs)
		}
	}()
	DB := db.DbClient.Client
	if DB == nil {
		glog.Errorln("writeOrder*************db is nil")
	}
	for job := range jobs {
		if mydb := DB.Create(&job); mydb.Error != nil {
			glog.Errorln("writeOrder*************db err", mydb.Error)
			continue
		}
	}
}

func getBook() {
	timer := time.NewTicker(time.Duration(config.ConfInst().GetBookTime) * time.Second)
	DB := db.DbClient.Client
	if DB == nil {
		glog.Errorln("writeOrder*************db is nil")
	}
	for {
		select {
		case <-timer.C:
			var orders []db.NumberOrder
			if mydb := DB.Where("number_orders.read = ? AND dump_num < ?", false, config.ConfInst().DumpNum).Order("create_time").Order("id").Limit(config.ConfInst().OrderNum).Find(&orders); mydb.Error != nil {
				glog.Errorln("getBook*******************db find err", mydb.Error)
				continue
			}
			tx := DB.Begin()
			for _, order := range orders {
				var req model.ReqGetABById
				var resp model.RespGetABById
				req.ScheduleName = order.SubAccountNo
				req.OrderId = order.ThirdTradeNo
				glog.Infoln("getBook************req:", req)
				if err := doPost("/getabbyid", req, &resp); err != nil {
					glog.Errorln("SaveOrders*****************doPost", err)
				}
				glog.Infoln("getBook***************resp:", resp)
				if resp.StatusCode != 0 {
					if mydb := tx.Model(&order).Update("dump_num", order.DumpNum + 1); mydb.Error != nil {
						glog.Errorln("getBook***************************db err:", mydb.Error)
					}
					glog.Errorln("getBook****************", resp.Msg, "| subCode:", order.SubAccountNo)
					continue
				}
				
				transComplete := true
				for _, abs := range resp.Abs {
					glog.Infoln("getBook*************abs.Rflag ", abs.Rflag, abs.Acco.Name)
					if abs.Rflag == true {

						db.InsertOrder(order.SubAccountNo, order.ThirdTradeNo,
							abs.Address, abs.Acco.Alipay, abs.Acco.Name,
							abs.Money, abs.Radio, uint32(abs.SubWay), order.OrderTime)
					} else {
						transComplete = false
					}
					
				}
				if transComplete {
					glog.Infoln("getBook***********************************成功")
					if mydb := tx.Model(&order).Update("read", true); mydb.Error != nil {
						glog.Errorln("getBook********************db update err", mydb.Error)
						continue
					}
				}
			}
			tx.Commit()
		default:
			time.Sleep(1*time.Second)
		}
	}
}

//GetMoney 查询收益
func GetMoney(c *gin.Context) {
	var req model.ReqGetMoney
	var resp model.RespGetMoney
	if err := lib.ParseReq(c, "GetMoney", &req); err != nil {
		return
	}
	glog.Infoln("GetMoney*****************req: ", req)

	all, date, month := db.GetAllMoney(req.UserAddress)
	b := time.Unix(req.StartTime, 0)
	e := time.Unix(req.EndTime, 0)
	if req.StartTime == 0 && req.EndTime == 0 {
		e = time.Now()
	}

	if (req.StartTime == 0 && req.EndTime != 0) || (req.EndTime == 0 && req.StartTime != 0) {
		resp.StatusCode = M_TIME_ERR
		resp.Msg = "Fail"
		glog.Infoln("GetMoney*********************resp: ", resp)
		c.JSON(http.StatusOK, resp)
		return
	}

	apporder, page := db.GetMoney(req.UserAddress, b, e, int(req.Page))

	reb := make([]model.Bill, 0)
	for _, v := range apporder {
		bill := model.Bill{
			Name: v.Name, 
			Money: lib.MoneyP(v.Price),
			Ratio: v.Ratio,
			SubWay: v.SubWay,
			PayAcco: v.PayAccount,
		}
		reb = append(reb, bill)
	}

	if len(reb) == 0 {
		resp.StatusCode = M_NO_HAVE_MORE
		resp.Msg = "没有更多"
		glog.Infoln("GetMoney*********************resp: ", resp)
		c.JSON(http.StatusOK, resp)
		return
	}

	resp.StatusCode = M_SUCCESS
	resp.AllMoney = lib.MoneyP(all)
	resp.Date = lib.MoneyP(date)
	resp.Month = lib.MoneyP(month)
	resp.Bills = reb
	resp.PageCount = uint32(page)
	resp.Msg = "Success"
	glog.Infoln("GetMoney*********************resp: ", resp)
	c.JSON(http.StatusOK, resp)
}

//todo 测试删除
func getOrder() {
	glog.Infoln("GetOrder****************start")
	db.InitOrderDB()
	var pages int
	pages = 1
	tk := time.NewTicker(5 * time.Second)
	for {
		select {
		case <- tk.C:
		nos, p := db.GetOrder2(pages)
		pages = p
		glog.Infoln("GetOrder**************orders:", nos)
		if len(nos) == db.Size {
			pages++
			for _, no := range nos {
				db.InsertOrder(no.SubAccountNo, no.TradeNo, no.Address, no.Alipay, no.Name, no.Money, no.Radio, uint32(no.SubWay), no.CreateTime)
			}
		} else if len(nos) > 0 {
			for _, no := range nos {
				db.InsertOrder(no.SubAccountNo, no.TradeNo, no.Address, no.Alipay, no.Name, no.Money, no.Radio, uint32(no.SubWay), no.CreateTime)
			}
			glog.Infoln("getOrder***********************no more")
		}
		default:
			time.Sleep(1 * time.Second)
		}
	}
}



