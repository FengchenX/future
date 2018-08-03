package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Id string
	Phone []string
}
func main() {

	user:=User{
		Id:"123",
		Phone:[]string{"A","B","C","D"},
	}

	bytes,err:=json.Marshal(&user)
	if err!=nil{
		panic(err)
	}
	user1:=User{}
	json.Unmarshal(bytes,&user1)

	fmt.Println("user1",user1)

	var A map[string]string

	fmt.Println(A["a"],"aaa")
}
