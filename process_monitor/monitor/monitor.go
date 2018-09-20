package monitor

import (
	"bytes"
	"fmt"
	"os/exec"
	"path/filepath"
	"sub_account_service/process_monitor/config"
	"sub_account_service/process_monitor/model"
	"sub_account_service/process_monitor/service"
	"time"

	"github.com/sirupsen/logrus"
)

func StartMonitor() {
	ticker := time.NewTicker(time.Duration(int64(time.Second) * config.GetConfigInstance().Interval))
	for {
		select {
		case <-ticker.C:
			handleProcess()
		}
	}
}

func handleProcess() {
	defer func() {
		if err := recover(); err != nil {
			logrus.Panicln("monitor process panic:", err)
		}
	}()
	for key, value := range config.GetConfigInstance().Process {
		monitorProcess(key, &value)
	}
}

func monitorProcess(processName string, process *model.Process) {
	shellScript := fmt.Sprintf(` ps -ef |grep %v |grep -v grep |awk '{print $2}'`, processName)
	out, err := execShell(shellScript)
	if err != nil {
		logrus.WithError(err).Errorln("exec " + shellScript + " failed")
	}
	if out == "" { //该进程不存在，重启服务
		out, err := execShellFile(process.StartScript)
		if err != nil {
			flag := restarted(processName, false)
			if flag {
				service.SendMail(getMailTitle(processName)+"重启失败!", processName+" 挂了，重启失败："+err.Error(), process.Receivers)
			}
		} else {
			flag := restarted(processName, true)
			if flag {
				service.SendMail("重启成功！"+getMailTitle(processName), processName+" 挂了，重启成功："+out, process.Receivers)
			}
		}
	}
}

func getMailTitle(processName string) string {
	str := time.Now().Format("2006-01-02 15:04:05")
	return processName + str
}

//阻塞式的执行外部shell命令的函数,等待执行完毕并返回标准输出
func execShell(s string) (string, error) {
	logrus.Infoln("begin exec shell script:", s)
	//函数返回一个*Cmd，用于使用给出的参数执行name指定的程序
	cmd := exec.Command("/bin/sh", "-c", s)
	//读取io.Writer类型的cmd.Stdout，再通过bytes.Buffer(缓冲byte类型的缓冲器)将byte类型转化为string类型(out.String():这是bytes类型提供的接口)
	var out bytes.Buffer
	cmd.Stdout = &out
	//Run执行c包含的命令，并阻塞直到完成。  这里stdout被取出，cmd.Wait()无法正确获取stdin,stdout,stderr，则阻塞在那了
	err := cmd.Run()
	str := out.String()
	logrus.Infoln("exec shell scrip:"+s+"end,out:", str, ",err:", err)
	if err != nil && err.Error() != "<nil>" {
		return "", err
	}
	return str, nil
}

func execShellFile(filePath string) (string, error) {
	parentDir := filepath.Dir(filePath)
	fileName := filepath.Base(filePath)
	shellScript := fmt.Sprintf("cd %v;sh ./%v", parentDir, fileName)
	return execShell(shellScript)
}
