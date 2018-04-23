
package main

import (
	//"io/ioutil"
	"net/url"
	"log"
	"net/http"

)

func main() {
	//Get方法
	resp, err := http.Get("http://localhost:8000/")
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	//body, err := ioutil.ReadAll(resp.Body)

	//_, err = http.PostForm("http://localhost:8000/PostForm",
	//url.Values{"key": {"Value"},"id": {"123"}})

	var vs url.Values
	vs = make(url.Values)
	vs.Add("key","Rust")
	vs.Add("id","go")
	_, err = http.PostForm("http://localhost:8000/PostForm",vs);
}