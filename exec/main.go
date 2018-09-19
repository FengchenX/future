package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"os/exec"
)

func main() {
	//execOutput()
	//execStart()
	//execStdoutPipe()
	execShutdown()
}

func execOutput() {
	out, err := exec.Command("powershell", "date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
	fmt.Println(out)
}

func execStart() {
	cmd := exec.Command("powershell", "sleep", "5")
	err := cmd.Start()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("waiting for command to finish...")
	err = cmd.Wait()
	log.Printf("Command finished with error: %v", err)
}

//windows not pass
func execStdoutPipe() {
	cmd := exec.Command("powershell", "echo", "-n", `{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age  int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err != nil {
		log.Fatal(err)
	}
	if err := cmd.Wait(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s is %d years old\n", person.Name, person.Age)
}

/**
	关机命令
	c := exec.Command("powershell", "shutdown","-s","-t 0")
    if err := c.Run(); err != nil {
        fmt.Println("Error: ", err)
	}
*/

/** 自动关机程序 */
type MyTime struct {
	H int
	M int
}

const (
	stdH = 3600
	stdM = 60
)

func execShutdown() {
	data, err := ioutil.ReadFile("config.ini")
	if err != nil {
		log.Fatal(err)
	}
	var delay MyTime
	err = json.Unmarshal(data, &delay)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(delay.H, delay.M)
	delSecond := delay.H*stdH + delay.M*stdM
	str := fmt.Sprintf("%d", delSecond)
	cmd := exec.Command("powershell", "shutdown", "-s", "-t "+str)
	err = cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
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

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func execPowershell(s string) (string, error) {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("powershell", "-c", s)

	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out

	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	checkErr(err)

	return out.String(), err
}

func execCommand(commandName string, params []string) bool {
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command(commandName, params...)

	//显示运行的命令
	fmt.Println(cmd.Args)
	//StdoutPipe方法返回一个在命令Start后与命令标准输出关联的管道。Wait方法获知命令结束后会关闭这个管道，一般不需要显式的关闭该管道。
	stdout, err := cmd.StdoutPipe()

	if err != nil {
		fmt.Println(err)
		return false
	}

	cmd.Start()
	//创建一个流来读取管道内内容，这里逻辑是通过一行一行的读取的
	reader := bufio.NewReader(stdout)

	//实时循环读取输出流中的一行内容
	for {
		line, err2 := reader.ReadString('\n')
		if err2 != nil || io.EOF == err2 {
			break
		}
		fmt.Println(line)
	}

	//阻塞直到该命令执行完成，该命令必须是被Start方法开始执行的
	cmd.Wait()
	return true
}

//错误处理函数
func checkErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}
