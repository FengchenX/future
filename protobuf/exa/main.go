package main

import (
	"fmt"
	"os"
	"io/ioutil"
	"github.com/golang/protobuf/proto"
	"github.com/feng/future/protobuf"
)

func main() {
	write()
	read()
}

func write() {
	p1:=&protobuf.Person{
		Id: 1,
		Name: "小张",
		Phones: []*protobuf.Phone{
			{protobuf.PhoneType_HOME, "111111111111"},
			{protobuf.PhoneType_WORK, "222222222222"},
		},
	}
	p2 := &protobuf.Person{
		Id: 2,
		Name: "小王",
		Phones: []*protobuf.Phone{
			{protobuf.PhoneType_HOME, "333333333333"},
			{protobuf.PhoneType_WORK, "444444444444"},
		},
	}
	//创建地址簿
	book := &protobuf.ContactBook{}
	book.Persons = append(book.Persons, p1)
	book.Persons = append(book.Persons, p2)

	//编码数据
	data, _ := proto.Marshal(book)
	//把数据写入文件
	ioutil.WriteFile("./test.txt",data,os.ModePerm)
}
func read() {
	//读取数据
	data,_:=ioutil.ReadFile("./test.txt")
	book := &protobuf.ContactBook{}
	//解码数据
	proto.Unmarshal(data,book)
	for _, v := range book.Persons {
		fmt.Println(v.Id, v.Name)
		for _, vv := range v.Phones {
			fmt.Println(vv.Type, vv.Number)
		}
	}

}