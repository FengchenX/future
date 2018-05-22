package apiserver

import (
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"math/big"
	"time"
	"znfz/server/lib"
	"znfz/server/protocol"
	"znfz/server/token-contract/subaccount"
)

// 6.1.三方存证
func (this *ApiService) ThreeSetOrder(ctx context.Context, req *protocol.ReqThreeSetOrder) (*protocol.RespThreeSetOrder, error) {
	defer this.RunMiddleWare()
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "提交订单提交订单"), req)

	buffer, err := proto.Marshal(req.Content)
	if err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespThreeSetOrder{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	if req.GetUserAddress() == "" || req.GetJobAddress() == "" || req.GetAccountDescribe() == "" {
		glog.Errorln("[api err]", "some of the req is nil")
		return &protocol.RespThreeSetOrder{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	hash, err := subaccount.SetOrdersContent(req.GetAccountDescribe(), req.GetPassWord(),
		req.GetJobAddress(), req.OrderId, string(buffer))

	if hash == "" || err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespThreeSetOrder{
			StatusCode: uint32(protocol.Status_SetContentFail),
		}, nil
	}

	glog.Infoln("success")
	return &protocol.RespThreeSetOrder{
		StatusCode: uint32(protocol.Status_Success),
	}, nil
}

// 6.2.三方支付
func (this *ApiService) ThreeSetBill(ctx context.Context, req *protocol.ReqThreeSetBill) (*protocol.RespThreeSetBill, error) {
	defer this.RunMiddleWare()
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "three 支付服务"), req)

	// 获取最近一次排班
	respj, err := this.GetNowJobAddress(context.Background(), &protocol.ReqGetNowJobAddr{
		CompanyName:    req.GetCompanyName(),
		SubCompanyName: req.GetSubCompanyName(),
		TimeStamp:      time.Now().String(),
	})

	if err != nil {
		glog.Errorln("[api err]", req.GetJobAddress(), err)
		return &protocol.RespThreeSetBill{
			StatusCode: uint32(protocol.Status_RegistError),
		}, nil
	}

	// 分账
	hash, err := subaccount.SettleAccounts(req.GetAccountDescribe(), req.GetPassWord(),
		respj.GetJobAddr(), big.NewInt(int64(req.GetMoney())))
	if hash == "" || err != nil {
		glog.Errorln("[api err]", req.GetJobAddress(), err)
		return &protocol.RespThreeSetBill{
			StatusCode: uint32(protocol.Status_RegistError),
		}, nil
	}

	return &protocol.RespThreeSetBill{
		StatusCode: uint32(protocol.Status_Success),
	}, nil
}
