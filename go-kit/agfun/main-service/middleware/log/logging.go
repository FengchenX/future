package log

import (
	"github.com/feng/future/go-kit/agfun/main-service/service"
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
