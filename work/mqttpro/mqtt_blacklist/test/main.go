
package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/feng/alg/mqttpro/mqtt_blacklist"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:feng@/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	var num = 100
	mqtt_blacklist.InitBlacklist(db, &num)
	http.HandleFunc("/postForm", foo)	
	err = http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func foo(w http.ResponseWriter, r *http.Request) {
	if mqtt_blacklist.DoFilter(r,w) {
		fmt.Println("foo***********此ip在黑名单中")
	}else {
		fmt.Println("foo*************go")
	}
}