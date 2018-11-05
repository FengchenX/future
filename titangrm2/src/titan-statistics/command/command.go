package command

import (
	"fmt"
	"strings"

	"github.com/emicklei/go-restful"
	"grm-service/command"
	"grm-service/dbcentral/etcd"
	"grm-service/dbcentral/pg"
	"grm-service/mq"
	"grm-service/service"

	. "titan-statistics/aggregation"
	. "titan-statistics/dbcentral/etcd"
	. "titan-statistics/dbcentral/pg"
	. "titan-statistics/statistics"
)

type StatCommand struct {
	command.Meta

	SystemDB     string
	MetaDB       string
	DBuser       string
	Dbpwd        string
	EtcdEndpoint string
	EsUrl        string
	MqUrl        string
}

func (c *StatCommand) Help() string {
	helpText := `
Usage: titan-grm statistics [registry_address] [server_address] [server_namespace] [data_dir] [config_dir]
Example: titan-grm statistics -registry_address consul:8500 -server_address :8080 -server_namespace titangrm
						-data_dir /opt/titangrm/data -config_dir /opt/titangrm/config
`
	return strings.TrimSpace(helpText)
}

func (c *StatCommand) Synopsis() string {
	return "titan statistics service"
}

func (c *StatCommand) Run(args []string) int {
	flags := c.Meta.FlagSet(service.StatService, command.FlagSetDefault)
	flags.StringVar(&c.SystemDB, "sysdb", "192.168.1.149:31771", "postgresql server address and port")
	flags.StringVar(&c.MetaDB, "metadb", "192.168.1.149:31771", "postgresql server address and port")

	flags.StringVar(&c.DBuser, "dbuser", "postgres", "postgresql user")
	flags.StringVar(&c.Dbpwd, "dbpwd", "otitan123", "postgresql user password")
	flags.StringVar(&c.EtcdEndpoint, "etcd endpoint", "192.168.1.149:31686", "etcd endpoint")
	flags.StringVar(&c.EsUrl, "es_url", "http://192.168.1.149:9200", "es url")
	flags.StringVar(&c.MqUrl, "mq", "amqp://admin:otitan123@192.168.1.149:5672/", "rammitmq url")

	if err := flags.Parse(args); err != nil {
		c.Ui.Error(c.Help())
		return 1
	}
	service := service.NewService(service.StatService, "v2")
	service.Init(&c.Meta)

	// 初始化数据库
	sysDB, err := pg.ConnectSystemDB(c.SystemDB, c.DBuser, c.Dbpwd)
	if err != nil {
		fmt.Println("Faile to connect system db:", err, c.SystemDB, c.DBuser, c.Dbpwd)
		return 1
	}
	defer sysDB.DisConnect()

	// 初始化数据库
	metaDB, err := pg.ConnectMetaDB(c.MetaDB, c.DBuser, c.Dbpwd)
	if err != nil {
		fmt.Println("Faile to connect meta db:", err, c.MetaDB, c.DBuser, c.Dbpwd)
		return 1
	}
	defer metaDB.DisConnect()

	// 初始化etcd连接
	dynamic := DynamicDB{etcd.DynamicEtcd{Endpoints: strings.Split(c.EtcdEndpoint, ";")}}
	if err := dynamic.Connect(); err != nil {
		fmt.Println("Faile to connect etcd v3:", err)
		return 1
	}
	defer dynamic.DisConnect()

	// 初始化消息队列
	mQueue := mq.RabbitMQ{URL: c.MqUrl}
	if err := mQueue.Connect(); err != nil {
		fmt.Print("Faile to connect mq(%s):%s\n", c.MqUrl, err)
		return 1
	}
	defer mQueue.Close()

	// 添加服务路由
	wc := restful.NewContainer()

	statSvc := StatSvc{SysDB: &SystemDB{sysDB}, MetaDB: &MetaDB{metaDB}, DynamicDB: &dynamic}
	wc.Add(statSvc.WebService())

	aggrSvc := AggrSvc{SysDB: &SystemDB{sysDB}, MetaDB: &MetaDB{metaDB}, DynamicDB: &dynamic}
	wc.Add(aggrSvc.WebService())

	service.Handle("/", wc)
	err = service.Run()
	if err != nil {
		fmt.Println("Faile to run service:", err)
		return 1
	}
	return 0
}
