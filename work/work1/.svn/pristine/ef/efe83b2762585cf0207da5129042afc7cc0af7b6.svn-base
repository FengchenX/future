package subaccount

import (
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/golang/glog"
	"math/big"
	"strings"
	"znfz/server/arguments"
	"znfz/server/config"
	"znfz/server/lib"
	"znfz/server/token-contract/utils"
)

type scheduleStaff struct {
	StaffAddr     string   //  当前申请者
	JobId         *big.Int // jobid
	Ratio         *big.Int // 分账比例
	FatherAddress string   // 发布者父地址
	Role          [32]byte // 角色
}

type scheduleDesired struct {
	Issuer         string          // 发布者地址
	IssueRatio     *big.Int        // 发布者初始比例
	AllCount       *big.Int        // 发布总人数 包括自己
	JobIds         []*big.Int      // 自增id
	IssuerDesireds []issuerDesired // 单个职位发布者信息
}

// 单个
type issuerDesired struct {
	Role   [32]byte
	Count  *big.Int
	Ratio  *big.Int
	JobId  *big.Int
	Whites string // 对应的可申请人的账户地址
}

// 排班发布
func PublishScheduleing(pargs arguments.ScheduleArguments) (string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	auth, err := utils.GetEthAuth(pargs.OperationKeyStore, pargs.OperationPassWord)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}

	token, err := NewToken(common.HexToAddress(pargs.SmartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}
	whitelists := []common.Address{}
	for wdx := 0; wdx < len(pargs.Whitelists); wdx++ {
		whitelists = append(whitelists, common.HexToAddress(pargs.Whitelists[wdx]))
	}

	tx, err := token.Scheduleing(auth, common.HexToAddress(pargs.FatherAddress), pargs.IssueRatio, whitelists, pargs.JobIds, pargs.Roles, pargs.Counts, pargs.Ratio)
	if err != nil {
		glog.Errorln("Failed to request sub account PublishScheduleing: ", err.Error())
		return "", err
	}
	glog.Infoln("sub account PublishScheduleing success pending:", tx.Hash().String(), "smartAddress:", pargs.SmartAddress)
	return tx.Hash().String(), nil
}

// 查询排班发布信息 --- 查询这个合约上发布的所有排班信息
func GetScheduleingCxt(smartAddress string, jobIds []*big.Int) ([]scheduleDesired, string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return []scheduleDesired{}, "", err
	}
	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return []scheduleDesired{}, "", err
	}

	publish_addrs, _ := token.GetAllpublishAddrs(nil)
	desiredCxts := []scheduleDesired{}

	for rdx := 0; rdx < len(publish_addrs); rdx++ {
		desired, _ := token.ScheduleDesireds(nil, publish_addrs[rdx])
		desiredCxt := scheduleDesired{}
		desiredCxt.Issuer = strings.ToLower(desired.Issuer.String())
		desiredCxt.IssueRatio = desired.IssueRatio
		desiredCxt.AllCount = desired.Count
		desiredCxt.JobIds = jobIds

		if len(jobIds) == 0 {
			jobIds, _ = token.GetScheduledJobIds(nil, desired.Issuer)
			glog.Infoln(lib.Log("subacc", "", "GetScheduleingCxt"), "GetScheduledJobIds after jobIds:", jobIds, "Issuer:", desired.Issuer.String())
		}

		for rdx := 0; rdx < len(jobIds); rdx++ {
			role, count, ratio, white, err := token.FindScheduledIssue(nil, desired.Issuer, jobIds[rdx])
			if err != nil {
				glog.Error("GetScheduleingCxt for FindScheduledIssu err:", err.Error(), "desired.Issuer:", desired.Issuer.String())
				continue
			}
			glog.Infoln("GetScheduleingCxt jobIds:", jobIds, "role:", role, "count:", count, "ratio:", ratio, "white:", white.String())

			issuerCxt := issuerDesired{}
			issuerCxt.Role = role
			issuerCxt.Count = count
			issuerCxt.Ratio = ratio
			issuerCxt.JobId = jobIds[rdx]
			issuerCxt.Whites = white.String()
			desiredCxt.IssuerDesireds = append(desiredCxt.IssuerDesireds, issuerCxt)
		}

		desiredCxts = append(desiredCxts, desiredCxt)
	}
	symbol, _ := GetTokenSymbol(smartAddress)
	glog.Infoln(lib.Log("subacc", "", "GetScheduleingCxt"), " desiredCxts:", desiredCxts, "smartAddress:", smartAddress, "symbol:", symbol)
	return desiredCxts, symbol, nil
}

// 岗位申请 --- 新增一个参数：fatherAddr 申请职位的发布者账户
func ApplicationJob(operationKey string, operationPhrase string, smartAddress string, staffAddr string, jobId *big.Int, fatherAddr string) (string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	auth, err := utils.GetEthAuth(operationKey, operationPhrase)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}

	tx, err := token.Application(auth, common.HexToAddress(fatherAddr), common.HexToAddress(staffAddr), jobId)
	if err != nil {
		glog.Errorln("Failed to request sub account ApplicationJob: ", err.Error())
		return "", err
	}
	glog.Infoln(lib.Log("subacc", "", "ApplicationJob"), "success pending:", tx.Hash().String(), "smartAddress:", smartAddress)
	return tx.Hash().String(), nil
}

// 判断是否已经申请过岗位
func CheckIsOkApplication(jobAddress string, staffAddr string) (bool, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return false, err
	}

	token, err := NewToken(common.HexToAddress(jobAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return false, err
	}

	staffCxt, err := token.ScheduleStaffs(nil, common.HexToAddress(staffAddr))
	if err != nil || staffCxt.JobId == nil || staffCxt.Ratio == nil {
		glog.Errorln("some of staffCxt is nil", staffCxt)
		return false, errors.New("some of staffCxt is nil")
	}
	if staffCxt.JobId.Int64() > 0 && staffCxt.Ratio.Int64() > 0 {
		return true, nil
	}
	return false, nil
}

// 查询岗位申请信息
func GetApplicationCxt(smartAddress string) ([]scheduleStaff, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return []scheduleStaff{}, err
	}
	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return []scheduleStaff{}, err
	}
	application_address, _ := token.GetScheduleingAddr(nil)

	staffCxts := []scheduleStaff{}
	for adx := 0; adx < len(application_address); adx++ {
		glog.Errorln("GetApplicationCxt for application_address:", application_address[adx].String())

		staffCxt := scheduleStaff{}
		token_staff, err := token.ScheduleStaffs(nil, application_address[adx])
		if err == nil {
			staffCxt.StaffAddr = token_staff.StaffAddr.String()
			staffCxt.Ratio = token_staff.Ratio
			staffCxt.JobId = token_staff.JobId
			staffCxt.Role = token_staff.Role
			staffCxt.FatherAddress = token_staff.FatherAddress.String()
			staffCxts = append(staffCxts, staffCxt)
		}
	}
	glog.Infoln(lib.Log("subacc", "", "GetApplicationCxt"), "staffCxt:", staffCxts, "smartAddress:", smartAddress)
	return staffCxts, nil
}

// 订单记录
func SetOrdersContent(operationKey string, operationPhrase string, smartAddress string, ordersId string, content string) (string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}

	auth, err := utils.GetEthAuth(operationKey, operationPhrase)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}

	tx, err := token.SaveContHash(auth, ordersId, content)
	if err != nil {
		glog.Errorln("Failed to request sub account SetOrdersContent: ", err.Error())
		return "", err
	}
	glog.Infoln(lib.Log("subacc", "", "SetOrdersContent"), " success pending:", tx.Hash().String(), "smartAddress:", smartAddress)
	return tx.Hash().String(), nil

}

// 订单查询
func GetOrdersContent(smartAddress string, ordersId string) (string, *big.Int, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", big.NewInt(0), err
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", big.NewInt(0), err
	}

	content, creatAt, err := token.GetContHash(nil, ordersId)
	if err != nil {
		glog.Errorln("Failed to request sub account GetOrdersContent: ", err.Error())
		return "", big.NewInt(0), err
	}
	glog.Infoln(lib.Log("subacc", "", "GetOrdersContent"), "success content:", content, "creatAt:", creatAt, "smartAddress:", smartAddress)
	return content, creatAt, nil
}

// 每单分账
func SettleAccounts(operationKey string, operationPhrase string, smartAddress string, totalConsume *big.Int) (string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	auth, err := utils.GetEthAuth(operationKey, operationPhrase)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}

	tx, err := token.SettleAccounts(auth, totalConsume)
	if err != nil {
		glog.Errorln("Failed to request sub account SettleAccounts: ", err.Error())
		return "", err
	}
	glog.Infoln(lib.Log("subacc", "", "SettleAccounts"), "success pending:", tx.Hash().String(), "smartAddress:", smartAddress)
	return tx.Hash().String(), nil
}

// 查询账户积分
func GetBalance(smartAddress string, faddress string) (*big.Int, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return big.NewInt(0), err
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return big.NewInt(0), err
	}
	balance, err := token.GetBalance(nil, common.HexToAddress(faddress))
	if err != nil {
		glog.Errorln("Failed to request sub account GetBalance: ", err.Error())
		return big.NewInt(0), err
	}
	glog.Infoln(lib.Log("subacc", "", "GetBalance"), "faddress:", faddress, "success balance:", balance)
	return balance, nil
}

func CheckOwnerIsOk(smartAddress string, accountAddr string) (bool, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return false, err
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return false, err
	}
	owerAddr, err := token.Owner(nil)
	if err != nil {
		glog.Errorln("sub account CheckOwnerIsOk Failed,err: ", err.Error())
		return false, err
	}
	if owerAddr.String() == accountAddr {
		return true, nil
	}
	return false, nil
}

// 获取当前合约的所有订单信息
func GetAllContentHashCxt(smartAddress string) ([]string, error) {
	fallHashCxt := []string{}
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return fallHashCxt, err
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return fallHashCxt, err
	}

	bklen, err := token.GetContentHashKeysLen(nil)
	if err != nil {
		glog.Errorln("Fail to find GetContentHashKeysLen: ", err.Error())
		return fallHashCxt, err
	}
	if bklen == nil {
		glog.Infoln("GetContentHashKeysLen bklen is 0,not data")
		return fallHashCxt, nil
	}

	var klen int64 = bklen.Int64()
	for kdx := 0; kdx < int(klen); kdx++ {
		fhash, err := token.ContentHashKeys(nil, big.NewInt(int64(kdx)))
		if err != nil {
			glog.Errorln("Failed to ContentHashKeys,err: ", err.Error())
			continue
		}

		content, _, err := token.GetContHash(nil, fhash)
		if err != nil {
			glog.Errorln("Failed to request sub account GetOrdersContent: ", err.Error())
			continue
		}
		fallHashCxt = append(fallHashCxt, content)
	}
	glog.Infoln("GetAllContentHashCxt fallHashCxt:", fallHashCxt)
	return fallHashCxt, nil
}

func FindScheduledIssue(smartAddress string, staffAddr string, jobIds []*big.Int) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return
	}

	for rdx := 0; rdx < len(jobIds); rdx++ {
		role, count, ratio, white, err := token.FindScheduledIssue(nil, common.HexToAddress(staffAddr), jobIds[rdx])
		if err != nil {
			continue
		}
		glog.Infoln("FindScheduledIssu jobIds:", jobIds, "role:", role, "count:", count, "ratio:", ratio, "white:", white.String())
	}
	return
}

func GetAllRatio(smartAddress string) int64 {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return 0
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return 0
	}

	ratio, _ := token.GetAllRatio(nil)
	glog.Infoln("GetAllRatio  ratio:", ratio.Int64())
	return ratio.Int64()
}

func GetCompanyCxt(smartAddress string) (string, string, int64) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", "", 0
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", "", 0
	}
	companyPayee, _ := token.CompanyPayee(nil)
	companyRatio, _ := token.CompanyRatio(nil)
	payer, _ := token.Payer(nil)
	glog.Infoln("GetCompanyCxt  companyPayee:", companyPayee.String(), "companyRatio:", companyRatio.Int64(), "Payer:", payer.String())
	return companyPayee.String(), payer.String(), companyRatio.Int64()
}

// 查询备注信息
func GetPostscriptCxt(smartAddress string) string {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return ""
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return ""
	}
	postscript, _ := token.Postscript(nil)
	glog.Infoln("GetPostscriptCxt  postscript:", postscript)
	return postscript
}

func UpdatePostscriptCxt(operationKey string, operationPhrase string, smartAddress string, u_postscript string) (string, error) {
	conn, err := utils.GetEthclient(config.Opts().IpcDir)
	if err != nil {
		glog.Errorln("Failed to connect to the Ethereum client: ", err.Error())
		return "", err
	}
	auth, err := utils.GetEthAuth(operationKey, operationPhrase)
	if err != nil {
		glog.Errorln("Failed to create authorized transactor: ", err.Error())
		return "", err
	}

	token, err := NewToken(common.HexToAddress(smartAddress), conn)
	if err != nil {
		glog.Errorln("Failed to instantiate a Token contract: ", err.Error())
		return "", err
	}

	tx, err := token.UpdatePostscript(auth, u_postscript)
	if err != nil {
		glog.Errorln("Failed to request sub account SettleAccounts: ", err.Error())
		return "", err
	}
	glog.Infoln("UpdatePostscriptCxt  u_postscript:", u_postscript, "hash code:", tx.Hash().String())

	return tx.Hash().String(), nil
}
