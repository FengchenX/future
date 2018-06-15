package main

import (
	"github.com/feng/alg/mqttpro/download"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/log/", download.DealStaticFiles)
	http.HandleFunc("/list", download.ListFunc)
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
