package executor

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"sync"
)

var (
	// finishedErr is the error message received when trying to kill and already
	// exited process.
	finishedErr = "os: process already finished"
)

type HandlerFunc func(msg string)

type Executor struct {
	CmdPath        string
	Args           []string
	MessageHandler HandlerFunc

	cmd       exec.Cmd
	MutexLock sync.RWMutex
}

// NewExecutor returns an Executor
func NewExecutor(cmdPath string, args []string) Executor {
	exec := Executor{
		CmdPath: cmdPath,
		Args:    args,
	}
	return exec
}

func KillProcess(pid int) error {
	proc, err := os.FindProcess(pid)
	if err != nil {
		return fmt.Errorf("executor.shutdown failed to find process: %v", err)
	}
	if runtime.GOOS == "windows" {
		if err := proc.Kill(); err != nil && err.Error() != finishedErr {
			return err
		}
		return nil
	}
	if err = proc.Signal(os.Interrupt); err != nil && err.Error() != finishedErr {
		return fmt.Errorf("executor.shutdown error: %v", err)
	}
	return nil
}

func (e *Executor) LaunchCmd(workDir string, handler HandlerFunc) error {
	// 初始化日志文件
	infoFile, err := os.Create(filepath.Join(workDir, "log/stdout.log"))
	defer infoFile.Close()
	if err != nil {
		return fmt.Errorf("open file error: %s", filepath.Join(workDir, "log/stdout.log"))
	}
	infoLog := log.New(infoFile, "", 0)

	errFile, err := os.Create(filepath.Join(workDir, "log/stderr.log"))
	defer errFile.Close()
	if err != nil {
		return fmt.Errorf("open file error: %s", filepath.Join(workDir, "log/stderr.log"))
	}
	errLog := log.New(errFile, "", 0)

	infoLog.Println("INFO: launching command")

	// Set the commands arguments
	e.cmd.Path = e.CmdPath
	e.cmd.Args = append([]string{e.cmd.Path}, e.Args...)
	fmt.Println("cmd args： ", e.cmd.Args)

	stdout, err := e.cmd.StdoutPipe()
	if err != nil {
		errLog.Println("ERROR:", err.Error())
		return err
	}

	stderr, err := e.cmd.StderrPipe()
	if err != nil {
		errLog.Println("ERROR:", err.Error())
		return err
	}

	// Start the process
	e.MessageHandler = handler
	if err := e.cmd.Start(); err != nil {
		err = fmt.Errorf("Failed to start command path=%q --- args=`%v`: %v", e.cmd.Path, e.cmd.Args, err)
		errLog.Println("ERROR:", err.Error())
		return err
	} else {
		msg := fmt.Sprintf("[Status:Running(%d)]\n", e.cmd.Process.Pid)
		go e.MessageHandler(msg)
	}

	// 循环读取输出流中的一行内容
	go func() {
		reader := bufio.NewReader(stderr)
		for {
			line, err := reader.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			errLog.Printf("%s", line)
		}
	}()

	go func() {
		reader := bufio.NewReader(stdout)
		for {
			line, err := reader.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			infoLog.Printf("%s", line)

			go e.MessageHandler(line)
		}
	}()

	if err := e.cmd.Wait(); err != nil {
		go e.MessageHandler("[Status:Terminated]\n")
		return err
	} else {
		go e.MessageHandler("[Status:Finished]\n")
	}
	return nil
}

func (e *Executor) LaunchCmdEx() error {
	// Set the commands arguments
	e.cmd.Path = e.CmdPath
	e.cmd.Args = append([]string{e.cmd.Path}, e.Args...)
	fmt.Println("cmd args： ", e.cmd.Args)

	stdout, err := e.cmd.StdoutPipe()
	if err != nil {
		log.Println("ERROR:", err.Error())
		return err
	}

	stderr, err := e.cmd.StderrPipe()
	if err != nil {
		log.Println("ERROR:", err.Error())
		return err
	}

	// Start the process
	if err := e.cmd.Start(); err != nil {
		err = fmt.Errorf("Failed to start command path=%q --- args=`%v`: %v", e.cmd.Path, e.cmd.Args, err)
		log.Println("ERROR:", err.Error())
		return err
	} else {
		msg := fmt.Sprintf("[Status:Running(%d)]\n", e.cmd.Process.Pid)
		log.Println(msg)
	}

	// 循环读取输出流中的一行内容
	go func() {
		reader := bufio.NewReader(stderr)
		for {
			line, err := reader.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			log.Printf("%s", line)
		}
	}()

	go func() {
		reader := bufio.NewReader(stdout)
		for {
			line, err := reader.ReadString('\n')
			if err != nil || io.EOF == err {
				break
			}
			log.Printf("%s", line)
		}
	}()

	if err := e.cmd.Wait(); err != nil {
		log.Printf("[Status:Terminated]\n")
		return err
	} else {
		log.Printf("[Status:Finished]\n")
	}
	return nil
}
