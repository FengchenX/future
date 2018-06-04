
package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
	//"github.com/feng/future/agfun/data"
	"github.com/feng/future/agfun/control"
)


func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("URL**************",r.URL.Path)
	})

	http.HandleFunc("/index", control.Agfun)
	http.HandleFunc("/login", control.Login)
	http.HandleFunc("/test", test)

	//文件系统的路由
	http.Handle("/css/", http.FileServer(http.Dir("./template")))
	http.Handle("/js/", http.FileServer(http.Dir("./template")))
	http.Handle("/images/", http.FileServer(http.Dir("./template")))
	
	
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		log.Fatal(err)
	}
}


func test(w http.ResponseWriter, r *http.Request) {
	//sess := globalSessions.SessionStart(w,r)
	//sess.Set("username", "feng")
	t, err := template.ParseFiles("./template/rendered-index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}

