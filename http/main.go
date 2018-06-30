package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", foo)
	http.HandleFunc("/PostForm", myPostForm)
	http.HandleFunc("/api/heroes", func(w http.ResponseWriter, r *http.Request){
		fmt.Println("hello angular")
		w.Write([]byte("angular feng test"))
	})
	err := http.ListenAndServe(":8000", nil)
	if err != nil {
		log.Fatal(err)
	}
}
func foo(w http.ResponseWriter, r *http.Request) {
	fmt.Println("feng")
	//get 的r.form.encode()为空
	//r.ParseForm()
	//fmt.Println(r.Form.Encode())
}
func myPostForm(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Fatal(err)
	}
	key := r.Form.Get("key") //Value
	id := r.Form.Get("id")   //123
	fmt.Println(key, id)
	fmt.Println(r.Form.Encode())
}
