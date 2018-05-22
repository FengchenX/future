package apiserver

import (
	"time"
	"strings"
	"znfz/server/token-contract/subaccount"
	"strconv"
	"math/big"
	"znfz/server/protocol"
	"golang.org/x/net/context"
	"znfz/server/lib"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
)

/****************************************************************************************
*  现在已经没用的接口，删除是不可能删除的，这辈子都不可能删除的，只能安安静静的呆在这里
*  才能维持的了生活的样子
****************************************************************************************/

// 支付服务
func (this *ApiService) Pay(ctx context.Context, req *protocol.ReqPay) (*protocol.RespPay, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "提交订单"), req)

	if req.GetContent() == nil {
		glog.Errorln("[api err]", "some of the req is nil")
		return &protocol.RespPay{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	buffer, err := proto.Marshal(req.Content)
	if err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespPay{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	if req.GetUserAddress() == "" || req.GetJobAddress() == "" || req.GetAccountDescribe() == "" {
		glog.Errorln("[api err]", "some of the req is nil")
		return &protocol.RespPay{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	orderId := time.Now().Unix()
	jAddr := strings.ToLower(req.GetJobAddress())
	hash, err := subaccount.SetOrdersContent(req.GetAccountDescribe(), req.GetPassWord(),
		jAddr, strconv.Itoa(int(orderId)), string(buffer))
	if hash == "" || err != nil {
		glog.Errorln("[api err]", jAddr, err)
		return &protocol.RespPay{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "支付服务"), req)

	hash, err = subaccount.SettleAccounts(req.GetAccountDescribe(), req.GetPassWord(),
		jAddr, big.NewInt(int64(req.GetMoney())))
	if hash == "" || err != nil {
		glog.Errorln("[api err]", jAddr, err)
		return &protocol.RespPay{
			StatusCode: uint32(protocol.Status_RegistError),
		}, nil
	}

	return &protocol.RespPay{
		StatusCode: uint32(protocol.Status_Success),
		HashCode:   hash,
		OrderId:    uint32(orderId),
	}, nil
}

// 提交订单（比如说饭店客户下单）
func (this *ApiService) SetContent(ctx context.Context, req *protocol.ReqSetContent) (*protocol.RespSetContent, error) {
	defer this.RunMiddleWare()
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "提交订单"), req)

	buffer, err := proto.Marshal(req.Content)
	if err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespSetContent{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	if req.GetUserAddress() == "" || req.GetJobAddress() == "" || req.GetAccountDescribe() == "" {
		glog.Errorln("[api err]", "some of the req is nil")
		return &protocol.RespSetContent{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	hash, err := subaccount.SetOrdersContent(req.GetAccountDescribe(), req.GetPassWord(),
		req.GetJobAddress(), req.OrderId, string(buffer))
	if hash == "" || err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespSetContent{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	return &protocol.RespSetContent{
		StatusCode: uint32(protocol.Status_Success),
	}, nil
}