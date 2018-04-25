package main

import (
	"fmt"

)

func main() {
	//slice()
	//slice2()
	//Map()
	//slice3()
	//slice4()
	//slice6()
}

//slice 一些需要注意的坑
func f1(x []int) {
	for i:=0;i<len(x);i++ {
		x[i] = x[i] + 3
	}
	for i:=0;i<10;i++{
		x=append(x,i)
	}
}

func slice() {
	var a = []int{1,2,3}
	f1(a)
	fmt.Println(a)  //[4 5 6]  已存在的值修改了，，新append的并没有传出来 append是新开辟内存空间所以会出现这个问题
}

func slice2() {
	var a = []int{1,2,3}
	f2(&a)
	fmt.Println(a)  //[4 5 6 0 1 2 3 4 5 6 7 8 9] //彻底改了

}

func f2(x *[]int) {
	for i:=0;i<len(*x);i++ {
		(*x)[i] = (*x)[i] +3
	}
	for i:=0;i<10;i++ {
		*x = append(*x,i)
	}

}

func f4(args ...int) {
	for i:=0;i<10;i++{
		args = append(args,i)
	}
}

func slice4() {
	var a []int
	f4(a...)  
	fmt.Println(a) //[] 可变参数满足和slice一样性质
}



func f5(args ...*int) {
	for i:=0;i<10;i++{
		args=append(args,&i)
	}
}

func slice5() {
	var a []*int
	f5(a...)
	fmt.Println(a) //[] 使用可变参数是千万不要再append
}

func f6(args ...int) {
	for i:=0;i<len(args);i++ {
		args[i]= args[i] +10
	}
}

func slice6() {
	var a =[]int{1,2,3}
	f6(a...)
	fmt.Println(a)  //[11 12 13]
}

/* 不能拿空接口去取值
func f7(args ...interface{}) {
	for i:=0;i<len(args);i++ {
		reflect.ValueOf(args).Index(i).SetInt(int64(i))
	}
}

func slice7() {
	var a = make([]interface{},7)
	f7(a)
	fmt.Println(a)
}
*/

func f3(x map[string]int) {
	x["uio"] = 10
	x["kk"] =20
}

func Map() {
	var a = make(map[string]int)
	f3(a)
	fmt.Println(a)  //map[kk:20 uio:10] 彻底改了，和slice不同
}