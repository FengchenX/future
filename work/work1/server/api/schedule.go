/*--------------------------------------------------------------
* package: 排班相关的服务
* time:    2018/04/17
*-------------------------------------------------------------*/

package apiserver

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"math/big"
	"strconv"
	"time"
	"znfz/server/arguments"
	"znfz/server/lib"
	"znfz/server/protocol"
	"znfz/server/token-contract/accmanager"
	"znfz/server/token-contract/addrmanager"
	"znfz/server/token-contract/subaccount"
)

// 发布排班服务
func (this *ApiService) SetSchedule(ctx context.Context, req *protocol.ReqScheduling) (*protocol.RespScheduling, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "发布排班"), req)
	var counts, ratio, jobIds []*big.Int
	var whiteLists []string
	roles := make([][32]byte, 0)

	if req.GetCompany() == "" {
		glog.Errorln("[api err] 发布企业为空")
		return &protocol.RespScheduling{
			StatusCode: uint32(protocol.Status_SetSchedueFail),
		}, nil
	}

	if req.GetUserAddress() == "" {
		glog.Errorln("[api err] 发布者地址为空")
		return &protocol.RespScheduling{
			StatusCode: uint32(protocol.Status_SetSchedueFail),
		}, nil
	}
	acc, err := accmanager.GetAccountByAddr(this.AccountAddress, req.UserAddress)
	glog.Infoln("查询账户", req.GetPassWord(), req.GetAccountDescribe(), acc)
	if err != nil {
		glog.Errorln("[api err] 查询账户失败", err)
		return &protocol.RespScheduling{
			StatusCode: uint32(protocol.Status_SetSchedueFail),
		}, nil
	}

	ids := int64(1)
	for _, job := range req.GetJobs() {
		for j := 0; j < int(job.GetCount()); j++ {
			roles = append(roles, lib.StrToByte32(job.GetRole()))
			counts = append(counts, big.NewInt(1))
			ratio = append(ratio, big.NewInt(int64(job.GetRadio())))
			jobIds = append(jobIds, big.NewInt(ids))
			whiteLists = append(whiteLists, job.GetWhiteList()...)
			ids++
		}
	}

	// 部署排班合约
	store := req.GetCompany() + req.GetStoresNumber()
	arg := arguments.DeployArguments{
		OperationKeyStore: req.GetAccountDescribe(), // key 是很长的字符串（key_store）
		OperationPassWord: req.GetPassWord(),
		TokenName:         req.GetCompany(),
		TokenSymbol:       strconv.Itoa(int(req.GetTimeStamp())),
		SubPayer:          req.GetPayAccount(),                      // 分账时的支付账户地址
		ManagerPayee:      req.GetManagerPayee(),                    // 管理费的收款地址
		ManagerRatio:      big.NewInt(int64(req.GetManagerRatio())), // 管理费的固定比例
		StoresNumber:      store,                                    // 门店编号
		Postscript:        req.GetPostscript(),                      // 备注信息
	}
	glog.Infoln("部署排班合约:", store, arg)
	addr, _, err := subaccount.Deploy(arg)
	if err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespScheduling{
			StatusCode: uint32(protocol.Status_SetSchedueFail),
		}, nil
	}
	glog.Infoln("发布成功", addr)

	// 绑定查询
	str, err := addrmanager.NewAddressAdd(
		arguments.BindSmartArguments{
			OperationKeyStore: req.GetAccountDescribe(),
			OperationPassWord: req.GetPassWord(),
			OperatingAddress:  this.AddressManager,
			SmartAddress:      addr,
			TokenName:         req.GetCompany(),
			TokenSymbol:       strconv.Itoa(int(req.GetTimeStamp())),
			StoresNumber:      store,
		})
	if err != nil {
		glog.Errorln("[api err]", str, err)
		return &protocol.RespScheduling{
			StatusCode: uint32(protocol.Status_SetSchedueFail),
		}, nil
	}

	// 排班发布
	publisher := arguments.ScheduleArguments{
		OperationKeyStore: req.GetAccountDescribe(),
		OperationPassWord: req.GetPassWord(),
		SmartAddress:      addr,
		FatherAddress:     req.GetUserAddress(),
		IssueRatio:        big.NewInt(int64(req.GetMyRatio())),
		Roles:             roles,
		Counts:            counts,
		Ratio:             ratio,
		JobIds:            jobIds,
		Whitelists:        whiteLists,
	}

	glog.Infoln("publisher ===> ", publisher)
	hash, err := subaccount.PublishScheduleing(publisher)

	if hash == "" || err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespScheduling{
			StatusCode: uint32(protocol.Status_SetSchedueFail),
		}, nil
	}
	return &protocol.RespScheduling{
		StatusCode: uint32(protocol.Status_Success),
		JobAddress: addr,
	}, nil
}

func (this *ApiService) GetCanJoin(ctx context.Context, req *protocol.ReqGetSchedue) (*protocol.RespGetSchedue, error) {
	defer this.RunMiddleWare()
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询当前排班"), req)
	var list []*protocol.Schedule
	list = this.checkList(req.GetTimeStamp(), 5)

	if len(list) == 0 {
		glog.Infoln(lib.Log("api err", req.GetUserAddress(), "查询排班接口失败：没有查到排班！"))
		return &protocol.RespGetSchedue{
			StatusCode: uint32(protocol.Status_GetSchedueFail),
		}, nil
	}

	return &protocol.RespGetSchedue{
		StatusCode: uint32(protocol.Status_Success),
		Schedules:  this.checkList(req.GetTimeStamp(), 5),
	}, nil
}

// 查询排班接口
func (this *ApiService) GetSchedule(ctx context.Context, req *protocol.ReqGetSchedue) (*protocol.RespGetSchedue, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询排班"), req)
	var list []*protocol.Schedule
	list = this.checkList(req.GetCompanyName(), 5)

	if len(list) == 0 {
		glog.Infoln(lib.Log("api err", req.GetUserAddress(), "查询排班接口失败：没有查到排班！"))
		return &protocol.RespGetSchedue{
			StatusCode: uint32(protocol.Status_Success),
		}, nil
	}
	resp := &protocol.RespGetSchedue{
		StatusCode: uint32(protocol.Status_Success),
		Schedules:  list,
	}
	glog.Infoln(lib.Log("api", "", "查询排班回包"), resp)
	return resp, nil
}

// 查询所有的排班信息
// todo 排班的订单和申请信息
func (this *ApiService) checkList(key string, timeout int) []*protocol.Schedule {
	defer this.RunMiddleWare()

	ret := make([]*protocol.Schedule, 0)
	// 查询本人发布的所有排班的地址
	glog.Infoln(lib.Log("api", "", "查询所有排班"), key)

	addrs, err := addrmanager.GetAddressByKey(this.AddressManager, key)
	if err != nil {
		glog.Errorln("[api err]", err)
	}
	glog.Infoln(lib.Log("api", "", "addrs"), addrs)
	c := make(chan *protocol.Schedule, len(addrs))
	enter_sec := time.Now().Unix()
	// 根据地址查询每一个排班的信息
	for _, addr := range addrs {
		go func(address string) {
			// 查询排班接口
			schs, time, err := subaccount.GetScheduleingCxt(address, all)
			t, _ := strconv.Atoi(time)

			company, _ := subaccount.GetTokenName(address)

			if err != nil {
				glog.Errorln("[api err]", err)
			}
			tmps := []*protocol.Job{}

			company_ratio := int64(0)
			fatherAddr := ""
			for _, v := range schs {
				if v.AllCount == nil || v.IssueRatio.Int64() == 0 {
					continue
				}
				company_ratio = v.IssueRatio.Int64()
				fatherAddr = v.Issuer
				for _, vv := range v.IssuerDesireds {
					tmp := &protocol.Job{}
					tmp.Role = lib.Byte32ToStr(vv.Role)
					tmp.Count = uint32(vv.Count.Int64())
					tmp.Radio = uint64(vv.Ratio.Int64())
					tmp.JobId = int32(vv.JobId.Int64())
					tmp.WhiteList = []string{vv.Whites}
					acco, _ := accmanager.GetAccountByAddr(this.AccountAddress, vv.Whites)
					tmp.WhiteName = acco.Name
					tmp.TimeStamp = int64(t)
					tmp.JobAddress = address
					tmp.Company = company
					tmps = append(tmps, tmp)
				}
			}
			cd := &protocol.Schedule{
				Jobaddr:       address,
				Jobs:          tmps,
				TimeStamp:     int64(t),
				CompanyName:   company,
				CompanyRatio:  int64(100) - company_ratio,
				FatherAddress: fatherAddr,
			}
			select {
			case c <- cd:
			default:
			}
		}(addr)
	}
	break_flag := false
	tick := time.NewTicker(1 * time.Second)
	count := 0
	for {
		if break_flag == true {
			break
		}
		select {
		case sch := <-c:
			ret = append(ret, sch)
			count++
			if count >= len(addrs) {
				// sort
				le := len(ret)
				for i := 0; i < le-1; i++ {
					for j := 0; j < le-1-i; j++ {
						if ret[j].TimeStamp < ret[j+1].TimeStamp {
							ret[j+1], ret[j] = ret[j], ret[j+1]
						}
					}
				}
				break_flag = true
				break
			}
		case <-tick.C:
			interval := int(time.Now().Unix() - enter_sec)
			if interval > timeout {
				break_flag = true
				break
			}
		}
	}
	tick.Stop()

	return ret
}

// 查询所有可以申请的工作
func (this *ApiService) GetJob(ctx context.Context, req *protocol.ReqGetCanApply) (*protocol.RespGetCanApply, error) {
	defer this.RunMiddleWare()
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询所有可以申请的工作"), req)

	list := this.checkList(req.CompanyName, 5)

	if len(list) == 0 {
		glog.Infoln(lib.Log("api err", req.GetUserAddress(), "查询排班接口失败：没有查到排班！"))
		return &protocol.RespGetCanApply{
			StatusCode: uint32(protocol.Status_Success),
		}, nil
	}

	jobs := make([]*protocol.Job, 0)
	for _, v := range list {
		jobs = append(jobs, v.Jobs...)
	}

	// 判断是否已申请
	for _, job := range jobs {
		flag, _ := subaccount.CheckIsOkApplication(job.JobAddress, req.GetUserAddress())
		job.HasApply = flag
	}

	resp := &protocol.RespGetCanApply{
		StatusCode: uint32(protocol.Status_Success),
		Jobs:       jobs,
	}
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询所有可以申请的工作"), resp)
	return resp, nil
}

// 申请工作服务
func (this *ApiService) ApplyJob(ctx context.Context, req *protocol.ReqFindJob) (*protocol.RespFindJob, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "申请工作"), req)

	// 业务层先判断用户符不符合申请要求
	account, err := accmanager.GetAccountByAddr(this.AccountAddress, req.GetUserAddress())

	// 判断账户信息是否完整
	if account.Password == "" || account.Name == "" {
		glog.Errorln("[api fail]未查询到该玩家", account, err)
		return &protocol.RespFindJob{
			StatusCode: uint32(protocol.Status_ReqNotFull),
		}, nil
	}

	// 申请信息不完整
	if req.GetMyJob() == nil || req.GetMyJob().GetJobAddress() == "" || req.GetMyJob().GetCompany() == "" {
		if account.Password == "" || account.Name == "" {
			glog.Errorln("[api fail]申请信息不完整", account, err)
			return &protocol.RespFindJob{
				StatusCode: uint32(protocol.Status_FindJobFail),
			}, nil
		}
	}

	// 查询是否可以申请
	staffs, _, err := subaccount.GetScheduleingCxt(req.GetMyJob().GetJobAddress(), all)
	glog.Infoln("查询是否可以申请", staffs)
	for _, v := range staffs {
		if v.AllCount.Int64() < 0 {
			glog.Errorln(lib.Log("api err", req.GetUserAddress(), "申请人数已满"))
			return &protocol.RespFindJob{
				StatusCode: uint32(protocol.Status_ApplyFull),
			}, nil
		}
	}

	// 查询是否重复申请
	flag, _ := subaccount.CheckIsOkApplication(req.GetMyJob().GetJobAddress(), req.GetUserAddress())
	if flag {
		glog.Errorln(lib.Log("api err", req.GetUserAddress(), "已经申请过"))
		return &protocol.RespFindJob{
			StatusCode: uint32(protocol.Status_HasAppliedJobFail),
		}, nil
	}

	// 申请工作
	hash, err := subaccount.ApplicationJob(req.GetAccountDescribe(),
		req.GetPassWord(),
		req.GetMyJob().GetJobAddress(),
		req.GetUserAddress(),
		big.NewInt(int64(req.GetMyJob().GetJobId())),
		req.GetMyJob().GetFatherAddr())
	if hash == "" || err != nil {
		glog.Errorln("[api err]", err)
		return &protocol.RespFindJob{
			StatusCode: uint32(protocol.Status_FindJobFail),
		}, nil
	}

	glog.Infoln("[api 申请工作成功]", hash)
	return &protocol.RespFindJob{
		StatusCode: uint32(protocol.Status_Success),
		HashCode:   hash,
	}, nil
}

// 查询是否申请
func (this *ApiService) CheckIsOkApplication(ctx context.Context, req *protocol.ReqCheckIsOkApplication) (*protocol.RespCheckIsOkApplication, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询是否申请"), req)
	flag, err := subaccount.CheckIsOkApplication(req.GetJobAddress(), req.GetUserAddress())
	var StatusCode uint32 = uint32(protocol.Status_Success)
	if err != nil {
		StatusCode = uint32(protocol.Status_HasAppliedJobFail)
	}
	resp := &protocol.RespCheckIsOkApplication{
		StatusCode: StatusCode,
		IsApplied:  flag,
	}
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询是否申请"), flag)
	return resp, err
}

// 查询申请情况
func (this *ApiService) GetApply(ctx context.Context, req *protocol.ReqGetStaff) (*protocol.RespGetStaff, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询申请情况"), req)
	staffs, err := subaccount.GetApplicationCxt(req.GetJobAddress())
	if err != nil {
		glog.Errorln("[api err]", err)
	}
	ret := make([]*protocol.ScheduleStaff, 0)
	for _, v := range staffs {
		ret = append(ret, &protocol.ScheduleStaff{
			StaffAddr:  v.StaffAddr,
			Role:       lib.Byte32ToStr(v.Role),
			Ratio:      v.Ratio.Uint64(),
			JobId:      v.JobId.Int64(),
			FatherAddr: v.FatherAddress,
		})
	}

	resp := &protocol.RespGetStaff{
		StatusCode: uint32(protocol.Status_Success),
		Staffs:     ret,
	}

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "查询申请情况"), resp)
	return resp, nil
}

// 查询玩家加入历史
func (this *ApiService) HistoryJoin(ctx context.Context, req *protocol.ReqHistoryJoin) (*protocol.RespHistoryJoin, error) {
	list := this.checkList(req.GetCompany(), 5)

	if len(list) == 0 {
		glog.Infoln(lib.Log("api err", req.GetUserAddress(), "查询排班接口失败：没有查到排班！"))
		return &protocol.RespHistoryJoin{
			StatusCode: uint32(protocol.Status_Success),
		}, nil
	}

	addr := req.GetUserAddress()
	jobs := make([]*protocol.Job, 0)
	for _, v := range list {
		staffs, _ := subaccount.GetApplicationCxt(v.GetJobaddr())
		for _, s := range staffs {
			if s.StaffAddr == addr {
				jobs = append(jobs, &protocol.Job{
					Role:      lib.Byte32ToStr(s.Role),
					Company:   req.GetCompany(),
					TimeStamp: v.TimeStamp,
				})
			}
		}
	}
	if len(jobs) == 0 {
		glog.Errorln(lib.Log("api err", req.UserAddress, "查询玩家加入历史失败：未找到历史"))
	}
	resp := &protocol.RespHistoryJoin{
		StatusCode: uint32(protocol.Status_Success),
		Jobs:       jobs,
	}
	glog.Infoln(lib.Log("api", req.UserAddress, "查询玩家加入历史"), resp)
	return resp, nil
}

// 4.3.获取当前排班
func (this *ApiService) GetNowJobAddress(ctx context.Context, req *protocol.ReqGetNowJobAddr) (*protocol.RespGetNowJobAddr, error) {
	glog.Infoln(lib.Log("api", "", "获取当前排班"), req.GetSubCompanyName())

	jobAddr, err := addrmanager.GetPaySamrtAddress(this.AddressManager, req.CompanyName+req.GetSubCompanyName())
	if err != nil {
		return &protocol.RespGetNowJobAddr{
			StatusCode: uint32(protocol.Status_RoleNotFitFail),
		}, nil
	}
	return &protocol.RespGetNowJobAddr{
		StatusCode: uint32(protocol.Status_Success),
		JobAddr:    jobAddr,
	}, nil
}
