
package main

import (
	"log"
	"fmt"
	"net/http"
	"html/template"
)

func main() {
	http.HandleFunc("/index", agfun)
	err := http.ListenAndServe(":8000",nil)
	if err != nil {
		log.Fatal(err)
	}
}

/**网站首页*/
func agfun(w http.ResponseWriter, r *http.Request) {
	fmt.Println("visit first index")	
	t, _ := template.ParseFiles("./view/agfun.html")
	t.Execute(w, nil)
}