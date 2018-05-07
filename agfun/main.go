
package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
	_ "github.com/feng/future/agfun/memory"
	"github.com/feng/future/agfun/session"
)

var globalSessions *session.Manager

func init(){
	globalSessions, _ = session.NewSessionManager("memory", "gosessionid", 3600)
	
	go globalSessions.GC()
}

func main() {
	http.HandleFunc("/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("URL**************",r.URL.Path)
	})

	http.HandleFunc("/index", agfun)
	http.HandleFunc("/login", login)
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

/**网站首页*/
func agfun(w http.ResponseWriter, r *http.Request) {
	fmt.Println("visit first index")	
	t, _ := template.ParseFiles("./template/index.html")
	t.Execute(w, nil)
}

//登录处理器
func login(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w, r)	
	r.ParseForm()
	if r.Method=="GET" {
		t,err := template.ParseFiles("./template/login.html")
		if err != nil {
			fmt.Println(err)
			return
		}
		t.Execute(w, nil)
	} else {
		//验证密码登录部分通过之后才set
		sess.Set("username", r.Form["username"])

	}
}

func test(w http.ResponseWriter, r *http.Request) {
	sess := globalSessions.SessionStart(w,r)
	sess.Set("username", "feng")
	t, err := template.ParseFiles("./template/rendered-index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}

