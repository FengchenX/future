package manager

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"github.com/satori/go.uuid"
	"io"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"sub_account_service/order_server_zhengdao/client"
	"sub_account_service/order_server_zhengdao/config"
	"sub_account_service/order_server_zhengdao/db"
	"sub_account_service/order_server_zhengdao/lib"
	"sub_account_service/order_server_zhengdao/models"
	"sub_account_service/order_server_zhengdao/payment"
	"sub_account_service/order_server_zhengdao/payment/alipayModels"
)

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

	if dberr := db.DbClient.Client.Where(&models.Bill{OrderId: bill.OrderId}).First(bill); dberr.Error == nil {
		c.JSON(http.StatusOK, gin.H{"server": "bill is exist", "echo": bill})
		return
	}

	md5Token := lib.CipherStr(models.ThreeKey, strconv.Itoa(int(bill.Time)))
	if strings.ToLower(md5Token) != strings.ToLower(bill.Token) {
		glog.Errorln("SaveOrder err:", bill.Token, md5Token)
		c.JSON(http.StatusOK, gin.H{"server": "error Token Error"})
		return
	}

	glog.Infoln(lib.Loger("[three] SaveOrder", "SaveBill"), bill)

	// find company by CompanyName and BanchName
	company := &models.Company{}
	dberr := db.DbClient.Client.Where(&models.Company{
		CompanyBanchName: bill.BranchShop,
	}).First(company)

	// if don't have branch_company try main company
	if dberr.Error != nil {
		glog.Errorln("haven't register company!")
		c.JSON(http.StatusOK, gin.H{"server": "haven't register company"})
		return
	}

	glog.Infoln("SaveBill   ", company.CompanyName, company.CompanyBanchName)

	/*	reqg := &protocol.ReqThreeSetBill{
		UserAddress:     company.CompanyAddress,
		PassWord:        company.CompanyPassWord,
		AccountDescribe: company.CompanyDespcrite,
		Money:           float64(bill.Price),
		CompanyName:     bill.Company,
		SubCompanyName:  bill.BranchShop,
	}*/
	// save orders to mysql
	bill.TradeStatus = models.Unpaid
	db.DbClient.Client.Create(bill)

	if v := config.Opts().SubAccountNoInfo[company.CompanyBanchName]; v == "" {
		glog.Errorln("company branch don't have subaccount number!")

		if dberr := db.DbClient.Client.Model(&bill).Update("trade_status", models.PayFailed); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", "", "SaveBill"), "", dberr.Error)
		}
		c.JSON(http.StatusOK, gin.H{"server": "company branch don't have subaccount number!"})
		return
	}

	// 向企业账户转账
	if err := TransferToCompany(*bill); err != nil {
		glog.Errorln("Transfer to company err:", err, string(body))
		c.JSON(http.StatusOK, gin.H{"server": "error transfer to company fall"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"server": "success", "echo": bill})

}

// 转账给企业用户
func TransferToCompany(billInfo models.Bill) error {

	if config.Opts().PaySwitch > 0 { //支付宝支付
		uid, err := uuid.NewV4()
		if err != nil {
			return err
		}
		taReq := alipayModels.TransferToAccountReq{
			OrderId:       fmt.Sprintf("%s", uid), //需要唯一id，例如uuid
			Amount:        billInfo.Price,
			PayeeType:     "ALIPAY_LOGONID",
			PayeeAccount:  config.Opts().AliPayAddress,
			PayerShowName: billInfo.BranchShop,
			PayeeRealName: "",
			Remark:        "",
		}
		tAccount, err := payment.TransferToAccount(taReq)
		if err != nil {
			glog.Errorln(lib.Log("transfer to account err", "", "TransferToCompany"), "", err)

			if dberr := db.DbClient.Client.Model(&billInfo).Update("trade_status", models.PayFailed); dberr.Error != nil {
				glog.Errorln(lib.Log("update trade status to database err", "", "TransferToCompany"), "", dberr.Error)
			}
			return err
		}

		if dberr := db.DbClient.Client.Model(&billInfo).Update("trade_status", models.PayOK); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", "", "TransferToCompany"), "", dberr.Error)
		}

		fmt.Println("转账成功*****************************")
		glog.Infoln("转账成功*****************************")

		transferInfo, err := json.Marshal(tAccount)
		if err != nil {
			glog.Errorln(lib.Log("marshal account info err", "", "TransferToCompany"), "", err)
		}

		if dberr := db.DbClient.Client.Model(&billInfo).Update("transfer_Info", string(transferInfo)); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", "", "TransferToCompany"), "", dberr.Error)
		}

		if err := AddOrderToServer(*tAccount, billInfo); err != nil {
			if dberr := db.DbClient.Client.Model(&billInfo).Update("trade_status", models.AddTradeFailed); dberr.Error != nil {
				glog.Errorln(lib.Log("update trade status to database err", "", "TransferToCompany"), "", dberr.Error)
			}
			glog.Errorln(lib.Log("transfer to account err", "", "TransferToCompany"), "", err)
			return err
		}

	}

	/*	if err := AddOrderToServerWithoutPay(alipayModels.TransferToAccountResp{TradeNo: strconv.Itoa(int(billInfo.OrderId)),
			OrderId: strconv.Itoa(int(billInfo.OrderId))}, billInfo); err != nil {
			if dberr := db.DbClient.Client.Model(&billInfo).Update("trade_status", models.AddTradeFailed); dberr.Error != nil {
				glog.Errorln(lib.Log("update trade status to database err", "", "TransferToCompany"), "", dberr.Error)
			}
			glog.Errorln(lib.Log("transfer to account err", "", "TransferToCompany"), "", err)
			return err
		}
	*/
	if dberr := db.DbClient.Client.Model(&billInfo).Update("trade_status", models.TransferOK); dberr.Error != nil {
		glog.Errorln(lib.Log("update trade status to database err", "", "TransferToCompany"), "", dberr.Error)
	}

	return nil
}

// 从数据库获取转账失败的流水信息，重新转账
func HandleTransferToCompany() error {

	bills := []models.Bill{}
	if dberr := db.DbClient.Client.Where("trade_status = ?",
		models.PayFailed).Find(&bills); dberr.Error != nil {
		glog.Errorln(lib.Log("get account from database err", "", "HandleTransferToCompany"), "", dberr.Error)
		return nil
	}

	for _, bill := range bills {
		fmt.Println("HandleTransferToCompany===================转账失败的订单Bill.ID=", bill.ID)
		//重新转账
		billOp := models.Bill{}
		tx := db.DbClient.Client.Begin()
		if dberr := tx.Where("id = ? and trade_status = ?",
			bill.ID, models.PayFailed).First(&billOp); dberr.Error != nil {
			glog.Errorln(lib.Log("get account from database err", "", "HandleTransferToCompany"), "", dberr.Error)
			tx.Rollback()
			continue
		}
		if dberr := db.DbClient.Client.Model(&billOp).Update("trade_status", models.Paying); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", "", "HandleTransferToCompany"), "", dberr.Error)
			tx.Rollback()
			continue
		}
		tx.Commit()

		if err := TransferToCompany(billOp); err != nil {
			glog.Errorln(lib.Log("transfer to company err", "", "HandleTransferToCompany"), "", err)
			continue
		}
	}

	return nil
}

// 从数据库获取添加到编号系统失败的交易流水，重新添加
func HandleTransferToOrderServer() error {

	bills := []models.Bill{}
	if dberr := db.DbClient.Client.Where("trade_status = ?",
		models.AddTradeFailed).Find(&bills); dberr.Error != nil {
		glog.Errorln(lib.Log("get account from database err", "", "HandleTransferToOrderServer"), "", dberr.Error)
		return nil
	}

	for _, bill := range bills {
		fmt.Println("HandleTransferToOrderServer===================添加到编号系统失败的订单Bill.ID=", bill.ID)
		//重新添加到编号系统
		tAccount := alipayModels.TransferToAccountResp{}
		if err := json.Unmarshal([]byte(bill.TransferInfo), &tAccount); err != nil {
			glog.Errorln(lib.Log("unmarshall transfer info err", "", "HandleTransferToOrderServer"), "", err)
			continue
		}
		if err := AddOrderToServer(tAccount, bill); err != nil {
			if dberr := db.DbClient.Client.Model(&bill).Update("trade_status", models.AddTradeFailed); dberr.Error != nil {
				glog.Errorln(lib.Log("update trade status to database err", "", "HandleTransferToOrderServer"), "", dberr.Error)
			}
			glog.Errorln(lib.Log("transfer to company err", "", "HandleTransferToOrderServer"), "", err)
			continue
		}

		if dberr := db.DbClient.Client.Model(&bill).Update("trade_status", models.TransferOK); dberr.Error != nil {
			glog.Errorln(lib.Log("update trade status to database err", "", "HandleTransferToOrderServer"), "", dberr.Error)
		}
	}

	return nil
}

// 添加账单信息到编号服务
func AddOrderToServer(tAccount alipayModels.TransferToAccountResp, billInfo models.Bill) error {
	fmt.Println("tAccount.TradeNo=", tAccount.TradeNo)
	fmt.Println("tAccount.OrderId=", tAccount.OrderId)
	aliTransfer, err := payment.QueryTransferToAccount(tAccount.TradeNo)
	if err != nil {
		glog.Errorln(lib.Log("query ali trade err, trade num: ", tAccount.TradeNo, "QueryAliTrade"), "", err)
		return err
	}
	tout, err := json.Marshal(aliTransfer)
	if err != nil {
		glog.Errorln(lib.Log("aliTransfer json marshal err, trade num: ", tAccount.TradeNo, "QueryAliTrade"), "", err)
		return err
	}
	glog.Infoln("aliTransfer.PayDate=", aliTransfer.PayDate)
	orderSvc := models.OrderServer{
		SubAccountNo:   config.Opts().SubAccountNoInfo[billInfo.BranchShop],
		ThirdTradeNo:   tAccount.TradeNo,
		OrderNo:        tAccount.OrderId,
		Company:        "zhengdao",
		OrderType:      0,
		PaymentType:    1,
		TransferAmount: billInfo.Price,
		TransferInfo:   string(tout),
		OrderTime:      aliTransfer.PayDate,
		OrderState:     1,
	}
	out, err := json.Marshal(orderSvc)
	if err != nil {
		return err
	}

	//发送添加账单请求
	addOrderUrl := fmt.Sprintf("http://%s%s/addOrder", config.Opts().OrderSvcAddress, config.Opts().OrderSvcPort)
	glog.Infoln("addOrderUrl=", addOrderUrl)
	_, err = SendHttpRequest("POST", addOrderUrl, bytes.NewReader(out), nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("add order err", "", "GetOrderListFromServer"), "", err)
		return err
	}
	return nil
}

// 添加账单信息到编号服务
func AddOrderToServerWithoutPay(tAccount alipayModels.TransferToAccountResp, billInfo models.Bill) error {
	fmt.Println("tAccount.TradeNo=", tAccount.TradeNo)
	fmt.Println("tAccount.OrderId=", tAccount.OrderId)

	orderSvc := models.OrderServer{
		SubAccountNo:   config.Opts().SubAccountNoInfo[billInfo.BranchShop],
		ThirdTradeNo:   tAccount.TradeNo,
		OrderNo:        tAccount.OrderId,
		Company:        "zhengdao",
		OrderType:      0,
		PaymentType:    1,
		TransferAmount: billInfo.Price,
		OrderTime:      time.Now(),
		OrderState:     1,
	}
	out, err := json.Marshal(orderSvc)
	if err != nil {
		return err
	}

	//发送添加账单请求
	addOrderUrl := fmt.Sprintf("http://%s%s/addOrder", config.Opts().OrderSvcAddress, config.Opts().OrderSvcPort)
	glog.Infoln("addOrderUrl=", addOrderUrl)
	_, err = SendHttpRequest("POST", addOrderUrl, bytes.NewReader(out), nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("add order err", "", "GetOrderListFromServer"), "", err)
		return err
	}
	return nil
}

// 发送http请求
func SendHttpRequest(method, url string, body io.Reader, head, fdata map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return []byte{}, err
	} else if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("http response status code: %v", resp.StatusCode)
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}
