/*--------------------------------------------------------------
*  package: 订单相关的服务
*  time:    2018/04/17
*-------------------------------------------------------------*/
package apiserver

import (
	"bytes"
	"github.com/golang/glog"
	"github.com/golang/protobuf/proto"
	"golang.org/x/net/context"
	"strconv"
	"znfz/server/lib"
	"znfz/server/protocol"
	"znfz/server/token-contract/addrmanager"
	"znfz/server/token-contract/subaccount"
)

// 查询订单
func (this *ApiService) GetContent(ctx context.Context, req *protocol.ReqGetContent) (*protocol.RespGetContent, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询订单"), req)
	context, _, err := subaccount.GetOrdersContent(req.GetJobAddress(), req.GetOrderId())
	if context == "" || err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespGetContent{
			StatusCode: uint32(protocol.Status_GetContentFail),
		}, nil
	}
	order := &protocol.Order{}
	proto.Unmarshal(bytes.NewBufferString(context).Bytes(), order)
	return &protocol.RespGetContent{
		StatusCode: uint32(protocol.Status_Success),
		Content:    order,
	}, nil
}

// 查询某一班的收入
func (this *ApiService) GetBalance(ctx context.Context, req *protocol.ReqGetBalance) (*protocol.RespGetBalance, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询收入"), req)
	money, err := subaccount.GetBalance(req.GetSchedueAddress(), req.GetUserAddress())
	if money == nil || err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespGetBalance{
			StatusCode: uint32(protocol.Status_GetBalanceFail),
		}, nil
	}
	return &protocol.RespGetBalance{
		StatusCode: uint32(protocol.Status_Success),
		Money:      uint64(money.Int64()),
	}, nil
}

// 獲取某一排班下所有訂單
func (this *ApiService) GetAllOrderBySchedule(ctx context.Context, req *protocol.ReqGetAllOrder) (*protocol.RespGetAllOrder, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "獲取所有訂單 in"), req)

	reOrders, sum := this.getorder(req.GetJobAddress(), req.GetUserAddress())

	resp := &protocol.RespGetAllOrder{
		StatusCode: uint32(protocol.Status_Success),
		Orders:     reOrders,
		Sum:        sum, // 用户挣的钱
	}
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "獲取所有訂單 out"), resp)
	return resp, nil
}

// 獲取用戶在某一公司所有的錢
func (this *ApiService) GetAllMoney(ctx context.Context, req *protocol.ReqGetAllMoney) (*protocol.RespGetAllMoney, error) {
	defer this.RunMiddleWare()
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "獲取用戶在某一公司所有的錢"), req)

	jobAddrs, err := addrmanager.GetAddressByKey(this.AddressManager, req.GetCompanyName())
	if err != nil {
		glog.Errorln("err", err)
		return &protocol.RespGetAllMoney{
			StatusCode: uint32(protocol.Status_GetContentFail),
		}, err
	}

	sum := int64(0)
	for _, jobAddr := range jobAddrs {
		num, _ := subaccount.GetBalance(jobAddr, req.GetUserAddress())
		sum += num.Int64()
	}

	return &protocol.RespGetAllMoney{
		StatusCode: uint32(protocol.Status_GetContentFail),
		Sum:        float64(sum),
	}, nil
}

// 获取用户收入详情
func (this *ApiService) GetAllIncome(ctx context.Context, req *protocol.ReqGetAllIncome) (*protocol.RespGetAllIncome, error) {
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "获取用户收入详情 GetAllIncome"), req)

	// 查询所有排班地址
	addrs, err := addrmanager.GetAddressByKey(this.AddressManager, req.GetCompanyName())
	if err != nil {
		glog.Errorln(lib.Log("api err", req.GetUserAddress(), "GetAllIncome"), err)
		return &protocol.RespGetAllIncome{
			StatusCode: uint32(protocol.Status_GetAllIncomeFail),
		}, nil
	}

	reOrders := make([]*protocol.Order, 0)
	for _, jobAddr := range addrs {

		// 用户是否在该排班中
		if orders, _ := this.getorder(jobAddr, req.GetUserAddress()); len(orders) != 0 {
			reOrders = append(reOrders, orders...)
		}
	}
	resp := &protocol.RespGetAllIncome{
		StatusCode: uint32(protocol.Status_Success),
		Orders:     reOrders,
	}
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "获取用户收入详情 GetAllIncome"), resp.GetOrders())
	return resp, nil
}

// 根据排班信息获取玩家收入详情
func (this *ApiService) GetIncomeBySchedule(ctx context.Context, req *protocol.ReqGetIncomeBySchedule) (*protocol.RespGetIncomeBySchedule, error) {

	// 查询订单信息
	addrs, err := addrmanager.GetAddressByKey(this.AddressManager, req.GetCompanyName())
	if err != nil {
		glog.Errorln(lib.Log("api err", req.GetUserAddress(), "GetAllIncome"), err)
		return &protocol.RespGetIncomeBySchedule{
			StatusCode: uint32(protocol.Status_GetAllIncomeFail),
		}, nil
	}

	reOrders := make([]*protocol.Order, 0)
	for _, jobAddr := range addrs {
		if orders, _ := this.getorder(jobAddr, req.GetUserAddress()); len(orders) != 0 {
			reOrders = append(reOrders, orders...)
		}
	}
	return &protocol.RespGetIncomeBySchedule{
		StatusCode: uint32(protocol.Status_Success),
		Orders:     reOrders,
	}, nil
}

// 根据排班地址获取订单详情
func (this *ApiService) getorder(jobAddress, userAddress string) ([]*protocol.Order, float64) {

	// 是否申请过
	flag, _ := subaccount.CheckIsOkApplication(jobAddress, userAddress)

	// 是否为发布者
	isOwner, _ := subaccount.CheckOwnerIsOk(jobAddress, userAddress)

	// 查询排班信息
	schs, time, _ := subaccount.GetScheduleingCxt(jobAddress, all)
	if len(schs) == 0 || schs[0].IssueRatio == nil {
		return nil, 0
	}
	reOrders := make([]*protocol.Order, 0)
	if flag || isOwner {
		percent := int64(0)
		for _, s := range schs[0].IssuerDesireds {
			if s.Whites == userAddress {
				percent = s.Ratio.Int64()
			}
		}
		// 如果用户是经理
		if isOwner {
			has := schs[0].IssueRatio.Int64()
			glog.Infoln("has =========>1 ", has)
			// 查询所有已申请的岗位
			staffs, err := subaccount.GetApplicationCxt(jobAddress)
			if err != nil {
				return nil, 0
			}

			glog.Infoln("has =========>3 ", has)
			percent = staffs[0].Ratio.Int64()
		}

		// 查询排班内所有订单
		orders, err := subaccount.GetAllContentHashCxt(jobAddress)
		if err != nil {
			glog.Errorln(lib.Log("api err", userAddress, "GetAllIncome err != nil"), err)
			return nil, 0
		}
		if len(orders) == 0 {
			glog.Errorln(lib.Log("api err", userAddress, "GetAllIncome len(orders) == 0"), err)
			return nil, 0
		}
		t, _ := strconv.Atoi(time)
		for _, order := range orders {
			o := &protocol.Order{}
			proto.Unmarshal(bytes.NewBufferString(order).Bytes(), o)
			o.JobAddress = jobAddress
			o.GetMoney_ = float64(o.GetMoney() * (float64(percent) / float64(100)))
			o.TimeStamp = int64(t)
			reOrders = append(reOrders, o)
		}
	}

	count := float64(0)
	for _, reorder := range reOrders {
		count += reorder.GetMoney_
	}

	return reOrders, count // count 为用户挣的钱
}
