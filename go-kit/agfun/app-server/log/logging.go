package log

import (
	"time"
	"github.com/sirupsen/logrus"
	"github.com/feng/future/go-kit/agfun/app-server/service"
	"github.com/feng/future/go-kit/agfun/app-server/model"
)

//LoggingMiddleware 日志中间件
func LoggingMiddleware() service.SvcMiddleware {
	return func(next service.AppService) service.AppService {
		return logmw{next}
	}
}
type logmw struct {
	service.AppService
}

func (mw logmw) GetAccount(userAddr string) (status uint32, msg string, userAccount model.UserAccount) {
	defer func(begin time.Time) {
		logrus.Infoln(
			"method", "GetAccount",
			"input", userAddr,
			"output", status, msg, userAccount,
			"took", time.Since(begin),
		)
	}(time.Now())
	status, msg, userAccount = mw.AppService.GetAccount(userAddr)
	return
}

func (mw logmw) SetAccount(userKeyStore, userParse, keyString string, userAccount model.UserAccount) (status uint32, msg string) {
	defer func(begin time.Time) {
		logrus.Infoln(
			"method", "SetAccount",
			"input", userKeyStore, userParse, keyString,
			"output", status, msg,
			"took", time.Since(begin),
		)
	}(time.Now())
	status, msg = mw.AppService.SetAccount(userKeyStore, userParse, keyString, userAccount)
	return
}

