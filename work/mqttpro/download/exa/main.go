
package main

import (
	"log"
	"net/http"
	"github.com/feng/alg/mqttpro/download"
)
func main() {
	http.HandleFunc("/log/",download.DealStaticFiles)
	http.HandleFunc("/list",download.ListFunc)
	err :=http.ListenAndServe(":8000",nil)
	if err != nil {
		log.Fatal(err)
	}
}