
package main

import (
	"log"
	"net/http"
	"net/url"

)

func main() {
	for i := 0; i<1000; i++ {
		var value url.Values
		value = make (url.Values)
		value.Add("id","news")
		_, err := http.PostForm("http://127.0.0.1:8080/postForm", value)
		if err != nil {
			log.Fatal(err)
		}
	}
}