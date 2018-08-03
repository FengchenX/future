package manager

import (
	"bytes"
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"sub_account_service/order_server_zhengdao/client"
	"sub_account_service/order_server_zhengdao/db"
	"sub_account_service/order_server_zhengdao/lib"
	"sub_account_service/order_server_zhengdao/models"
)

func InitAutoTB() {
	if db.AutoMigrate == true {
		glog.Infoln("init AutoMigrate mysql db tables")
		db.DbClient.Client.AutoMigrate(&models.Order{})
		db.DbClient.Client.AutoMigrate(&models.OrderSave{})
		db.DbClient.Client.AutoMigrate(&models.Bill{})
	}
}

// save orders form Thirty-party api
func SaveOrder(c *gin.Context, cli *client.Client) {
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

	glog.Infoln("SaveOrder********************", order)
	glog.Infoln(lib.Loger("[three] SaveOrder", "order_msg"), order.Company)

	// check md5
	md5Token := lib.CipherStr(models.ThreeKey, strconv.Itoa(int(order.Time)))
	if strings.ToLower(md5Token) != strings.ToLower(order.Token) {
		glog.Errorln("SaveOrder err:", order.Token, md5Token)
		c.JSON(http.StatusOK, gin.H{"server": "error Token Error"})
		return
	}

	// find company by CompanyName and BanchName
	company := &models.Company{}
	dberr := db.DbClient.Client.Where(&models.Company{
		CompanyBanchName: string(order.BranchShop),
	}).First(company)

	// if don't have branch_company try main company
	if dberr.Error != nil {
		glog.Errorln("haven't register BranchShop!")
		c.JSON(http.StatusOK, gin.H{"success": "haven't register BranchShop!"})
		return
	}

	glog.Infoln(lib.Loger("[three] SaveOrder", "order_msg"), order.OrderId, company)

	// get job address throuth block chain
	/*	respj, err := cli.C.GetNowJobAddress(context.Background(), &protocol.ReqGetNowJobAddr{
			CompanyName:    company.CompanyName,
			SubCompanyName: company.CompanyBanchName,
			TimeStamp:      time.Now().String(),
		})

		glog.Infoln(lib.Log("three", respj.JobAddr, "job addr is"))

		b, err := json.Marshal(order.OrderContent)
		str := bytes.NewBuffer(b).String()
		reqg := &protocol.ReqThreeSetOrder{
			UserAddress:     company.CompanyAddress,
			PassWord:        company.CompanyPassWord,
			AccountDescribe: company.CompanyDespcrite,
			Money:           float64(order.Price),
			OrderId:         strconv.Itoa(int(order.OrderId)),
			JobAddress:      respj.GetJobAddr(),
			Content: &protocol.Order{
				Table:      uint32(order.OrderId),
				TimeStamp:  lib.SetTime(order.CreateTime).Unix(),
				Money:      float64(order.Price),
				Content:    str,
				JobAddress: respj.GetJobAddr(),
				OrderId:    strconv.Itoa(int(order.OrderId)),
			},
		}
	*/
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

	c.JSON(http.StatusOK, gin.H{"server": "success", "echo": order})
	// sending msg to block chain
	//	cli.C.ThreeSetOrder(context.Background(), reqg)
	//	glog.Infoln(lib.Log("three", respj.JobAddr, "order req"), reqg)
}
