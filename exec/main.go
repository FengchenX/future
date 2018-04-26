

package main

import (
	"flag"
	"io/ioutil"
	"encoding/json"
	"fmt"
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
	out, err := exec.Command("powershell","date").Output()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("The date is %s\n", out)
	fmt.Println(out)
}

func execStart() {
	cmd := exec.Command("powershell","sleep", "5")
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
	cmd := exec.Command("powershell","echo", "-n",`{"Name": "Bob", "Age": 32}`)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		log.Fatal(err)
	}
	if err := cmd.Start(); err != nil {
		log.Fatal(err)
	}
	var person struct {
		Name string
		Age int
	}
	if err := json.NewDecoder(stdout).Decode(&person); err!= nil {
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
	err = json.Unmarshal(data,&delay)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(delay.H,delay.M)
	delSecond := delay.H*stdH + delay.M*stdM
	str := fmt.Sprintf("%d",delSecond)
	cmd := exec.Command("powershell","shutdown", "-s", "-t "+ str)
	err = cmd.Run()	
	if err != nil {
		log.Fatal(err)
	}
}

//删除文件
var path *string = flag.String("path", "","Use -path <filename>")
func Delete() {
	flag.Parse()
	cmd := exec.Command("powershell","rm", *path)
	err := cmd.Run()
	if err != nil {
		log.Fatal(err)
	}
}

