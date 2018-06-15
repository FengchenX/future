package main

import (
	"flag"
	"fmt"
)

func main() {
	flag.Parse()        //使用前必须要，，不然只会是默认值
	fmt.Println(*in)    //使用命令 go run main.go -path uio,输出 uio
	fmt.Println(*count) //使用命名 go run main.go -path uio -count 56,输出 uio  56
}

var in = flag.String("path", ".", "Use -path <filesource>")
var count = flag.Int("count", 10, "Use -count <counters>")
