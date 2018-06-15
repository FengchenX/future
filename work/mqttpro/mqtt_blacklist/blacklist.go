package mqtt_blacklist

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"math"
	"net/http"
	"strings"
	"sync"
	"time"
)

var mu sync.RWMutex                                              //guards access to fields below
var blackLists map[string]time.Time = make(map[string]time.Time) //ip: 到期时间
var lock sync.RWMutex                                            //guards access to fields below
var whiteLists map[string]bool = make(map[string]bool)           //key 是ip此处考虑使用map而不是slice是出于添加删除查找方便些考虑
var db *sql.DB                                                   //数据库

const Queue = 10000 //一般通道缓存
const maxCount = 5  //10s访问莫个url最大访问次数
const (
	stdCounterT = 100 * time.Second //counter 检查时间周期,用于对counter.clents检查时间间隔判断
	stdClientT  = 10 * time.Second  //client 生命周期,用于判断client是否已经过期
	stdDay      = 10                //访问次数过多时封禁ip天数
)

const (
	A float64 = 2 + iota
	B
	C
)

var rlogin = map[int]float64{100: 1.5, 500: 1.2, 1000: 1.1, 2000: 1.1, 3000: 1.1, 4000: 1.1, 5000: 1.1, 6000: 1.05, 7000: 1.05, 8000: 1.05} //登录url 人数和访问概率map
var rurl = map[int]float64{100: 2, 500: 1, 1000: 0.8, 2000: 0.6, 3000: 0.6, 4000: 0.6, 5000: 0.55, 6000: 0.55, 7000: 0.5, 8000: 0.5}        //普通url人数和访问概率map

var threatLevel = A //威胁等级

var queuelist = Queuelist{tail: 0, list: make([]string, 50)} //最近五十条访问内容

var (
	ipUrls = make(chan IPURL, Queue) //ip和url组合的通道
	jobs   = make(chan job, Queue)   //存放需要加入黑名单的任务

	loginCh = make(chan int, Queue) //登录url计数
	urlCh   = make(chan int, Queue) //一般url计数
)

/**初始化，外部app调用执行
*indb:外部app传入的数据库对象
 */
func InitBlacklist(indb *sql.DB, num *int) {

	/*
		db, err := sql.Open("mysql", "root:feng@/test?charset=utf8")
		if err != nil {
			fmt.Println(err)
		}
		defer db.Close()
	*/
	db = indb
	if rows, err := db.Query("SELECT * FROM blacklist"); err != nil {
		fmt.Println("读取黑名单数据库所有行错误:", err)
	} else {
		//遍历每一行
		for rows.Next() {
			var uid int
			var ip string
			var expiretime string
			err = rows.Scan(&uid, &ip, &expiretime) //将每行数据写入ip expiretime中
			if err != nil {
				fmt.Println("rows.Scan失败:", err)
			}
			fmt.Println(uid)
			fmt.Println(ip)
			fmt.Println(expiretime)
			const shortForm = "2006-1-02"
			t, _ := time.Parse(shortForm, expiretime) //将从数据库中取到的时间按格式解析为time.Time
			fmt.Println(t)
			blackLists[ip] = t //将ip和他的到期时间写到黑名单中
		}
		rows.Close()
	}

	for ip, expiretime := range blackLists {
		if time.Now().After(expiretime) {
			//这个黑名单中的ip已经过期了，删除
			delete(blackLists, ip)
		}
	}

	if rows, err := db.Query("SELECT * FROM whitelist"); err != nil {
		fmt.Println("读取白名单所有行错误:", err)
	} else {
		for rows.Next() {
			var uid int
			var ip string
			err = rows.Scan(&uid, &ip)
			if err != nil {
				fmt.Println(err)
			}
			whiteLists[ip] = true //将数据库中取到的这个ip添加到白名单中
		}
		rows.Close()
	}

	go writeDb() //开启往数据库中写黑白名单gorutine
	go stat()    //开启统计gorutine,统计某个ip在给定时间内访问某个url次数，如果流量过大，就放入黑名单通道jobs中，jobs是个全局变量

	go writeBlacklists() //开启往黑名单中写gorutine，不断从黑名单通道jobs中取出ip并添加

	go calThreatLevel(num) //开启计算威胁等级goroutine
}

/**往数据库中写黑白名单*/
func writeDb() {
	t := time.NewTicker(24 * time.Hour)
	for {
		<-t.C //每隔一天允许通过一次
		//写黑名单部分
		//准备一个删除状态
		del, err := db.Prepare("truncate table blacklist") //清空数据库命令
		if err != nil {
			fmt.Println("清空数据库命令错误", err)
			return
		}
		//准备一个任务
		tx, err := db.Begin()
		if err != nil {
			fmt.Println(err)
			return
		}
		//删除黑名单表中所有行
		tx.Stmt(del).Exec() //tx通过del状态生成一个特定的删除状态，然后去执行删除任务

		//将黑名单写到数据库中
		insert, err := db.Prepare("insert into blacklist (ip, expiretime) values(?,?)")
		if err != nil {
			fmt.Println(err)
			return
		}
		mu.RLock()
		for ip, expiretime := range blackLists {
			tx.Stmt(insert).Exec(ip, expiretime) //通过insert状态，然后执行insert任务. 将黑名单写到数据库中
		}
		mu.RUnlock()
		//提交到数据库
		if err = tx.Commit(); err != nil {
			fmt.Println("提交到数据库错误:", err)
		}
		insert.Close()
		del.Close()

		/**写白名单部分*/
		del, err = db.Prepare("truncate table whitelist") //清空数据库命令
		if err != nil {
			fmt.Println(err)
			return
		}
		tx, err = db.Begin()
		if err != nil {
			fmt.Println(err)
			return
		}
		//删除白名单表中所有行
		tx.Stmt(del).Exec()
		//将白名单写到数据库中
		insert, err = db.Prepare("insert into whitelist (ip) values(?)")
		if err != nil {
			fmt.Println(err)
			return
		}
		lock.RLock()
		for ip := range whiteLists {
			tx.Stmt(insert).Exec(ip)
		}
		lock.RUnlock()
		//提交到数据库
		if err = tx.Commit(); err != nil {
			fmt.Println("提交到whitelist错误:", err)
		}
		insert.Close()
		del.Close()
	}
}

type client struct {
	counters int           //给定时间内已经访问次数
	t        time.Time     //客户端第一次访问时间
	T        time.Duration //生命周期 当time.Now()-t>T时，说明这个client对象过期了应该被处理了
}

type counter struct {
	clients map[IPURL]*client //key:ip+url value:*client

	lastTime time.Time     //上次检查时间 初始化为time.Now()
	T        time.Duration //检查时间周期 当time.Now() - lasttime > T时，说明我们需要去检查一下clients中哪些选项应该被删除了
}

type IPURL string //IPURL是string的别名

/**创建一个统计对象counter
*t:检查时间周期会赋值给counter.T
 */
func newCounter(t time.Duration) *counter {
	c := &counter{
		clients:  make(map[IPURL]*client),
		lastTime: time.Now(),
		T:        t,
	}
	return c
}

/**统计ip访问莫个url次数并将其放入黑名单通道jobs中*/
func stat() {
	coun := newCounter(stdCounterT)
	for {
		select {
		case ipUrl := <-ipUrls:
			if _, ok := coun.clients[ipUrl]; ok {
				if client := coun.clients[ipUrl]; time.Now().After(client.t.Add(client.T)) {
					//这个client过期了
					client.counters, client.t = 0, time.Now()
				} else {
					//这个client没有过期
					client.counters++
				}
			} else {
				client := client{
					counters: 1,
					t:        time.Now(),
					T:        stdClientT,
				}
				coun.clients[ipUrl] = &client
			}
			fmt.Println("访问次数:", ipUrl, coun.clients[ipUrl].counters)
			if coun.clients[ipUrl].counters > maxCount {
				//需要加入黑名单了
				ip := strings.SplitAfter(string(ipUrl), "/")[0] //ipUrl格式为 127.0.0.1/postForm 将字符串按 / 分离 得到 127.0.0.1/
				ip = strings.TrimSuffix(ip, "/")                //将字符串去尾 / ，得到127.0.0.1
				var j = job{ip: ip, day: stdDay}
				jobs <- j //将当前需要加入黑名单的工作放入jobs管道
			}
		default:
			if t := time.Now().Sub(coun.lastTime); t > coun.T {
				//该清理clients了
				for ip, client := range coun.clients {
					if dua := time.Now().Sub(client.t); dua > client.T {
						delete(coun.clients, ip)
					}
				}
			}
			coun.lastTime = time.Now()
		}
	}
}

/**不断从jobs管道中取出ip然后写入黑名单中*/
func writeBlacklists() {
	defer func() {
		close(jobs)
	}()

	lastTime := time.Now()
	T := 600 * time.Second //todo 待修改
	for {
		select {
		case job := <-jobs:
			expireTime := time.Now().AddDate(0, 0, job.day)
			mu.Lock()
			blackLists[job.ip] = expireTime
			mu.Unlock()
		default:
			if dua := time.Now().Sub(lastTime); dua > T {
				lastTime = time.Now()

				mu.Lock()
				for ip, expiretime := range blackLists {
					//该删除一些过期的黑名单了
					if time.Now().After(expiretime) {
						delete(blackLists, ip)
					}
				}
				mu.Unlock()
			}
		}
	}
}

type job struct {
	ip  string //ip地址
	day int    //要封ip多少天
}

/**如果返回值为true则ip在黑名单中,false 则ip 不在黑名单中*/
func DoFilter(r *http.Request, w http.ResponseWriter) bool {
	switch r.URL.Path {
	case "/login":
		loginCh <- 1 //计算访问loginUrl的危险等级
	default:
		urlCh <- 1 //计算访问其他url的威胁等级
	}
	ip := strings.SplitAfter(r.RemoteAddr, ":")[0]
	ip = strings.TrimSuffix(ip, ":")
	lock.RLock()
	_, inWhite := whiteLists[ip]
	lock.RUnlock()
	if inWhite {
		//在白名单中直接返回false
		fmt.Println("此 ip 在白名单中")
		return false
	}
	mu.RLock()
	_, ok := blackLists[ip]
	mu.RUnlock()

	if ok {
		//此IP在黑名单中
		fmt.Println("doFilter*****************在黑名单中")
		return true

	} else {
		//此IP不在黑名单中
		ipUrl := ip + r.URL.Path
		ipUrls <- IPURL(ipUrl)
		if threatLevel > A { //威胁等级大于A，要采取除了A以外更保险措施
			keywords, badwords := getInfo(r.URL.Path) //获取关键词和敏感词列表
			size := len(r.Form.Encode())              //获得r.form大小
			var isBadIp = false
			for _, word := range keywords {
				if !strings.Contains(r.Form.Encode(), word) {
					//文本内容缺少关键词，所以是恶意ip
					isBadIp = true
					break
				}
			}
			if !isBadIp {
				for _, word := range badwords {
					if strings.Contains(r.Form.Encode(), word) {
						//文本内容有敏感词汇，所以是恶意ip,此处可以考虑放到下一等级措施中，因为这个地方花费也挺多，如果敏感词汇库不大的话可以放这
						isBadIp = true
						break
					}
				}
			}
			if !isBadIp {
				if size > 2000 {
					//文本内容太大，所以是恶意ip
					isBadIp = true
				}
			}
			if isBadIp {
				Add(ip, 10)
				//todo 告知其他服务器
			}
		} else if threatLevel > B && r.Form.Encode() != "" { //威胁等级大于B, 采取一些代价较大的措施
			var percent float64
			text := r.Form.Encode()
			if len(text) > 100 {
				text = text[:100]
			}
			for _, str := range queuelist.list {
				if len(str) > 100 {
					str = str[:100]
				}
				if SimilarText(text, str, &percent); percent > 0.85 {
					//内容很相似，，说明有两个不同ip发送相同内容，，所以判断这两个ip都为恶意ip
					Add(ip, 10)
					//todo 告知其他服务器
					break
				}
			}
			queuelist.list[queuelist.tail] = r.Form.Encode() //将本次内容存入文本队列中
			if queuelist.tail == 49 {
				queuelist.tail = 0 //尾等于 49，需要从头开始存放文本内容
			} else {
				queuelist.tail++
			}
		}
		return false
	}
}

//创建个队列数据结构
type Queuelist struct {
	tail int      //尾index，范围为1..50
	list []string //最近50名访问者访问内容
}

/**主动添加黑名单*/
func Add(ip string, day int) {
	mu.RLock()
	_, ok := blackLists[ip]
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
	if _, ok := whiteLists[ip]; ok {
		fmt.Println("此ip已经在白名单中请不要重复添加")
	} else {
		whiteLists[ip] = true //添加白名单
	}
}

/**计算威胁等级 num: 在线人数*/
func calThreatLevel(num *int) {
	var loginCounters int          //登录url被访问次数
	var urlCounters int            //一般url被访问次数
	var loginLastTime = time.Now() //登录url上次计算频率时间
	var urlLastTime = time.Now()   //一般url上次计算频率时间
	for {
		select {
		case <-loginCh:
			loginCounters++
			if loginCounters > 1000 {
				f := float64(loginCounters) / (time.Now().Sub(loginLastTime)).Seconds() //计算访问频率f
				threatLevel = f / (float64(*num) * getNearest(num, rlogin))             //计算当前威胁等级
				loginCounters, loginLastTime = 0, time.Now()                            //数据清空
			}
		case <-urlCh:
			urlCounters++
			if urlCounters > 1000 {
				f := float64(urlCounters) / (time.Now().Sub(urlLastTime)).Seconds()
				x := f / (float64(*num) * getNearest(num, rurl))
				urlCounters, urlLastTime = 0, time.Now()
				if x > threatLevel {
					threatLevel = x //如果普通url的威胁等级比登录url的威胁等级大，就赋值给当前威胁等级.目的是获得最大的威胁等级
				}
			}
		}
	}
}

/**获取和当前人数匹配的r系数，r系数表示当前人数点击某个url的概率 .
num : 当前人数指针
r : 人数和系数字典
*/
func getNearest(num *int, r map[int]float64) float64 {
	var nearest = float64(10000) //最近距离
	var key int                  // 最近距离时的人数
	for n, _ := range r {
		if dist := math.Abs(float64(*num - n)); dist < nearest {
			nearest = dist
			key = n
		}
	}
	return r[key]
}

/**获取关键词和敏感词列表，，这个需要以后接着弄，需要敏感词库，和各个url对应的关键词列表*/
func getInfo(url string) ([]string, []string) {
	//todo
	return []string{}, []string{}
}

//两个字符串相似度算法  对长度为100个字符的算法时间最差为10ms，好的情况更快
func SimilarText(first, second string, percent *float64) int {
	var similarText func(string, string, int, int) int
	similarText = func(str1, str2 string, len1, len2 int) int {
		var sum, max int
		pos1, pos2 := 0, 0

		// Find the longest segment of the same section in two strings
		for i := 0; i < len1; i++ {
			for j := 0; j < len2; j++ {
				for l := 0; (i+l < len1) && (j+l < len2) && (str1[i+l] == str2[j+l]); l++ {
					if l+1 > max {
						max = l + 1
						pos1 = i
						pos2 = j
					}
				}
			}
		}

		if sum = max; sum > 0 {
			if pos1 > 0 && pos2 > 0 {
				sum += similarText(str1, str2, pos1, pos2)
			}
			if (pos1+max < len1) && (pos2+max < len2) {
				s1 := []byte(str1)
				s2 := []byte(str2)
				sum += similarText(string(s1[pos1+max:]), string(s2[pos2+max:]), len1-pos1-max, len2-pos2-max)
			}
		}

		return sum
	}

	l1, l2 := len(first), len(second)
	if l1+l2 == 0 {
		return 0
	}
	sim := similarText(first, second, l1, l2)
	if percent != nil {
		*percent = float64(sim*200) / float64(l1+l2)
	}
	return sim
}
