
package main

import (
	"os"
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
	http.HandleFunc("/login", login)
	http.HandleFunc("/test", test)

	http.HandleFunc("/css/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("css**************",r.URL.Path)
		//w.Header().Set("Content-Type", "text/css")当css显示有问题时试试这个
		file:= "./view/" + r.URL.Path[1:]
		f,err := os.Open(file)
		if err!= nil {
			log.Fatal(err)
		}
		defer f.Close()
		http.ServeFile(w,r,file)
	})
	http.HandleFunc("/js/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("js**************",r.URL.Path)
		http.ServeFile(w,r,"./view/"+r.URL.Path[1:])
	})
	
	http.HandleFunc("/images/",func(w http.ResponseWriter, r *http.Request){
		fmt.Println("images**************",r.URL.Path)
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
		t,err := template.ParseFiles("./view/login.html")
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
	fmt.Println("a*******************a")
	t, err := template.ParseFiles("./view/rendered-index.html")
	if err != nil {
		fmt.Println(err)
		return
	}
	t.Execute(w, nil)
}

