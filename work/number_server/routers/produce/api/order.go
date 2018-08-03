package api

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"sub_account_service/finance/lib"

	"sub_account_service/number_server/models"
	"sub_account_service/number_server/pkg/nsq"

	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"sub_account_service/number_server/config"
	"sub_account_service/number_server/pkg/log"
)

func AddOrder(c *gin.Context) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		glog.Errorln("request  error")
		c.JSON(http.StatusOK, lib.Result.Fail(-1, "request错误"))
		return
	}
	order := models.OrderServer{}
	err = json.Unmarshal(body, &order)
	if err != nil {
		log.Error("AddOrder 订单解析失败",&order)
		c.JSON(http.StatusOK, lib.Result.Fail(-1, err.Error()))
		return
	}
	//验证数据有效性
	err = models.Validate(&order)
	if err != nil {
		c.JSON(http.StatusOK, lib.Result.Fail(-1, err.Error()))
		return
	}

	var orderNew = models.Orders{
		ThirdTradeNo:   order.ThirdTradeNo,
		OrderNo:        order.OrderNo,
		SubAccountNo:   order.SubAccountNo,
		MchId:      order.MchId,
		Company:        order.Company,
		BranchShop:     order.BranchShop,
		OrderType:      order.OrderType,
		PaymentType:    order.PaymentType,
		TransferAmount: order.TransferAmount,
		TransferInfo:   order.TransferInfo,
		AutoTransfer:   order.AutoTransfer,
		OrderTime:      order.OrderTime,  //下单时间
		OrderState:     order.OrderState, //订单状态
	}

	msg, _ := json.Marshal(&orderNew)
	//加入到消息队列写数据库
	producer.Publish("order", msg)

	log.Info("AddOrder",orderNew)

	c.JSON(http.StatusOK, lib.Result.Success("添加成功！", order))
}

var producer *nsq.Producer

func Setup() error {
	var err error
	producer, err = nsq.NewProducer(config.Opts().Nsqd_Producer_Tcp)
	if err != nil {
		glog.Errorln(err)
		return err
	}
	glog.Info("connecting to nsqd")
	return nil
}
