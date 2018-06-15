package main

import (
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	logSingleton = newLog(os.Stderr, "", log.LstdFlags)
	logSingleton.Println("feng")
	var c = child{key: 20}
	c.base = base{i: 30}
	f1(c)
	f2()
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
