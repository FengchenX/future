package apiserver

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"znfz/server/lib"
	"znfz/server/protocol"
	"znfz/server/token-contract/accmanager"
)

// 5.1.添加用人白名单
func (this *ApiService) AddWhiteList(ctx context.Context, req *protocol.ReqAddWhite) (*protocol.RespAddWhite, error) {
	defer this.RunMiddleWare()
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "添加用人白名单"), req)

	_, addr, err := accmanager.GetAccountByTel(this.AccountAddress, req.GetPhone())
	if err != nil && len(addr) != 0 {
		return &protocol.RespAddWhite{
			StatusCode: uint32(protocol.Status_RoleNotFitFail),
		}, nil
	}

	str, err := accmanager.AddEmployStaff(req.GetAccountDescribe(), req.GetPassWord(), this.AccountAddress, addr[0])
	if str == "" && err != nil {
		return &protocol.RespAddWhite{
			StatusCode: uint32(protocol.Status_RoleNotFitFail),
		}, nil
	}
	return &protocol.RespAddWhite{
		StatusCode: uint32(protocol.Status_Success),
	}, nil
}

// 5.2.根据账户地址查询用人白名单
func (this *ApiService) GetWhiteList(ctx context.Context, req *protocol.ReqGetWhite) (*protocol.RespGetWhite, error) {
	defer this.RunMiddleWare()
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "根据账户地址查询用人白名单"), req)

	findAccount, str, err := accmanager.GetEmployStaffs(this.AccountAddress, req.GetUserAddress())

	friends := make([]*protocol.Friend, 0)

	if len(findAccount) == len(str) {
		for i, v := range findAccount {
			friend := &protocol.Friend{
				Name:    v.Name,
				Phone:   v.Telephone,
				Address: str[i],
			}
			friends = append(friends, friend)
		}
	}

	if err != nil {
		return &protocol.RespGetWhite{
			StatusCode: uint32(protocol.Status_RoleNotFitFail),
		}, nil
	}
	return &protocol.RespGetWhite{
		StatusCode: uint32(protocol.Status_Success),
		Friends:    friends,
	}, nil
}

// 5.3.删除用人白名单
func (this *ApiService) DelWhiteList(ctx context.Context, req *protocol.ReqDelWhite) (*protocol.RespDelWhite, error) {
	defer this.RunMiddleWare()
	glog.Infoln(lib.Log("api", req.GetUserAddress(), "删除用人白名单"), req)

	str, err := accmanager.DelEmployStaff(req.GetAccountDescribe(), req.GetPassWord(), this.AccountAddress, req.GetAdderAddress())
	if str == "" && err != nil {
		return &protocol.RespDelWhite{
			StatusCode: uint32(protocol.Status_RoleNotFitFail),
		}, nil
	}
	return &protocol.RespDelWhite{
		StatusCode: uint32(protocol.Status_Success),
	}, nil
}
