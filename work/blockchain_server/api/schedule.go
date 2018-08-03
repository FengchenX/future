/*--------------------------------------------------------------
* package: 排班相关的服务
* time:    2018/04/17
*-------------------------------------------------------------*/

package api

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"

	"net/http"
	"strings"
	"sub_account_service/blockchain_server/arguments"
	"sub_account_service/blockchain_server/contracts"
	"sub_account_service/blockchain_server/lib"
	"sub_account_service/blockchain_server/lib/eth"
	"sub_account_service/blockchain_server/model"
	"time"
)

const TIMES = 10000

// SetSchedule 发布分配表服务
func SetSchedule(c *gin.Context) {
	var req model.ReqSchedule
	var resp model.RespSchedule
	if err := lib.ParseReq(c, "SetSchedule", &req); err != nil {
		return
	}

	//defer this.RunMiddleWare()
	if req.UserAddress == "" {
		glog.Errorln("[api err] 发布者地址为空")
		resp = model.RespSchedule{
			StatusCode: uint32(model.Status_SetSchedueFail),
			Msg:        "20001 发布分配表失败! 发布者地址为空",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	if req.ScheduleName == "" {
		glog.Errorln("[api err] GetScheduleName为空")
		resp = model.RespSchedule{
			StatusCode: uint32(model.Status_SetSchedueFail),
			Msg:        "20002 发布分配表失败! GetScheduleName为空",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	glog.Infoln(lib.Log("api", req.UserAddress, "发布分配表"), req)

	sum := float64(0)
	allByCount := true
	for _, rs := range req.Rss {
		if rs.SubWay == 0 {
			sum += rs.Radios
			allByCount = false
		}
	}

	for i := 0; i < len(req.Rss); i++ {
		for j := i + 1; j < len(req.Rss); j++ {
			if strings.ToLower(req.Rss[i].Job) == strings.ToLower(req.Rss[j].Job) {
				resp = model.RespSchedule{
					StatusCode: uint32(model.Status_SetSchedueFail),
					Msg:        "20003 发布分配表失败! 参与分账的工作必须唯一！",
				}
				c.JSON(http.StatusOK, resp)
				return
			}
		}
	}

	if sum != 100 && !allByCount {
		glog.Errorln("[api err] sum != 100", sum)
		resp = model.RespSchedule{
			StatusCode: uint32(model.Status_SetSchedueFail),
			Msg:        "20004 发布分配表失败! 发布者地址为空 sum(ratio) != 100",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	acc := make([]string, 0)
	rs := make([]float64, 0)
	sw := make([]int64, 0)
	qw := make([]int64, 0)
	rt := make([]int64, 0)
	rss := Sort(req.Rss)

	glog.Infoln("+++++++++++++++after sort", )
	for _, item := range rss {
		glog.Infoln(*item)
	}
	glog.Infoln("+++++++++++++++", )
	for _, v := range rss {
		ra := v.Radios
		if v.SubWay == 1 {
			ra *= TIMES
		} else {
			ra = lib.RadioIn(ra)
		}
		acc = append(acc, v.Job)
		rs = append(rs, ra)
		sw = append(sw, v.SubWay)
		qw = append(qw, v.ResetWay)
		rt = append(rt, v.ResetTime)
	}

	md5Token := ""
	if req.ScheduleName == "" {
		md5Token = "AC" + ":" + lib.CipherStr("AC", time.Now().String())
	} else {
		md5Token = req.ScheduleName
	}

	// 发布分配表
	args := arguments.DistributionArguments{
		SmartAddress: DeployAddress,
		IssueCode:    md5Token,
		SubRoles:     lib.ReParseStrArr(acc),
		Rtaios:       lib.Float64BigIntArr(rs),
		SubWays:      lib.Int64BigIntArr(sw),
		QuotaWays:    lib.Int64BigIntArr(qw),
		ResetTimes:   lib.Int64BigIntArr(rt),
	}

	if len(req.UserKeyStore) > 0 && len(req.KeyString) == 0 {
		req.KeyString = myeth.ParseKeyStoreToString(req.UserKeyStore, req.UserParse)
	}

	hash, err := contracts.SetDistributionRatio(req.KeyString, args)
	if hash == "" {
		glog.Errorln("[api err] 发布分配表失败")
		resp = model.RespSchedule{
			StatusCode: uint32(model.Status_SetSchedueFail),
			Msg:        "20005 发布分配表失败! hash == ",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	if err != nil {
		glog.Errorln("[api err] 发布分配表失败", err)
		resp = model.RespSchedule{
			StatusCode: uint32(model.Status_SetSchedueFail),
			Msg:        "20006 发布分配表失败!" + err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp = model.RespSchedule{
		StatusCode:   uint32(model.Status_Success),
		Hash:         hash,
		ScheduleName: req.ScheduleName,
	}
	glog.Infoln("SetSchedule************** resp:", resp)
	c.JSON(http.StatusOK, resp)
}

func Sort(rs []*model.Rs) []*model.Rs {
	for i := 0; i < len(rs); i++ {
		for j := 0; j < len(rs)-i-1; j++ {
			if rs[j].Level > rs[j+1].Level {
				rs[j], rs[j+1] = rs[j+1], rs[j]
			}
		}
	}
	return rs
}

// 查询排班接口
func GetSchedule(c *gin.Context) {
	//defer this.RunMiddleWare()

	var req model.ReqGetSchedue
	var resp model.RespGetSchedue
	if err := lib.ParseReq(c, "GetSchedule", &req); err != nil {
		return
	}
	glog.Infoln(lib.Log("api", req.UserAddress, "查询排班接口 req"), req)

	jobs, radios, subWays, quos, ts, err := contracts.GetDistributionRatioByCode(DeployAddress, req.ScheduleName)
	if err != nil {
		glog.Errorln("[api err] 查询排班接口失败", err)
		resp = model.RespGetSchedue{
			StatusCode: uint32(model.Status_SetSchedueFail),
			Msg:        "30002 查询排班失败!" + err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	reAccounts := make([]*model.UserAccount, 0)
	_, newAcco, err := contracts.GetSchedulingCxt(DeployAddress, req.ScheduleName)

	if err == nil {
		for ndx := 0; ndx < len(newAcco); ndx++ {
			acc, err := contracts.GetAccountCxt(DeployAddress, common.HexToAddress(newAcco[ndx]))
			if err != nil {
				acc = arguments.AccountArguments{}
			}

			reAccounts = append(reAccounts, &model.UserAccount{
				Address:   newAcco[ndx],
				Name:      acc.Name,
				BankCard:  acc.BankCard,
				WeChat:    acc.WeChat,
				Alipay:    acc.Alipay,
				Telephone: acc.Telephone,
			})
		}
	} else {
		glog.Error(lib.Log("api", req.UserAddress, "GetSchedule GetSchedulingCxt failed,err:"), err.Error(), "req.ScheduleName:", req.ScheduleName)
	}

	sds := make([]*model.Rs, 0)
	for i, _ := range radios {

		ra := float64(radios[i].Int64())
		if subWays[i].Int64() == 1 {
			ra /= TIMES
		} else {
			ra = lib.RadioOut(ra)
		}

		sds = append(sds, &model.Rs{
			Radios:    ra,
			SubWay:    subWays[i].Int64(),
			ResetWay:  quos[i].Int64(),
			ResetTime: ts[i].Int64(),
			Job:       lib.Byte32ToStr(jobs[i]),
		})
	}

	resp = model.RespGetSchedue{
		StatusCode: uint32(model.Status_Success),
		Accounts:   reAccounts,
		Schedules:  sds,
	}
	glog.Infoln(lib.Log("api", "", "查询排班接口 resp"), resp)
	//return resp, nil
	c.JSON(http.StatusOK, resp)
}

// get accounts by schedules
func GetAccountsBySchedule(accos []string) map[string]*model.UserAccount {
	reAccount := make(map[string]*model.UserAccount)
	for _, acc := range accos {
		acc0, err := contracts.GetAccountCxt(DeployAddress, common.HexToAddress(acc))
		if err != nil {
			glog.Infoln(lib.Log("api", acc, "GetAccountBook Error"), err)
			continue
		}
		if _, exist := reAccount[acc]; !exist {
			reAccount[acc] = &model.UserAccount{
				Address:   acc,
				Name:      acc0.Name,
				BankCard:  acc0.BankCard,
				Alipay:    acc0.Alipay,
				WeChat:    acc0.WeChat,
				Telephone: acc0.Telephone,
			}
		}
	}
	return reAccount
}

// NewScheduleId
func NewScheduleId(c *gin.Context) {
	var req model.ReqNewScheduleId
	var resp model.RespNewScheduleId

	if err := lib.ParseReq(c, "NewScheduleId", &req); err != nil {
		return
	}
	glog.Infoln("NewScheduleId************", req)
	md5Token := req.Index + ":" + lib.CipherStr(req.Index, time.Now().String())
	if md5Token == "" {
		resp = model.RespNewScheduleId{
			StatusCode: uint32(model.Status_SetSchedueFail),
			Msg:        "30023 新建排班id失败!",
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp = model.RespNewScheduleId{
		StatusCode:   0,
		ScheduleName: md5Token,
	}
	c.JSON(http.StatusOK, resp)
}

// 2.5.重设按总量分账
func ResetQuo(c *gin.Context) {
	var req model.ReqResetQuo
	var resp model.RespResetQuo
	if err := lib.ParseReq(c, "ResetQuo", &req); err != nil {
		return
	}
	glog.Infoln("Beginning ResetQuo,subCode:", req.ScheduleId,"userAddr:",req.UserAddr)
	hash, err := contracts.ResetSubCodeQuotaData(req.KeyStore, DeployAddress, req.ScheduleId, req.UserAddr)
	if hash == "" {
		resp = model.RespResetQuo{
			Flag: false,
			Msg:  "ResetQuo hash==nil",
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	if err != nil {
		resp = model.RespResetQuo{
			Flag: false,
			Msg:  err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp = model.RespResetQuo{
		Hash: hash,
		Flag: true,
	}
	glog.Infoln("ResetQuo Success,userAddr:", req.UserAddr , "subCode:", req.ScheduleId ,"hash:", hash)
	c.JSON(http.StatusOK, resp)
}

// 2.6.查询按总量分账
func GetQuo(c *gin.Context) {
	var req model.ReqGetQuo
	if err := lib.ParseReq(c, "GetQuo", &req); err != nil {
		return
	}
	var resp model.RespGetQuo
	money, err := contracts.GetSubCodeQuotaData(DeployAddress, req.SubCode, req.UserAddr)

	if err != nil {
		resp = model.RespGetQuo{
			StatusCode: uint32(model.Status_GetSchedueFail),
			Msg:        err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	resp = model.RespGetQuo{
		StatusCode: uint32(model.Status_Success),
		Money:      lib.MoneyOut(float64(money)),
	}
	glog.Infoln(" GetQuo******************resp:", resp)
	c.JSON(http.StatusOK, resp)
}

// 2.7.发布排班  <--（客户端）v2新增
func SetPaiBan(c *gin.Context) {
	var req model.ReqSetPaiBan
	var resp model.RespSetPaiBan
	if err := lib.ParseReq(c, "SetPaiBan", &req); err != nil {
		return
	}
	ro := make([][32]byte, 0)
	jo := make([]string, 0)
	sid := req.SubCode
	for _, v := range req.PaiBans {
		ro = append(ro, lib.StrToByte32(v.JobName))
		jo = append(jo, v.UserAddress)
	}
	//如果没有传输keyString,那么自己编码
	if len(req.UserKeyStore) > 0 && len(req.KeyString) == 0 {
		req.KeyString = myeth.ParseKeyStoreToString(req.UserKeyStore, req.UserParse)
	}

	hash, err := contracts.IssueScheduling(req.KeyString, DeployAddress, sid, ro, jo)
	if err != nil {
		resp = model.RespSetPaiBan{
			StatusCode: uint32(40001),
			Msg:        err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	glog.Infoln("hash", hash)
	resp = model.RespSetPaiBan{
		StatusCode: 0,
		Hash:       hash,
	}
	resp.Msg = "Success"
	glog.Infoln("  SetPaiBan****************resp:", resp)
	c.JSON(http.StatusOK, resp)
}

// 2.8.查询排班  <--（客户端）v2新增
func GetPaiBan(c *gin.Context) {
	var req model.ReqGetPaiBan
	var resp model.RespGetPaiBan
	if err := lib.ParseReq(c, "GetPaiBan", &req); err != nil {
		resp = model.RespGetPaiBan{
			StatusCode: uint32(40001),
			Msg:        err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	// 传入参数，合约地址和分配表编号
	roles, accounts, err := contracts.GetSchedulingCxt(DeployAddress, req.SubCode)
	if err != nil {
		resp = model.RespGetPaiBan{
			StatusCode: uint32(40001),
			Msg:        err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	resp = model.RespGetPaiBan{
		StatusCode:   0,
		Roles:        roles,
		AddressArray: accounts,
	}
	c.JSON(http.StatusOK, resp)
	return
}

// 设置每个账户的已分配定额数
func SetSubCodeQuotaData(c *gin.Context) {
	var req model.ReqSetQuota
	if err := lib.ParseReq(c, "GetPaiBan", &req); err != nil {
		resp := model.RespCommonSet{
			StatusCode: uint32(40001),
			Msg:        err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	if len(req.UserKeyStore) > 0 && len(req.KeyString) == 0 {
		req.KeyString = myeth.ParseKeyStoreToString(req.UserKeyStore, req.UserParse)
	}

	hash, err := contracts.SetSubCodeQuotaData(req.KeyString, DeployAddress, req.SubCode, req.UserAddress, req.SetNumber)

	if err != nil {
		resp := model.RespCommonSet{
			StatusCode: uint32(40001),
			Msg:        err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	glog.Infoln("SetSubCodeQuotaData hash", hash)
	resp := model.RespCommonSet{
		StatusCode: 0,
		Msg:        hash,
	}
	glog.Infoln("  SetSubCodeQuotaData****************resp:", resp)
	c.JSON(http.StatusOK, resp)
}

// 修改财务平台的付款账户地址
func ChanngeSmartPayer(c *gin.Context) {
	var req model.ReqChangePayer
	if err := lib.ParseReq(c, "ChanngeSmartPayer", &req); err != nil {
		resp := model.RespCommonSet{
			StatusCode: uint32(40001),
			Msg:        err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}
	if len(req.UserKeyStore) > 0 && len(req.KeyString) == 0 {
		req.KeyString = myeth.ParseKeyStoreToString(req.UserKeyStore, req.UserParse)
	}

	hash, err := contracts.ChangeSamrtPayer(req.KeyString, DeployAddress, req.PayerAddress)

	if err != nil {
		resp := model.RespCommonSet{
			StatusCode: uint32(40001),
			Msg:        err.Error(),
		}
		c.JSON(http.StatusOK, resp)
		return
	}

	glog.Infoln("ChanngeSmartPayer hash", hash, "PayerAddress:", req.PayerAddress)
	resp := model.RespCommonSet{
		StatusCode: 0,
		Msg:        hash,
	}
	glog.Infoln("  ChanngeSmartPayer****************resp:", resp)
	c.JSON(http.StatusOK, resp)
}
