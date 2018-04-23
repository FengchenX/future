
package main

import (
	"strconv"
	"fmt"

)

func main() {
	fmt.Println(strconv.ParseBool("1"))//true
	fmt.Println(strconv.ParseBool("t"))
	fmt.Println(strconv.ParseBool("T"))
	fmt.Println(strconv.ParseBool("true"))
	fmt.Println(strconv.ParseBool("True"))

	fmt.Println(strconv.FormatBool(0<1))
	fmt.Println(strconv.FormatBool(0>1))

	rst := make([]byte, 0)
	rst = strconv.AppendBool(rst, 0<1)
	fmt.Printf("%s\n", rst)
	rst = strconv.AppendBool(rst, 0> 1)	
	fmt.Printf("%s\n",rst)

	s := "0.12345678901234567890"
	f, err := strconv.ParseFloat(s,32)
	fmt.Println(f, err)
	fmt.Println(float32(f), err)
	f, err = strconv.ParseFloat(s, 64)
	fmt.Println(f, err)

	//base: 进位制， bitSize:指定整数类型int8,int32,in64, 0 代表int
	fmt.Println(strconv.ParseInt("123", 10, 8))
	fmt.Println(strconv.ParseInt("12345", 10, 8))// value out of range
	fmt.Println(strconv.ParseInt("2147483647", 10, 0))//解析为int类型
	fmt.Println(strconv.ParseInt("0xFF", 16, 0))//invalid syntax 错误语法
	fmt.Println(strconv.ParseInt("FF", 16, 0))//255
	fmt.Println(strconv.ParseInt("0xFF", 0, 0))//255

	fmt.Println(strconv.Atoi("2147483647"))


}