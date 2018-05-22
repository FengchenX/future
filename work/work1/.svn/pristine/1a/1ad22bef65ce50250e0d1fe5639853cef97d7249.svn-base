package main

import (
	"github.com/golang/glog"
	ctx "golang.org/x/net/context"
	"time"
	"znfz/server/protocol"
)

type Client struct {
	C protocol.ApiServiceClient
}

// 测试
func (this *Client) Check() {
	this.C.SayHello(ctx.Background(), &protocol.Req{
		Name: "liangsihao",
	})
}

// 注册服务
func (this *Client) SetRegister(password string) *protocol.RespRegister {
	usr, _ := this.C.Register(ctx.Background(), &protocol.ReqRegister{
		PassWord: password,
	})
	glog.Infoln("1--注册服务", usr)
	glog.Infoln("1--注册服务", usr.GetAccountDescribe())
	time.Sleep(10 * time.Second)
	return usr
}

func (this Client) band(psw, name, desp, addr, phone string) {
	glog.Infoln(desp)
	req := &protocol.ReqBand{
		Name:            name,
		PassWord:        psw,
		Phone:           phone,
		AccountDescribe: desp,
		UserAddress:     addr,
	}
	glog.Infoln("1.1--绑定服务 1", req)
	resp, _ := this.C.Bind(ctx.Background(), req)
	glog.Infoln("1.1--绑定服务 2", resp.GetStatusCode())
}

// 注册服务
func (this *Client) GetRegister(address string) {
	acc, _ := this.C.GetAccount(ctx.Background(), &protocol.ReqCheckAccount{
		UserAddress: address,
	})
	glog.Infoln("2--查询账户", acc, acc.StatusCode)
}

// 发布排班
func (this *Client) SetSchedule(addr,payAccount, pass, desp, name, company, role string, timestamp int64) {

	glog.Infoln("发布排班", timestamp)
	resp, _ := this.C.SetSchedule(ctx.Background(), &protocol.ReqScheduling{
		UserAddress:     addr,
		PassWord:        pass,
		AccountDescribe: desp,
		Company:         company,
		TimeStamp:       timestamp,
		PayAccount:      payAccount,
		MyRatio:         uint32(88),
		StoresNumber:    "南山分店",
		ManagerPayee:    addr,
		ManagerRatio:    int64(12),
		Jobs: []*protocol.Job{
			&protocol.Job{
				JobAddress: "",
				Role:       role,
				Count:      uint32(3),
				Radio:      uint64(10),
				Company:    company,
				WhiteList:  []string{managerAddr, huangAddr, cookerAddr},
			},
		&protocol.Job{
			JobAddress: "",
			Role:       "小可爱",
			Count:      uint32(2),
			Radio:      uint64(10),
			Company:    company,
			WhiteList:  []string{managerAddr, cookerAddr},
		}},
	})
	glog.Infoln("3--发布排班", resp.GetJobAddress())
}

// 查询排班
func (this *Client) GetSchedule(userAddr, company, timestamp string) {
	resp, err := this.C.GetSchedule(ctx.Background(), &protocol.ReqGetSchedue{
		UserAddress: userAddr,
		CompanyName: company,
		TimeStamp:   timestamp,
	})
	glog.Infoln("3--查询排班", err)
	for i, v := range resp.GetSchedules() {
		glog.Infoln(i, "--查询排班", v.GetTimeStamp(),v.GetCompanyRatio(), v.GetCompanyName(), v.GetJobaddr())
		for _, vv := range v.GetJobs() {
			glog.Infoln("    -------查询", vv.GetJobId(), vv.GetWhiteList(), vv.GetRole(), vv.GetCompany(), vv.GetRadio(), vv.GetCount(), time.Unix(v.GetTimeStamp()/1000, 0).Format("2006-01-02 15:04:05"))
		}
	}
}

// 查询排班
func (this *Client) GetLastDeploy(company, sub, jaddr string) {
	resp, err := this.C.GetNowJobAddress(ctx.Background(), &protocol.ReqGetNowJobAddr{
		CompanyName:    company,
		SubCompanyName: sub,
		TimeStamp:      time.Now().String(),
		JobAddress:     jaddr,
	})
	glog.Infoln("3--查询排班", resp.GetJobAddr(), err)
}

// 获取以太坊余额
func (this *Client) GetEthBalance(usr string) {
	resp, err := this.C.GetEthBalance(ctx.Background(), &protocol.ReqGetEthBalance{
		UserAddress: usr,
	})
	glog.Infoln("获取以太坊余额", resp, err)
}

func (this *Client) RpcTest() {
	t1 := time.Now()
	resp, err := this.C.SayHello(ctx.Background(), &protocol.Req{Name: "sihao "})
	if err != nil {
		glog.Errorln("Do Format error:" + err.Error())
	} else {
		glog.Infoln("client recving msg ->>", resp, " time:", time.Now().Sub(t1))
	}
}

func (this *Client) FindJob(managerAddr, cookerAddr, jobAddress, psw, desp, role string, id int) {
	resp, _ := this.C.ApplyJob(ctx.Background(), &protocol.ReqFindJob{
		UserAddress:     cookerAddr,
		PassWord:        psw,
		AccountDescribe: desp,
		MyJob: &protocol.Job{
			JobId:      int32(id),
			JobAddress: JobAddress,
			Role:       role,
			FatherAddr: managerAddr,
		},
	})
	glog.Infoln("4--申请上班", resp.StatusCode)
}

func (this *Client) GetCanapplyJob(company, cookerAddr string) {
	resp, _ := this.C.GetJob(ctx.Background(), &protocol.ReqGetCanApply{
		UserAddress: cookerAddr,
		CompanyName: company,
		TimeStamp:   "",
	})
	for i, v := range resp.GetJobs() {
		glog.Infoln(i, "--查询可以上班", v, v.GetJobId(), v.GetRole(), v.GetRole(), v.GetTimeStamp(), v.GetHasApply())
	}
}

func (this *Client) GetFindJob(usr, job string) {
	resp, err := this.C.CheckIsOkApplication(ctx.Background(), &protocol.ReqCheckIsOkApplication{
		UserAddress: usr,
		JobAddress:  job,
	})
	glog.Infoln("4--是否申请上班", resp.GetIsApplied(), err)
}

//
//func (this *Client) Order(id, Context, addr string, table uint32, money float64) {
//	order := &protocol.Order{
//		Table:     uint32(table),
//		TimeStamp: time.Now().String(),
//		Money:     float64(money),
//		Content:   Context,
//	}
//	resp, err := this.C.(ctx.Background(), &protocol.ReqSetContent{
//		OrderId:    id,
//		Content:    order,
//		JobAddress: addr,
//	})
//	glog.Infoln("5--发布订单", err, resp)
//}

func (this *Client) GetOrder(id, addr string) {
	resp, err := this.C.GetContent(ctx.Background(), &protocol.ReqGetContent{
		OrderId:    id,
		JobAddress: addr,
	})
	glog.Infoln("5--查询订单", err, string(resp.GetContent().GetContent()), resp)
}

func (this *Client) Pay(table, money uint64, userAddr, pass, desp, jaddr string) {
	order := &protocol.Order{
		Table: uint32(table),
		//TimeStamp: time.Now().String(),
		Money: float64(money),
	}
	resp, err := this.C.Pay(ctx.Background(), &protocol.ReqPay{
		UserAddress:     userAddr,
		PassWord:        pass,
		AccountDescribe: desp,
		Money:           money,
		Content:         order,
		JobAddress:      jaddr,
	})
	glog.Infoln("6--支付订单", err, resp)
}

func (this *Client) GetMoney(job_addr, account_addr string) {
	resp, err := this.C.GetBalance(ctx.Background(), &protocol.ReqGetBalance{
		SchedueAddress: job_addr,
		UserAddress:    account_addr,
	})
	glog.Infoln("6--查詢分張", resp.GetMoney(), " err:", err)
}

func (this *Client) GetApply(usr_Addr, job_addr string) {
	glog.Infoln("5--查詢申请!", usr_Addr, job_addr)
	resp, _ := this.C.GetApply(ctx.Background(), &protocol.ReqGetStaff{
		UserAddress: usr_Addr,
		JobAddress:  job_addr,
	})
	for _, s := range resp.GetStaffs() {
		glog.Infoln("5--查詢申请", s.GetStaffAddr(), string(s.GetRole()), s.GetRatio())
		if s.GetStaffAddr() == usr_Addr {
			glog.Infoln("          success!")
		}
	}
}

func (this *Client) Login(phone string) {
	resp, err := this.C.GetBind(ctx.Background(), &protocol.ReqLogin{
		Phone: phone,
	})
	glog.Infoln("8--登录", resp.GetLogins(), resp, " err:", err)
}

func (this *Client) GetAllOrder(user, job string) {
	resp, _ := this.C.GetAllOrderBySchedule(ctx.Background(), &protocol.ReqGetAllOrder{
		UserAddress: user,
		JobAddress:  job,
	})
	for _, v := range resp.GetOrders() {
		glog.Infoln("8--查詢所有訂單 sum = ", resp.GetSum(), " v = ", v)
	}
}

func (this *Client) GetAllMoney(user, company string) {
	resp, _ := this.C.GetAllMoney(ctx.Background(), &protocol.ReqGetAllMoney{
		UserAddress: user,
		CompanyName: company,
	})
	glog.Infoln("9--查詢所有收入 sum = ", resp.Sum)
}

func (this *Client) HistoryJoin(uaddr, comp string) {
	resp, _ := this.C.HistoryJoin(ctx.Background(), &protocol.ReqHistoryJoin{
		UserAddress: uaddr,
		Company:     comp,
	})
	glog.Infoln("HistoryJoin", resp)
}

func (this *Client) GetAllIncome(uaddr, comp string) {
	glog.Infoln("get all income")
	resp, _ := this.C.GetAllIncome(ctx.Background(), &protocol.ReqGetAllIncome{
		UserAddress: uaddr,
		CompanyName: comp,
	})
	orders := resp.GetOrders()
	for _, order := range orders {
		//t := order.TimeStamp / 1000
		//realt := time.Unix(t, 0).Format("2006-01-02 15:04:05")
		glog.Infoln("GetAllIncome", order.TimeStamp, order.GetMoney(), order.GetMoney_, string(order.GetContent()))
	}
}
