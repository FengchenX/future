package main

import (
	"fmt"
	"io"
	"log"
	"os"
	"encoding/json"
)

func main() {
	logSingleton = newLog(os.Stderr, "", log.LstdFlags)
	logSingleton.Println("feng")
	var c = child{key: 20}
	c.base = base{i: 30}
	f1(c)
	f2()
	fmt.Println("********************************")
	//多重继承测试
	var tt testC

	tt.A = "aaaa"
	tt.B =20
	tt.C = "ccccc" 
	tt.D = "ddddd"
	buf, err := json.Marshal(tt)
	if err != nil {
		log.Fatalln(err)
	}
	var yy struct {
		A string
		B int
		C string
		D string
	}

	json.Unmarshal(buf, &yy)
	fmt.Printf("%v\n", yy)
}

var logSingleton *_Logger

type _Logger struct {
	*log.Logger

	key int
}

func newLog(out io.Writer, prefix string, flag int) *_Logger {
	var myLg _Logger
	myLg.key = 0
	myLg.Logger = log.New(out, prefix, flag)
	return &myLg
}

type base struct {
	i int
}

func (b base) say() {
	fmt.Println("I am base")
}

type child struct {
	base
	key int
}

func (c child) say() {
	fmt.Println("I am child")
}

type i interface {
	say()
}

func f1(a i) {
	a.say()
}

//interface 组合
type a interface {
	addOne()
}
type b interface {
	subOne()
}

type c interface {
	a
	b
}

type myint int

func (m *myint) addOne() {
	*m = *m + 1
}

func (m *myint) subOne() {
	*m = *m - 1
}

func f2() {
	var A myint
	var d c
	d = &A
	d.addOne()
	d.addOne()
	fmt.Println(A)
	var e b
	e = &A
	e.subOne()
	fmt.Println(A)
}

type testA struct {
	A string
	B int
}

type testB struct {
	C string
	D string
}

type testC struct {
	testA
	testB
}
