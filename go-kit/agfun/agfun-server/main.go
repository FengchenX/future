package main

import (
	"fmt"
	kitlogmw "github.com/feng/future/go-kit/agfun/agfun-server/middleware/log"
	"github.com/feng/future/go-kit/agfun/agfun-server/router"
	"github.com/feng/future/go-kit/agfun/agfun-server/service"
	"github.com/sirupsen/logrus"
	"net/http"
	"os"
)

func init() {
	// Log as JSON instead of the default ASCII formatter.
	logrus.SetFormatter(&logrus.TextFormatter{})

	// Output to stdout instead of the default stderr
	// Can be any io.Writer, see below for File example
	logrus.SetOutput(os.Stdout)
	//go log.RenameLogFile()

	// Only log the warning severity or above.
	//logrus.SetLevel(logrus.WarnLevel)
	logrus.SetLevel(logrus.InfoLevel)
}

func main() {

	var svc service.AppService
	svc = &service.AppSvc{}
	svc = kitlogmw.LoggingMiddleware()(svc)
	router.Start(svc)

}

func consulCheck(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "consulCheck")
}
