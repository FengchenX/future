/*--------------------------------------------------------------
*  package: 注册相关的服务
*  time：   2018/04/18
*-------------------------------------------------------------*/
package apiserver

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"znfz/server/config"
	"znfz/server/lib"
	"znfz/server/protocol"
	"znfz/server/token-contract/accmanager"
)

// 用户注册服务
func (this *ApiService) Register(ctx context.Context, req *protocol.ReqRegister) (*protocol.RespRegister, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", "", "用户注册"), req)
	if req.GetPassWord() == "" {
		glog.Errorln(lib.Log("api err", "", "注册密码不可为空"), req)
		return &protocol.RespRegister{
			StatusCode: uint32(protocol.Status_RegistError),
		}, nil
	}

	addr, desp, err := accmanager.NewAccount(req.GetPassWord()) // 使用秘钥注册账户
	if err != nil {
		glog.Errorln(lib.Log("api err", "", "使用秘钥注册账户"), err)
		return &protocol.RespRegister{
			StatusCode: uint32(protocol.Status_RegistError),
		}, err
	}

	resp := &protocol.RespRegister{
		StatusCode:      uint32(protocol.Status_Success),
		PassWord:        req.GetPassWord(),
		AccountDescribe: desp,
		UserAddress:     addr,
	}
	glog.Infoln(lib.Log("api", addr, "用户注册"), resp)
	return resp, err
}

// 用户绑定服务
func (this *ApiService) Bind(ctx context.Context, req *protocol.ReqBand) (*protocol.RespBand, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetName(), "user banding"), this.AccountAddress, req.GetAccountDescribe())
	if req.GetUserAddress() == "" {
		glog.Errorln(lib.Log("api err", req.GetUserAddress(), "绑定关键数据不可为空"), req)
		return &protocol.RespBand{
			StatusCode: uint32(protocol.Status_RegistError),
		}, nil
	}

	// 绑定账户信息
	hash, err := accmanager.NewAccountAdd(req.GetAccountDescribe(), req.GetPassWord(), this.AccountAddress,
		accmanager.Account{
			AccountAddr: req.GetUserAddress(),
			Name:        req.GetName(),
			Password:    req.GetPassWord(),
			Telephone:   req.GetPhone(),
		})
	if hash == "" || err != nil {
		glog.Errorln(lib.Log("api err", req.GetUserAddress(), "用户绑定"), hash, err)
		return &protocol.RespBand{
			StatusCode: uint32(protocol.Status_RegistError),
		}, err
	}
	resp := &protocol.RespBand{
		StatusCode: uint32(protocol.Status_Success),
	}
	glog.Infoln(lib.Log("api", req.GetName(), "用户绑定回包"), resp.GetStatusCode())
	return resp, err
}

// 查询用户注册信息
func (this *ApiService) GetAccount(ctx context.Context, req *protocol.ReqCheckAccount) (*protocol.RespCheckAccount, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetUserAddress(), " 查询用户注册信息"), req)
	f, err := accmanager.GetAccountByAddr(this.AccountAddress, req.GetUserAddress())
	if err != nil {
		glog.Errorln(lib.Log("api err", req.GetUserAddress(), "查询用户注册信息"), err)
		return &protocol.RespCheckAccount{
			StatusCode: uint32(protocol.Status_LoginFail)}, err
	}
	if f.Password == "" || f.Name == "" {
		glog.Errorln(lib.Log("api err", req.GetUserAddress(), "查询用户注册"), "未查询到该玩家")
		return &protocol.RespCheckAccount{
			StatusCode: uint32(protocol.Status_LoginFail),
		}, nil
	}
	resp := &protocol.RespCheckAccount{
		StatusCode: uint32(protocol.Status_Success),
		Name:       f.Name,
		PassWord:   f.Password,
		Phone:      f.Telephone,
	}
	glog.Infoln(lib.Log("api", req.GetUserAddress(), " 查询注册"), resp)
	return resp, err
}

// 查询登录信息
func (this *ApiService) GetBind(ctx context.Context, req *protocol.ReqLogin) (*protocol.RespLogin, error) {
	defer this.RunMiddleWare()

	glog.Infoln(lib.Log("api", req.GetPhone(), " 查询登录信息"), req)
	accos, addrs, err := accmanager.GetAccountByTel(this.AccountAddress, req.GetPhone())

	if err != nil {
		glog.Errorln(lib.Log("api err", "", "查询登录信息"), err)
		return &protocol.RespLogin{
			StatusCode: uint32(protocol.Status_LoginFail)}, err
	}

	resp := &protocol.RespLogin{}
	for i, acco := range accos {
		l := &protocol.Login{
			StatusCode:      uint32(protocol.Status_Success),
			Name:            acco.Name,
			PassWord:        acco.Password,
			Phone:           acco.Telephone,
			Address:         addrs[i],
			AccountDescribe: acco.AccountDescribe,
		}
		resp.Logins = append(resp.Logins, l)
	}
	resp.StatusCode = uint32(protocol.Status_Success)
	glog.Infoln(lib.Log("api", req.GetPhone(), " 登录接口"), resp)
	return resp, err
}

// 获取以太币余额
func (this *ApiService) GetEthBalance(ctx context.Context, req *protocol.ReqGetEthBalance) (*protocol.RespGetEthBalance, error) {
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "获取以太币余额"))
	money, err := accmanager.GetBalance(req.GetUserAddress(), config.Opts().EthAddress)
	if err != nil {
		glog.Errorln(lib.Log("api err", req.GetUserAddress(), "获取以太币余额失败"), err)
	}

	resp := &protocol.RespGetEthBalance{
		StatusCode: uint32(protocol.Status_Success),
		Balance:    money,
	}

	glog.Infoln(lib.Log("api", req.GetUserAddress(), "获取以太币余额"), resp)
	return resp, nil
}
