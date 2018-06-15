package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	//osOpen()
	name, _ := os.Hostname() //返回主机名
	fmt.Println(name)
	//osStat()
	//osMakeDir()
	//osRename()
	//osCreate()
	Delete()
}

var f1path *string = flag.String("f1path", "./mkDir/file.go", "Use -f1path <filesource>")

func osOpen() {
	flag.Parse()
	file, err := os.Open(*f1path)
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
	flag.Parse()
	fileInfo, err := os.Stat(*f1path)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(fileInfo.ModTime()) //最后修改时间
	fmt.Println(fileInfo.Size())    //文件字节数
}

var newDir *string = flag.String("newDir", "./mkDir", "Use -newDir <dir path>")

func osMakeDir() {
	flag.Parse()
	if err := os.Mkdir(*newDir, os.ModeDir); err != nil {
		log.Fatal(err)
	}
}

func osRename() {
	//修改一个文件名字，移动一个文件
	old := "./feng/alg/os/file.go"
	new := "./feng/alg/os/mkDir/file.go"
	if err := os.Rename(old, new); err != nil {
		log.Fatal(err)
	}
}

func osCreate() {
	//创建文件
	file, err := os.Create("./feng/alg/os/cre.txt")
	if err != nil {
		log.Fatal(err)
	}
	reader := strings.NewReader("hello, world!")
	reader.WriteTo(file)
}

//删除文件
var path *string = flag.String("path", "", "Use -path <filename>")

func Delete() {
	flag.Parse()
	cmd := exec.Command("powershell", "rm", *path)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}
