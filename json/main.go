




package main

import (
	"flag"
	"path/filepath"
	"bytes"
	"fmt"
	"io"
	"strings"
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

)
//json使用是结构体字段变量必须大写才能被识别

var in *string = flag.String("path",".","Use -path <filesource>")

var jsonStream string
//此时控制台应处于github.com目录下
func getCurDir() string {
	dir, err := filepath.Abs(filepath.Dir(*in))
	if err != nil {
		log.Fatal(err)
	}
	return strings.Replace(dir,"\\","/",-1)
}
func init() {
	flag.Parse()
	
	filepath:=getCurDir()
	
	fmt.Println(filepath)
	f,err:=os.Open(filepath+"/"+"file.txt")
	if err != nil {
		log.Println(err)
	}
	defer f.Close()
	b,err:=ioutil.ReadAll(f)
	jsonStream=string(b)

}
func main() {
	
	//jsonNewDecoder()
	//jsonIndent()
	//jsonMarshal()
	jsonRawMessage()
}

type Message struct {
	Name, Text string
}
func jsonNewDecoder() {
	dec:= json.NewDecoder(strings.NewReader(jsonStream))	
	for {
		var m Message
		if err:=dec.Decode(&m); err == io.EOF {
			break
		} else if err != nil {
			log.Fatal(err)
		}
		fmt.Printf("%s: %s\n",m.Name,m.Text)
	}
}


type Road struct {
	Name string
	Number int
}
func jsonIndent() {
	roads:=[]Road{
		{"Diamond Fork",29},
		{"Sheep Creek",51},
	}
	b, err := json.Marshal(roads)
	if err != nil {
		log.Fatal(err)
	}
	var out bytes.Buffer
	json.Indent(&out,b,"=","\t")
	out.WriteTo(os.Stdout)
}

type ColorGroup struct {
	ID int
	Name string
	Colors []string
}
func jsonMarshal() {
	group := ColorGroup{
		ID: 1,
		Name: "Reds",
		Colors: []string{"Crimson", "Red", "Ruby", "Maroon" },
	}
	b, err := json.Marshal(group)
	if err != nil {
		fmt.Println("error: ", err)
	}
	os.Stdout.Write(b)
	fmt.Println()
}

type Color struct {
	Space string
	Point json.RawMessage
}
type RGB struct {
	R uint8
	G uint8
	B uint8
}
type YCbCr struct {
	Y uint8
	Cb int8
	Cr int8
}

//延迟解析很吊
func jsonRawMessage() {
	var j = []byte(`[
					{"Space":"YCbCr", "Point": {"Y": 255, "Cb": 0, "Cr": -10}},
					{"Space": "RGB", "Point": {"R": 98, "G": 218, "B": 255}}
	]`)
	var colors []Color
	err := json.Unmarshal(j, &colors)
	if err != nil {
		log.Fatalln("error: ", err)
	}
	for _, c := range colors {
		var dst interface{}
		switch c.Space {
		case "RGB":
			dst = new(RGB)
		case "YCbCr":
			dst = new(YCbCr)
		}
		err := json.Unmarshal(c.Point, dst)
		if err != nil {
			log.Fatalln("error: ", err)
		}
		fmt.Println(c.Space, dst)
	}
}
