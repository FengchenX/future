package apiserver

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"math/big"
	"znfz/server/protocol"
)

var all = []*big.Int{} // 所有员工

// 测试服务
func (this *ApiService) SayHello(ctx context.Context, req *protocol.Req) (*protocol.Resp, error) {
	defer this.RunMiddleWare()

	glog.Infoln("[api 测试服务] receive request from ->>", req.Name)
	return &protocol.Resp{
		Message: req.GetName() + " , rpc请求成功啦！",
	}, nil
}