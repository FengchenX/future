package control

import (
	"net/http"
	"fmt"
	"html/template"
	//引用memory包
	_ "github.com/feng/future/agfun/memory"
	"github.com/feng/future/agfun/session"
)

var globalSessions *session.Manager

func init(){
	globalSessions, _ = session.NewSessionManager("memory", "gosessionid", 3600)
	
	go globalSessions.GC()
}

//Agfun 网站首页
func Agfun(w http.ResponseWriter, r *http.Request) {
	fmt.Println("visit first index")	
	t, _ := template.ParseFiles("./template/index.html")
	t.Execute(w, nil)
}

//Login 登录处理器
func Login(w http.ResponseWriter, r *http.Request) {
	fmt.Println("login******************login")
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
		username := r.Form["username"][0]
		password := r.Form["password"][0]
		fmt.Println(username,password)

		//验证密码登录部分通过之后才set
		sess.Set("username", r.Form["username"])

	}
}