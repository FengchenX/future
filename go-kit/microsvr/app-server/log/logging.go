package log

import (
	"time"
	"github.com/sirupsen/logrus"
	"github.com/feng/future/go-kit/microsvr/app-server/service"
	"github.com/feng/future/go-kit/microsvr/app-server/model"
)

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
