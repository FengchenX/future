package log

import (
	"time"

	"github.com/feng/future/go-kit/agfun/main-service/protocol"
	"github.com/sirupsen/logrus"
)

func (mw logmw) CreateAccount(req protocol.CreateAccountReq) (resp protocol.Resp, err error) {
	defer func(begin time.Time) {
		logrus.Infof(
			"method:%s, input:%+v, output: %+v, err: %v, took:%d",
			"CreateAccount", req, resp, err, time.Since(begin),
		)
	}(time.Now())
	resp, err = mw.AppService.CreateAccount(req)
	return
}

func (mw logmw) Login(req protocol.LoginReq) (resp protocol.Resp, err error) {
	defer func(begin time.Time) {
		logrus.Infof(
			"method:%s, input:%+v, output: %+v, err: %v, took:%d",
			"Login", req, resp, err, time.Since(begin),
		)
	}(time.Now())
	resp, err = mw.AppService.Login(req)
	return
}

func (mw logmw) Account(req protocol.AccountReq) (resp protocol.Resp, err error) {
	defer func(begin time.Time) {
		logrus.Infof(
			"method:%s, input:%+v, output: %+v, err: %v, took:%d",
			"Account", req, resp, err, time.Since(begin),
		)
	}(time.Now())
	resp, err = mw.AppService.Account(req)
	return
}
