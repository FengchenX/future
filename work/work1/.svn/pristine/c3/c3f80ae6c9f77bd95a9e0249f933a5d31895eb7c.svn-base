package manager

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"io/ioutil"
	"net/http"
	"time"
	"znfz/web_server/client"
	"znfz/web_server/db"
	"znfz/web_server/lib"
	"znfz/web_server/models"
	"znfz/web_server/protocol"
)

func InitAutoTB() {
	if db.AutoMigrate == true {
		glog.Infoln("init AutoMigrate mysql db tables")
		db.DbClient.Client.AutoMigrate(&models.Order{})
		db.DbClient.Client.AutoMigrate(&models.OrderSave{})
		db.DbClient.Client.AutoMigrate(&models.Bill{})
		db.DbClient.Client.AutoMigrate(&models.Binder{})
	}
}

// save orders form Thirty-party api
func SaveOrder(c *gin.Context, cli *client.Client) {
	t1 := time.Now()
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Errorln("SaveOrder err:", err)
		c.JSON(http.StatusOK, gin.H{"server": "error json unmarshall fall"})
		return
	}
	order := &models.Order{}
	err = json.Unmarshal(body, order)
	if err != nil {
		glog.Errorln("SaveOrder err:", err, string(body))
		c.JSON(http.StatusOK, gin.H{"server": "error json unmarshall fall"})
		return
	}

	glog.Infoln(lib.Loger("SaveOrder", "order_msg"), order)

	// find company by CompanyName and BanchName
	company := &models.Company{}
	dberr := db.DbClient.Client.Where(&models.Company{
		CompanyName: string(order.Company),
	}).First(company)

	// if don't have branch_company try main company
	if dberr.Error != nil {
		glog.Errorln(dberr.Error)
		if dberr.Error != nil {
			glog.Errorln("haven't register company!")
			c.JSON(http.StatusOK, gin.H{"server": "haven't register company"})
			return
		}
		dberr2 := db.DbClient.Client.Where(&models.Binder{
			CompanyName: string(order.BranchShop),
		}).First(company)
		if dberr2.Error != nil {
			glog.Errorln("haven't register company!")
			c.JSON(http.StatusOK, gin.H{"success": "haven't register company!"})
			return
		}
	}
	// get job address throuth block chain
	respj, err := cli.C.GetNowJobAddress(context.Background(), &protocol.ReqGetNowJobAddr{
		CompanyName:    company.CompanyName,
		SubCompanyName: company.CompanyBanchName,
		TimeStamp:      time.Now().String(),
	})

	b, err := json.Marshal(order.OrderContent)
	str := bytes.NewBuffer(b).String()
	reqg := &protocol.ReqThreeSetOrder{
		UserAddress:     company.CompanyAddress,
		PassWord:        company.CompanyPassWord,
		AccountDescribe: company.CompanyDespcrite,
		Money:           float64(order.Price),
		Content: &protocol.Order{
			Table:     uint32(order.OrderId),
			TimeStamp: lib.SetTime(order.CreateTime).Unix(),
			Money:     float64(order.Price),
			Content:   str,
			JobAddress:respj.GetJobAddr(),
		},
		JobAddress:respj.GetJobAddr(),
	}

	cont, _ := json.Marshal(order.OrderContent)
	sontStr := bytes.NewBuffer(cont).String()

	// save orders to mysql
	err = db.DbClient.Client.Create(&models.OrderSave{
		OrderId:      order.OrderId,
		CreateTime:   order.CreateTime,
		Company:      order.Company,
		BranchShop:   order.BranchShop,
		Price:        order.Price,
		OrderContent: sontStr,
	}).Error

	if err != nil {
		glog.Errorln("SaveOrder create db err:", err, string(body))
		c.JSON(http.StatusOK, gin.H{"server": "data illegal"})
		return
	}

	// 存证上链太慢了，800ms+，简易先回包
	c.JSON(http.StatusOK, gin.H{"server": "success", "echo": order})
	// sending msg to block chain
	cli.C.ThreeSetOrder(context.Background(), reqg)
	t7 := time.Now()
	glog.Infoln("其实慢的原因是公私钥加密算法太慢，可以不要的", t7.Sub(t1))
}

// save orders form Thirty-party api
func SaveBill(c *gin.Context, cli *client.Client) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Errorln("SaveOrder err:", err)
		c.JSON(http.StatusOK, gin.H{"server": "error json unmarshall fall"})
		return
	}
	bill := &models.Bill{}
	err = json.Unmarshal(body, bill)
	if err != nil {
		glog.Errorln("SaveOrder err:", err, string(body))
		c.JSON(http.StatusOK, gin.H{"server": "error json unmarshall fall"})
		return
	}

	glog.Infoln(lib.Loger("SaveOrder", "SaveBill"), bill)

	// find company by CompanyName and BanchName
	company := &models.Company{}
	dberr := db.DbClient.Client.Where(&models.Company{
		CompanyName: bill.Company,
	}).First(company)

	// if don't have branch_company try main company
	if dberr.Error != nil {
		glog.Errorln(dberr.Error)
		dberr2 := db.DbClient.Client.Where(&models.Company{
			CompanyName: bill.BranchShop,
		}).First(company)
		if dberr2.Error != nil {
			glog.Errorln("haven't register company!")
			c.JSON(http.StatusOK, gin.H{"server": "haven't register company"})
			return
		}
	}

	glog.Infoln("SaveBill   ", company.CompanyName, company.CompanyBanchName)

	// get job address throuth block chain
	respj, err := cli.C.GetNowJobAddress(context.Background(), &protocol.ReqGetNowJobAddr{
		CompanyName:    company.CompanyName,
		SubCompanyName: company.CompanyBanchName,
		TimeStamp:      time.Now().String(),
	})

	glog.Infoln("respj   ", respj.GetJobAddr(), time.Now().String())

	reqg := &protocol.ReqThreeSetBill{
		UserAddress:     company.CompanyAddress,
		PassWord:        company.CompanyPassWord,
		AccountDescribe: company.CompanyDespcrite,
		Money:           float64(bill.Price),
		JobAddress:      respj.GetJobAddr(),
		CompanyName:     bill.Company,
		SubCompanyName:  bill.BranchShop,
	}
	// save orders to mysql
	db.DbClient.Client.Create(bill)
	c.JSON(http.StatusOK, gin.H{"server": "success", "echo": bill})
	// sending msg to block chain
	cli.C.ThreeSetBill(context.Background(), reqg)
}
