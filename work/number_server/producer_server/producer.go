package main

import (
	"flag"
	"net/http"
	"sub_account_service/number_server/config"
	"sub_account_service/number_server/routers/produce"
	"sub_account_service/number_server/routers/produce/api"
)

func main() {
	flag.Parse()
	//接受消息
	err := api.Setup()

	if err != nil {
		panic(err)
	}
	router := produce.InitRouter()

	err = http.ListenAndServe(config.Opts().Producer_Server_Http, router)
	if err != nil {
		panic(err)
	}
}
