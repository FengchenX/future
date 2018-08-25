package service

import (
	"github.com/sirupsen/logrus"
)

// GetEthBalance 获取eth 币
func (AppSvc) GetEthBalance(userAddress string) (uint32, string, string) {
	var req struct {
		UserAddress string
	}
	var resp struct {
		StatusCode uint32
		Msg        string
		Balance    string
	}
	req.UserAddress = userAddress
	logrus.Infoln("GetEthBalance*****************req:", req)	
	if err := doPost("/getethbalance", req, &resp); err != nil {
		resp.Msg = err.Error()
		resp.StatusCode = 10001 //todo 
	}
	logrus.Infoln("GetEthBalance*****************resp:", resp)
	return resp.StatusCode, resp.Msg, resp.Balance
}