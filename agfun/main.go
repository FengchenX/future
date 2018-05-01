
package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
)


func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("URL**************",r.URL.Path)
	})

	http.HandleFunc("/index", agfun)
	http.HandleFunc("/template/login", login)
	http.HandleFunc("/test", test)
	
	
	http.HandleFunc("/static/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("static*************",r.URL.Path)
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
	t, _ := template.ParseFiles("./view/template/index.html")
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

func test(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("view/a.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}

