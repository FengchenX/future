package config

// Config 配置类型
type Config struct {
	Producer_Server_Http string
	Query_Server_Http    string
	Mysql                string
	Nsqd_Producer_Tcp    string
	Nsqd_Consumer_Tcp    string
}

// 默认配置
var optional = Config{
	//生产者服务器接口端点 可用nginx做一层代理
	//Producer_Server_Http:"192.168.83.79:9897",
	Producer_Server_Http: "172.18.22.48:9897",
	//Producer_Server_Http: "127.0.0.1:9897",
	//查询服务器接口端点
	//Query_Server_Http:"192.168.83.79:9898",
	Query_Server_Http: "172.18.22.48:9898",
	//Query_Server_Http: "127.0.0.1:9898",
	//mysql链接的url
	//Mysql:"root:root@tcp(39.108.80.66:3306)/test_order?charset=utf8&parseTime=true&loc=Local",
	//Mysql:"launch:root@tcp(192.168.83.79:3306)/number_server?charset=utf8&parseTime=true&loc=Local",
	//Mysql: "root:12345678@tcp(127.0.0.1:3306)/finance_orders?charset=utf8&parseTime=true&loc=Local",

	Mysql: "isbs:isbs*2018@tcp(rm-wz961aqvp3aq3gs334o.mysql.rds.aliyuncs.com:3306)/number_server?charset=utf8&parseTime=true&loc=Local",
	//nsq 可以部署成分布式的接口 后期根据量来扩容吧
	//消息队列nsq的生产者端点
	//Nsqd_Producer_Tcp:"192.168.83.79:4150",
	Nsqd_Producer_Tcp: "172.18.22.48:4150",
	//Nsqd_Producer_Tcp: "127.0.0.1:4150",
	//消息队列nsq的消费者端点
	//Nsqd_Consumer_Tcp:"192.168.83.79:4150",
	Nsqd_Consumer_Tcp: "172.18.22.48:4150",
	//Nsqd_Consumer_Tcp: "127.0.0.1:4150",
}

// Opts 获取配置
func Opts() Config {
	return optional
}
