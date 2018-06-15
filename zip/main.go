package main

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	//zipWrite()
	zipRead()
}

type File struct {
	name, body string
}

func zipWrite() {
	f, err := os.Create("10.zip")
	if err != nil {
		fmt.Println(err)
	}
	zw := zip.NewWriter(f)
	defer zw.Close()
	var files []File
	dir_list, err := ioutil.ReadDir("./test")
	for _, file := range dir_list {
		var one File
		one.name = file.Name()
		b, err := ioutil.ReadFile("./test/" + file.Name())
		if err != nil {
			fmt.Println(err)
		}
		one.body = string(b)
		files = append(files, one)
		fmt.Println(one.name, one.body)
	}
	for _, file := range files {
		w, err := zw.Create(file.name)
		if err != nil {
			fmt.Println(err)
		}
		w.Write([]byte(file.body))
	}
}

func zipRead() {
	zrc, err := zip.OpenReader("10.zip")
	defer zrc.Close()
	if err != nil {
		fmt.Println(err)
	}
	files := zrc.File
	for _, file := range files {
		f, err := os.Create(file.Name)
		if err != nil {
			fmt.Println(err)
		}
		defer f.Close()
		rc, err := file.Open()
		if err != nil {
			fmt.Println(err)
		}
		io.Copy(f, rc)
	}
}
