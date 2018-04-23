package mqtt_blacklist

import (
	"net/http"
	"sync"
	"strings"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"time"

)

var mu sync.RWMutex //guards access to fields below
var blackLists map[string]time.Time = make(map[string]time.Time) //ip: 到期时间
var lock sync.RWMutex //guards access to fields below
var whiteLists map[string]bool = make(map[string]bool)  //key 是ip此处考虑使用map而不是slice是出于添加删除查找方便些考虑
var db *sql.DB //数据库

const Queue = 10000  //一般通道缓存
const maxCount = 1000  //访问莫个url最大访问次数
const (
	stdCounterT = 100*time.Second  //counter 检查时间周期,用于对counter.clents检查时间间隔判断
	stdClientT = 60*time.Second    //client 生命周期,用于判断client是否已经过期
	stdDay = 10                    //访问次数过多时封禁ip天数
)

var (
	ipUrls = make(chan IPURL,Queue)
	jobs = make(chan job,Queue)
)

/**初始化，外部app调用执行
*indb:外部app传入的数据库对象
*/
func InitBlacklist(indb *sql.DB) {
	
	/*
	db, err := sql.Open("mysql", "root:feng@/test?charset=utf8")
	if err != nil {
		fmt.Println(err)
	}
	defer db.Close()
	*/
	db=indb
	if rows, err := db.Query("SELECT * FROM blacklist"); err != nil {
		fmt.Println("读取黑名单数据库所有行错误:",err)
	} else {
		//遍历每一行
		for rows.Next() {
			var uid int
			var ip string
			var expiretime string
			err=rows.Scan(&uid,&ip,&expiretime) //将每行数据写入ip expiretime中
			if err != nil {
				fmt.Println("rows.Scan失败:",err)
			}
			fmt.Println(uid)
			fmt.Println(ip)
			fmt.Println(expiretime)
			const shortForm = "2006-1-02"
			t, _ := time.Parse(shortForm, expiretime)  //将从数据库中取到的时间按格式解析为time.Time
			fmt.Println(t)
			blackLists[ip]=t //将ip和他的到期时间写到黑名单中
		}
		rows.Close()
	}

	for ip,expiretime:=range blackLists {
		if time.Now().After(expiretime) {
			//这个黑名单中的ip已经过期了，删除
			delete(blackLists,ip)
		}
	}
	
	if rows, err := db.Query("SELECT * FROM whitelist"); err != nil {
		fmt.Println("读取白名单所有行错误:",err)
	} else {
		for rows.Next() {
			var uid int
			var ip string
			err = rows.Scan(&uid, &ip)
			if err != nil {
				fmt.Println(err)
			}
			whiteLists[ip]= true  //将数据库中取到的这个ip添加到白名单中
		}
		rows.Close()
	}
	
	go writeDb()  //开启往数据库中写黑白名单gorutine
	go stat()	  //开启统计gorutine,统计某个ip在给定时间内访问某个url次数，如果流量过大，就放入黑名单通道jobs中，jobs是个全局变量
	
	go writeBlacklists()  //开启往黑名单中写gorutine，不断从黑名单通道jobs中取出ip并添加
}


/**往数据库中写黑白名单*/
func writeDb() {
	t:=time.NewTicker(24*time.Hour)
	for {
		<-t.C   //每隔一天允许通过一次
		//写黑名单部分
		//准备一个删除状态
		del,err:=db.Prepare("truncate table blacklist");//清空数据库命令
		if err != nil {
			fmt.Println("清空数据库命令错误",err)
			return
		}
		//准备一个任务
		tx, err := db.Begin();
		if err != nil {
			fmt.Println(err)
			return	
		}
		//删除黑名单表中所有行
		tx.Stmt(del).Exec()//tx通过del状态生成一个特定的删除状态，然后去执行删除任务

		//将黑名单写到数据库中	
		insert,err := db.Prepare("insert into blacklist (ip, expiretime) values(?,?)");
		if  err != nil {
			fmt.Println(err)
			return
		}
		mu.RLock()
		for ip,expiretime:=range blackLists {
			tx.Stmt(insert).Exec(ip,expiretime)  //通过insert状态，然后执行insert任务. 将黑名单写到数据库中
		}
		mu.RUnlock()
		//提交到数据库
		if err = tx.Commit(); err!= nil {
			fmt.Println("提交到数据库错误:",err)
		}
		insert.Close()
		del.Close()

		/**写白名单部分*/
		del,err = db.Prepare("truncate table whitelist");//清空数据库命令
		if err != nil {
			fmt.Println(err)
			return
		}
		tx,err = db.Begin();
		if err != nil {
			fmt.Println(err)
			return
		}
		//删除白名单表中所有行
		tx.Stmt(del).Exec()
		//将白名单写到数据库中	
		insert,err = db.Prepare("insert into whitelist (ip) values(?)");
		if err != nil {
			fmt.Println(err)
			return
		}
		lock.RLock()
		for ip:=range whiteLists{
			tx.Stmt(insert).Exec(ip)
		}
		lock.RUnlock()
		//提交到数据库
		if err = tx.Commit(); err!= nil {
			fmt.Println("提交到whitelist错误:",err)
		}
		insert.Close()
		del.Close()
	}
}

type client struct {
	counters int //给定时间内已经访问次数
	t time.Time //客户端第一次访问时间
	T time.Duration //生命周期 当time.Now()-t>T时，说明这个client对象过期了应该被处理了
}

type counter struct {
	clients map[IPURL]*client //key:ip+url value:*client
	
	lastTime time.Time //上次检查时间 初始化为time.Now()
	T time.Duration//检查时间周期 当time.Now() - lasttime > T时，说明我们需要去检查一下clients中哪些选项应该被删除了
}

type IPURL string  //IPURL是string的别名

/**创建一个统计对象counter
*t:检查时间周期会赋值给counter.T
*/
func newCounter(t time.Duration) *counter {
	c:=&counter{
		clients: make(map[IPURL]*client),
		lastTime: time.Now(),
		T: t,
	}
	return c
}

/**统计ip访问莫个url次数并将其放入黑名单通道jobs中*/
func stat() {
	coun:=newCounter(stdCounterT)
	for {
		select {
		case ipUrl:= <-ipUrls:
			if _,ok:=coun.clients[ipUrl]; ok {
				coun.clients[ipUrl].counters++
			} else {
				client:=client{
					counters: 1,
					t: time.Now(),
					T: stdClientT, 
				}
				coun.clients[ipUrl]=&client
			}
			fmt.Println("访问次数:",ipUrl,coun.clients[ipUrl].counters)
			if coun.clients[ipUrl].counters>maxCount {
				//需要加入黑名单了
				ip:=strings.SplitAfter(string(ipUrl),":")[0]
				ip=strings.TrimSuffix(ip,":")
				var j = job{ip:ip,day:stdDay}
				jobs <- j   //将当前需要加入黑名单的工作放入jobs管道
			}
		default:
			if t:=time.Now().Sub(coun.lastTime); t>coun.T {
				//该清理clients了
				for ip, client:=range coun.clients {
					if dua:=time.Now().Sub(client.t); dua>client.T {
						delete(coun.clients,ip)
					}
				}
			}
			coun.lastTime=time.Now()
		}
	}
}

/**不断从jobs管道中取出ip然后写入黑名单中*/
func writeBlacklists() {
	defer func() {
		close(jobs)
	}()

	lastTime:=time.Now()
	T:= 600*time.Second //todo 待修改
	for {
		select{
		case job:=<-jobs:
			expireTime:=time.Now().AddDate(0,0,job.day)
			mu.Lock()
			blackLists[job.ip]=expireTime
			mu.Unlock()
		default:
			if dua:=time.Now().Sub(lastTime); dua > T {  
				lastTime=time.Now()

				mu.Lock()
				for ip,expiretime:=range blackLists {
					//该删除一些过期的黑名单了
					if time.Now().After(expiretime) {
						delete(blackLists,ip)
					}
				}
				mu.Unlock()
			}
		}
	}
}

type job struct {
	ip string  //ip地址
	day int    //要封ip多少天
}

/**如果返回值为true则ip在黑名单中,false 则ip 不在黑名单中*/
func DoFilter(r *http.Request, w http.ResponseWriter) bool {
	ip:=strings.SplitAfter(r.RemoteAddr,":")[0]
	ip=strings.TrimSuffix(ip,":")
	lock.RLock()
	_,inWhite := whiteLists[ip]
	lock.RUnlock()
	if inWhite {
		//在白名单中直接返回false
		fmt.Println("此 ip 在白名单中")
		return false
	}
	mu.RLock()
	_,ok:=blackLists[ip]
	mu.RUnlock()
	
	if ok {
		//此IP在黑名单中
		return true

	} else {
		//此IP不在黑名单中
		ipUrl:=ip+r.URL.Path
		ipUrls <- IPURL(ipUrl)
		return false
	}
}

/**主动添加黑名单*/
func Add(ip string, day int) {
	mu.RLock()
	_,ok:=blackLists[ip]
	mu.RUnlock()
	
	if ok {
		//此IP已在黑名单中
		fmt.Println("此IP已在黑名单中，请不要重复添加")

	} else {
		//此IP不在黑名单中,可以添加
		var j = job{ip: ip, day: day}
		jobs <- j
	}
}
/**添加白名单*/
func AddWhite(ip string) {
	lock.Lock()
	defer lock.Unlock()
	if _,ok:=whiteLists[ip]; ok {
		fmt.Println("此ip已经在白名单中请不要重复添加")
	} else {
		whiteLists[ip]=true  //添加白名单
	}
}