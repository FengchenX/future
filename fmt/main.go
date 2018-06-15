package main

import (
	"fmt"
	"os"
)

func main() {
	temp := myint{i: 10}
	fmt.Printf("%T\n", temp)
	fmt.Printf("%v\n", temp)
	fmt.Printf("%+v\n", temp)
	fmt.Printf("%#v\n", temp)
	r := fmt.Sprintf("%[2]d %[1]d\n", 11, 22)
	fmt.Println(r)
	r = fmt.Sprintf("%[3]*.[2]*[1]f", 1.0, 2, 6) //最少总体6位两位小数，不够前面用空格表示,等价于fmt.Sprintf("%6.2f", 12.0),
	fmt.Println(r)
	s := "world!"
	fmt.Fprintf(os.Stderr, "hello,%s\n", s)
	fmt.Print(true, 123)   //不同参数间加空格
	fmt.Println(true, 123) //不同参数间加空格并且结尾加换行
	fmt.Print(fmt.Sprintln(false, "nihao", 56, "world"))
	fmt.Println(fmt.Sprintln(false, "nihao", 56, "world"))
}

type myint struct {
	i int
}
