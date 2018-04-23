

package main

import (
	"bytes"
	"bufio"
	"strings"
	"fmt"
	"os"
	"io"

)
func main() {
	//ioreader()
	//ioReaderAt()
	//ioWriteAt()
	//ioReadFrom()
	ioCopy()
}
func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p:= make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p, err
}

func ioreader() {
	data, err := ReadFrom(os.Stdin,11)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(data))
}

func ioReaderAt() {
	reader:=strings.NewReader("Go语言学习园地")
	p:= make([]byte, 6)
	n, err:= reader.ReadAt(p,2)
	if err != nil {
		panic(err)
	}
	fmt.Printf("%s, %d\n", p,n)
}

func ioWriteAt() {
	file, err := os.Create("./feng/alg/io/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.WriteString("Golang中文社区——这里是多余的")
	n, err := file.WriteAt([]byte("Go语言学习园地"), 24)
	if err != nil {
		panic(err)
	}
	fmt.Println(n)
}
func ioReadFrom() {
	file, err := os.Open("./feng/alg/io/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	writer := bufio.NewWriter(os.Stdout)
	writer.ReadFrom(file)
	writer.Flush()
}


func ioCopy() {
	
	file, err := os.Open("./feng/alg/io/writeAt.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var b []byte
	var buf = bytes.NewBuffer(b)//b在这之后不能用了
	
	_,err =io.Copy(buf,file)
	if err != nil {
		fmt.Println(err)
	}
	s:=string(buf.Bytes())
	fmt.Println(s)
	
}