package download

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
)

const logDir = "/log" //log目录
type File struct {
	Id  string //文件名
	Url string //下载文件时用的url
}

type OnlineFile struct {
	Files []File
}

/**列出log目录下的所有log文件*/
func ListFunc(w http.ResponseWriter, r *http.Request) {
	fmt.Println(r.URL.Path)
	dir_list, e := ioutil.ReadDir(".." + logDir)
	if e != nil {
		fmt.Println("read dir error")
		return
	}
	var myOnline OnlineFile
	for _, v := range dir_list {
		ss := strings.TrimSuffix(v.Name(), ".log")
		file := File{Id: ss, Url: logDir + "/" + v.Name()}
		myOnline.Files = append(myOnline.Files, file)
	}
	t, err := template.ParseFiles("../list.html") //模板解析html文件
	if err != nil {
		panic(err)
	}

	err = t.Execute(w, myOnline) //将myOnline数据动态放到list.html中
	if err != nil {
		panic(err)
	}
}

/**处理下载文件*/
func DealStaticFiles(w http.ResponseWriter, r *http.Request) {
	rootDir := "./.." //项目的根目录

	//访问的url路径里包含logDir
	if strings.HasPrefix(r.URL.Path, logDir) {
		file := rootDir + r.URL.Path //文件名
		fmt.Println(file)
		f, err := os.Open(file)
		if err != nil && os.IsNotExist(err) {
			fmt.Fprintln(w, "File not exist")
			return
		}
		defer f.Close()
		http.ServeFile(w, r, file) //将文件内容写到客户端
		return
	} else {
		fmt.Fprintln(w, "Hello world")
	}
}
