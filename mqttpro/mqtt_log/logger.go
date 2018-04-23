

package mqtt_log

import (
	"fmt"
	"log"
	"os"
	"sync"
	"time"
	"net/http"
	"strings"
)


const Queue = 1000 //一些常用的缓存大小

var (
	dir string   //log目录
	filename string  //文件名
	maxNum int    //文件最大数量
	maxSize int64 //单个文件最大尺寸

	url string    //post的url路径

	maxWaitTime time.Duration  //最大等待时间
)

type _Logger struct {
	*log.Logger      //组合标准库Logger对象或者说继承

	suffix int    //每个文件尾

	mu sync.Mutex  //警戒进入下面区域
	logfile *os.File  //记录log的文件

	nowFileSize int64  //当前文件大小

	cache *Cache  //缓存对象
}

//log对象单例
var logSingleton *_Logger

func newLogger() *_Logger {
	logger:= _Logger {
		nowFileSize: 0,
	}

	//移除已经存在的log文件确保初始化nowFileSize=0正确
	if isExist(dir+"/"+filename) {
		os.Remove(dir+"/"+filename)
	}

	logfile, _ := os.OpenFile(dir+"/"+filename, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0666)
	logger.Logger=log.New(logfile,"logger: ",log.LstdFlags)
	
	logger.logfile = logfile
	return &logger
}


type Cache struct {
	CacheData string        //已经缓存的字节
	logs chan string        //Debug(xx,xx)函数会往logs通道中传数据
	lasttime time.Time      //上次写入文件时间
	dataLen int64           //已经缓存数据字节数
}

//初始化log库，最先执行
func InitLog(logDir string,
			fileName string,
			MaxNum int,
			MaxSize int64,
			maxWait time.Duration,
			flag int,
			Url string,
) {
	dir=logDir
	filename=fileName
	maxNum=MaxNum
	maxSize=MaxSize
	url=Url
	maxWaitTime=maxWait

	logSingleton=newLogger()

	cache:=logSingleton.newCacher()
	logSingleton.cache=cache

	logSingleton.SetFlags(flag)

	logSingleton.start()
}

//外部app结束时一定要调用,关闭log库
func Close() {
	if logSingleton != nil && logSingleton.logfile != nil {
		logSingleton.logfile.Close()
	}
}

//log库开始运行
func (l *_Logger) start() {
	go l.cache.run()
}

func (l *_Logger) newCacher() *Cache {
	return &Cache{
		CacheData: "",
		logs: make(chan string,Queue),
		lasttime: time.Now(),
		dataLen: 0,
	}
}

//cache对象不断从Cache.logs管道中读数据并写入文件中
func (c *Cache) run() {
	
	defer func() {
		if c.logs != nil {
			close(c.logs)
		}
	}()

	waitTime:=0*time.Second
	lastTime:=time.Now()
	var newSize int64 = 0
	for {
		select {
		case data:=<-c.logs:
			if len(data)==0 {
				continue
			}
			c.CacheData+=data
			c.dataLen=c.dataLen+int64(len(data))
			newSize=c.dataLen + logSingleton.nowFileSize
			if newSize>maxSize {

				logSingleton.Println(string(c.CacheData))//写入文件
				logSingleton.rename()//重命名

				logSingleton.nowFileSize=0
				c.CacheData=""
				c.dataLen=0
				newSize = 0
				lastTime=time.Now()
				waitTime = time.Now().Sub(lastTime)
			} else {
				if waitTime > maxWaitTime {
					if c.dataLen == 0 {
						lastTime=time.Now()
						waitTime = time.Now().Sub(lastTime)
					} else {
						logSingleton.nowFileSize=logSingleton.nowFileSize+c.dataLen//计算文件新的尺寸
						logSingleton.Println(string(c.CacheData))//写入文件

						if logSingleton.nowFileSize > maxSize {
							logSingleton.rename()
							logSingleton.nowFileSize = 0
						}
						c.CacheData=""
						c.dataLen=0
						newSize = 0
						lastTime=time.Now()
						waitTime = time.Now().Sub(lastTime)
					}
				} else {
					waitTime = time.Now().Sub(lastTime)
				}
			}
		default:
			time.Sleep(100*time.Millisecond)
			if waitTime > maxWaitTime {
				if c.dataLen == 0 {
					lastTime=time.Now()
					waitTime = time.Now().Sub(lastTime)
				} else {
					logSingleton.nowFileSize=logSingleton.nowFileSize+c.dataLen
					logSingleton.Println(string(c.CacheData))//写入文件

					if logSingleton.nowFileSize > maxSize {
						logSingleton.rename()
						logSingleton.nowFileSize = 0
					}
					c.CacheData=""
					c.dataLen=0
					newSize = 0
					lastTime=time.Now()
					waitTime = time.Now().Sub(lastTime)
				}
			} else {
				waitTime = time.Now().Sub(lastTime)
			}
		}
	}
}

//对文件进行改名
func (log *_Logger) rename() {
	log.suffix = log.suffix%maxNum + 1

	if log.logfile != nil {
		log.logfile.Close()
	}
	newpath := fmt.Sprintf("%s/%s.%d.%d.log", dir, filename,time.Now().Day(), log.suffix) // dir/filename.day.suffix.log
	if isExist(newpath) {
		os.Remove(newpath)
	}

	log.mu.Lock()
	defer log.mu.Unlock()

	filepath := dir + "/" + filename
	os.Rename(filepath,newpath)
	log.logfile, _ = os.Create(filepath)
	log.SetOutput(log.logfile)//设置log输出对象
}

//判断文件是否存在
func isExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}

func Debug(isPost bool, v ...interface{}) {
	//打印到控制台
	log.Println(v)//使用log库的标准输出

	//将要写到日志中的信息放入logs管道中供cache.run线程使用
	logSingleton.cache.logs<-fmt.Sprintln(v)
	
	if isPost {
		go DonormalPostorGet("","",fmt.Sprintln(v))
	}
}

//普通的post / get请求
func DonormalPostorGet(id string, name string, errmsg string) {
	var r http.Request
	r.ParseForm()
	r.Form.Add("id", id)   					 //  post id
	r.Form.Add("name", name) 				 //  post name
	r.Form.Add("err",errmsg)			 		 //  post error
	bodystr := strings.TrimSpace(r.Form.Encode())
	request, err := http.NewRequest("POST", url, strings.NewReader(bodystr))
	if err != nil {
		print("todo")
	}
	request.Header.Set("Content-Type", "application/x-www-form-urlencoded")
	request.Header.Set("Connection", "Keep-Alive")
	resp, err := http.DefaultClient.Do(request)
	
	if err != nil {
		fmt.Println("post fail")
	} else {
		fmt.Println("post ok")
	}
	defer resp.Body.Close()

}