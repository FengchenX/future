package main

import (
	"flag"
	"net/http"
	"sub_account_service/number_server/config"
	"sub_account_service/number_server/models"
	"sub_account_service/number_server/routers/query"

)

func main() {

	flag.Parse()

	models.Setup()

	router := query.InitRouter()

	err := http.ListenAndServe(config.Opts().Query_Server_Http, router)

	if err != nil {
		panic(err)
	}
}
