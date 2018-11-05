package log

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
	"time"

	"github.com/lestrrat/go-file-rotatelogs"
	"github.com/pkg/errors"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"

	"grm-service/path"
)

type LogConfs uint

const (
	LogNameComputer LogConfs = 1 << iota
	LogNameUser
	LogNoFileCreated
	LogStdStream
	FlagSetDefault = LogNameComputer | LogNameUser
)

/*
@param name 日志的标识性名称。
	 * @param path 日志文件所在的路径（若不指定则为当前路径）。
	 * @param config 日志的配置选项（默认为 LogNameComputer|LogNameUser）。
	 * @param saveTime 日志保留时间（默认3个月）。
     * @param rotationTime 日志分割时间（默认7天）。
	 * @note 日志的配置选项可以是以下值的按位或：
	 * - LogNameComputer    = 0x01： 日志名中包含计算机名
	 * - LogNameUser        = 0x02： 日志名中包含用户名
	 * - LogNoFileCreated   = 0x04： 不创建日志文件
	 * - LogStdStream       = 0x08： 同时输出日志到标准流(未提供该功能)
*/

func InitLog(name, logPath string, configs LogConfs, saveTime, rotationTime time.Duration) {
	logFile := name
	// 主机名
	if configs&LogNameComputer != 0 {
		host, _ := os.Hostname()
		logFile = fmt.Sprintf("%s_%s", logFile, host)
	}
	// 用户名
	if configs&LogNameUser != 0 {
		currentUser, _ := user.Current()
		host, _ := os.Hostname()

		username := strings.Replace(currentUser.Username, host, "", -1)
		username = strings.Replace(username, "\\", "", -1)
		logFile = fmt.Sprintf("%s_%s", logFile, username)
	}
	logFile = logFile + ".log"

	// 是否创建日志文件
	if configs&LogNoFileCreated == 0 {
		if !path.Exists(logPath) {
			if err := path.CreateAllPath(logPath); err != nil {
				logrus.Error("Failed to create log path:", logPath, ":", err)
				return
			}
		}

		baseLogPaht := filepath.Join(logPath, logFile)
		writer, err := rotatelogs.New(
			baseLogPaht+".%Y%m%d%H%M",
			rotatelogs.WithLinkName(baseLogPaht),      // 生成软链，指向最新日志文件
			rotatelogs.WithMaxAge(saveTime),           // 文件最大保存时间
			rotatelogs.WithRotationTime(rotationTime), // 日志切割时间间隔
		)
		if err != nil {
			logrus.Errorf("config local file system logger error. %+v", errors.WithStack(err))
			return
		}

		lfHook := lfshook.NewHook(lfshook.WriterMap{
			logrus.DebugLevel: writer, // 为不同级别设置不同的输出目的
			logrus.InfoLevel:  writer,
			logrus.WarnLevel:  writer,
			logrus.ErrorLevel: writer,
			logrus.FatalLevel: writer,
			logrus.PanicLevel: writer,
		}, &logrus.JSONFormatter{TimestampFormat: "2006-01-02 15:04:05"})
		logrus.AddHook(lfHook)
	}
}

// Print logs a message at level Info on the standard logger.
func Print(args ...interface{}) {
	logrus.Print(args...)
}

// Printf logs a message at level Info on the standard logger.
func Printf(format string, args ...interface{}) {
	logrus.Printf(format, args...)
}

// Println logs a message at level Info on the standard logger.
func Println(args ...interface{}) {
	logrus.Println(args...)
}

// Debug logs a message at level Debug on the standard logger.
func Debug(args ...interface{}) {
	logrus.Debug(args...)
}

// Debugf logs a message at level Debug on the standard logger.
func Debugf(format string, args ...interface{}) {
	logrus.Debugf(format, args...)
}

// Info logs a message at level Info on the standard logger.
func Info(args ...interface{}) {
	logrus.Info(args...)
}

// Infof logs a message at level Info on the standard logger.
func Infof(format string, args ...interface{}) {
	logrus.Infof(format, args...)
}

// Warn logs a message at level Warn on the standard logger.
func Warn(args ...interface{}) {
	logrus.Warn(args...)
}

// Warnf logs a message at level Warn on the standard logger.
func Warnf(format string, args ...interface{}) {
	logrus.Warnf(format, args...)
}

// Error logs a message at level Error on the standard logger.
func Error(args ...interface{}) {
	logrus.Error(args...)
}

// Errorf logs a message at level Error on the standard logger.
func Errorf(format string, args ...interface{}) {
	logrus.Errorf(format, args...)
}

// Panic logs a message at level Panic on the standard logger.
func Panic(args ...interface{}) {
	logrus.Panic(args...)
}

// Panicf logs a message at level Panic on the standard logger.
func Panicf(format string, args ...interface{}) {
	logrus.Panicf(format, args...)
}

// Fatal logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatal(args ...interface{}) {
	logrus.Fatal(args...)
}

// Fatalf logs a message at level Fatal on the standard logger then the process will exit with status set to 1.
func Fatalf(format string, args ...interface{}) {
	logrus.Fatalf(format, args...)
}
