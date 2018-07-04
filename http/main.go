package main

import (
	"github.com/tealeg/xlsx"
	"fmt"
	"log"
	"net/http"
	"strings"
	"os"
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

/**处理下载文件*/
func DealStaticFiles(w http.ResponseWriter, r *http.Request) {

	var xf *xlsx.File
	var sheet *xlsx.Sheet
	var row *xlsx.Row
	var cell *xlsx.Cell

	xf = xlsx.NewFile()
	sheet, err := xf.AddSheet("Sheet1")
	if err != nil {
		fmt.Println(err)
	}
	

	//访问的url路径里包含logDir
	if strings.HasPrefix(r.URL.Path, "/test1") {


		file := "test1.xlsx"
		fmt.Println(file)
		f, err := os.Open(file)
		if err != nil && os.IsNotExist(err) {
			fmt.Fprintln(w, "File not exist")
			return
		}
		defer f.Close()
		http.ServeFile(w, r, file) //将文件内容写到客户端
	} else {
		fmt.Fprintln(w, "Hello world")
	}
}
