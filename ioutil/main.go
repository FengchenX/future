package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

func main() {
	//ioutilReadAll()
	//ioutilReadFile()
	//ioutilWriteFile()
	ioutilReadDir()
}
func ioutilReadAll() {
	reader := strings.NewReader("hello,world")
	if b, err := ioutil.ReadAll(reader); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}

func ioutilReadFile() {
	if file, err := os.Create("./feng/alg/ioutil/file.txt"); err != nil {
		fmt.Println(err)
	} else {
		str := "hello, world gogogogo"
		file.WriteString(str)
		file.Close()
	}
	if b, err := ioutil.ReadFile("./feng/alg/ioutil/file.txt"); err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(b))
	}
}

func ioutilWriteFile() {
	ioutil.WriteFile("./feng/alg/ioutil/writefile.ini", []byte("need go"), os.ModeType)
}
func ioutilReadDir() {
	if dirList, err := ioutil.ReadDir("./feng/alg/ioutil"); err == nil {
		for _, file := range dirList {
			fmt.Println(file.Name())
		}
	}
}
