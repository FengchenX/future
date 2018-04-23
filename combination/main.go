



package main
import (
	"fmt"
	"os"
	"log"
	"io"
	
)

func main() {
	logSingleton=newLog(os.Stderr,"",log.LstdFlags)
	logSingleton.Println("feng")
	var c = child{key:20}
	c.base=base{i:30}
	f1(c)
}

var logSingleton *_Logger

type _Logger struct {
	*log.Logger

	key int
}

func newLog(out io.Writer, prefix string, flag int) *_Logger {
	var myLg _Logger
	myLg.key=0
	myLg.Logger=log.New(out,prefix,flag)
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

type i interface{
	say()
}

func f1(a i) {
	a.say()
}
