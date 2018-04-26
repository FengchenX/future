package main

import (
	"net/http"
	"log"
	"html/template"
)
//todo
func main() {
	http.HandleFunc("/f1",f1)
	http.HandleFunc("/f2",f2)
	err := http.ListenAndServe(":8080",nil)
	if err != nil {
		log.Fatal(err)
	}
}

func f1(w http.ResponseWriter, r *http.Request) {
	temp,err := template.ParseFiles("./a.html")
	if err!= nil {
		log.Fatal(err)
	}
	temp.Execute(w,"fuio")

}

func f2(w http.ResponseWriter, r *http.Request) {
	temp, err := template.ParseFiles("./b.html")
	if err != nil {
		log.Fatal(err)
	}
	var p = PP{10,"uio",true}
	temp.Execute(w,p)
}

//要大写不然导不出去
type PP struct {
	A int
	B string
	C bool
}