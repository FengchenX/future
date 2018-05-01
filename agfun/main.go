
package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
)


func main() {
	
	http.HandleFunc("/index", agfun)
	http.HandleFunc("/login", login)

	


	
	http.HandleFunc("/statics/",func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"./view/"+r.URL.Path[1:])
	})
	
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		log.Fatal(err)
	}
}




/**网站首页*/
func agfun(w http.ResponseWriter, r *http.Request) {
	fmt.Println("visit first index")	
	t, _ := template.ParseFiles("./view/index.html")
	t.Execute(w, nil)
}

//登录处理器
func login(w http.ResponseWriter, r *http.Request) {
	r.ParseForm()
	if r.Method=="GET" {
		t,err := template.ParseFiles("view/login.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(w, nil)
	} else {
		//验证密码登录部分
		

	}
}

