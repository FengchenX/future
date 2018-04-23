
package main

import (
	"strings"
	"fmt"
	"log"
	"os"

)

func main() {
	//osOpen()
	name,_ := os.Hostname() //返回主机名
	fmt.Println(name)
	//osStat()
	//osMakeDir()
	//osRename()
	osCreate()
}
func osOpen() {
	file, err := os.Open("./feng/alg/os/file.go")
	if err != nil {
		log.Fatal(err)
	}
	data := make([]byte, 100)
	count, err := file.Read(data)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("read %d bytes: %q\n", count, data[:count])
}

func osStat() {
	fileInfo, err := os.Stat("./feng/alg/os/file.go")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.ModTime())  //最后修改时间
	fmt.Println(fileInfo.Size()) //文件字节数
}

func osMakeDir() {
	if err:=os.Mkdir("./feng/alg/os/mkDir",os.ModeDir); err != nil {
		log.Fatal(err)
	}
}

func osRename() {
	//修改一个文件名字，移动一个文件
	old:="./feng/alg/os/file.go"
	new:= "./feng/alg/os/mkDir/file.go"
	if err:= os.Rename(old,new); err != nil {
		log.Fatal(err)
	}
}

func osCreate() {
	//创建文件
	file, err := os.Create("./feng/alg/os/cre.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader:=strings.NewReader("hello, world!")
	reader.WriteTo(file)
}