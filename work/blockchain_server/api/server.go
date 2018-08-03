package api

import (
	"github.com/golang/glog"
)

// api服务结构体，必须实现了所有service ApiService中的方法
type ApiService struct {
	DeployAddress string
	middleWare    []func()
}

func (this *ApiService) init() {
	this.middleWare = make([]func(), 0)
	this.SetMiddleWare(func() {
		if r := recover(); r != nil {
			glog.Errorln("panic", r)
		}
	})
}

func (this *ApiService) SetMiddleWare(f func()) {
	this.middleWare = append(this.middleWare, f)
}

func (this *ApiService) RunMiddleWare() {
	for _, f := range this.middleWare {
		f()
	}
}
