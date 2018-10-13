package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

//RenameLogFile 按日期分包
func RenameLogFile() {
	fileBase := "applog"
	dir := "./log/static/"
	lastDay := time.Now().Day()
	var oldfile *os.File
	defer func() {
		if oldfile != nil {
			oldfile.Close()
		}
	}()
	f := func() {
		now := time.Now()
		year, month, day := now.Year(), now.Month(), now.Day()
		lastDay = day

		filePath := fmt.Sprintf("%s%s.%d-%d-%d.log", dir, fileBase, year, month, day)
		file, err := os.OpenFile(filePath, os.O_CREATE|os.O_RDWR|os.O_APPEND, os.ModeAppend|os.ModePerm)
		if err != nil {
			fmt.Println(err)
		}
		if oldfile != nil {
			oldfile.Close()
		}
		logrus.SetOutput(file)
		oldfile = file
	}
	f()
	for {
		if time.Now().Day()-lastDay > 0 {
			f()
		}
		time.Sleep(1 * time.Second)
	}
}
