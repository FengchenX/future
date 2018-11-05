package command

import (
	"fmt"
	"strings"

	"github.com/emicklei/go-restful"

	. "data-importer/dbcentral/es"
	. "data-importer/dbcentral/etcd"
	. "data-importer/dbcentral/pg"
	. "data-importer/importer"
	. "data-importer/office"
	. "data-importer/tasks"

	"grm-service/command"
	"grm-service/dbcentral/es"
	"grm-service/dbcentral/etcd"
	"grm-service/dbcentral/pg"
	"grm-service/geoserver"
	"grm-service/mq"
	"grm-service/service"
)

type ImporterCommand struct {
	command.Meta

	SysDB        string
	MetaDB       string
	DBuser       string
	DBpwd        string
	EtcdEndpoint string
	MqUrl        string
	EsUrl        string
	GeoServer    string
	GeoUser      string
	GeoPwd       string
	OfficeServer string
	OfficeUser   string
	OfficePwd    string
}

func (c *ImporterCommand) Help() string {
	helpText := `
Usage: titan-grm data-importer [registry_address] [server_address] [server_namespace] [data_dir] [config_dir]
Example: titan-grm data-importer -registry_address consul:8500 -server_address :8080 -server_namespace titangrm
						-data_dir /opt/titangrm/data -config_dir /opt/titangrm/config
`
	return strings.TrimSpace(helpText)
}

func (c *ImporterCommand) Synopsis() string {
	return "GRM Data Importer Service"
}

func (c *ImporterCommand) Run(args []string) int {
	flags := c.Meta.FlagSet(service.DataImporterService, command.FlagSetDefault)

	flags.StringVar(&c.SysDB, "sysdb", "192.168.1.149:31771", "postgresql server address and port")
	flags.StringVar(&c.MetaDB, "metadb", "192.168.1.149:31771", "postgresql server address and port")
	flags.StringVar(&c.DBuser, "dbuser", "postgres", "postgresql user")
	flags.StringVar(&c.DBpwd, "dbpwd", "otitan123", "postgresql user password")
	flags.StringVar(&c.EtcdEndpoint, "etcd", "192.168.1.149:31686", "etcd endpoint")
	flags.StringVar(&c.MqUrl, "mq", "amqp://admin:otitan123@192.168.1.149:5672/", "rammitmq url")
	flags.StringVar(&c.EsUrl, "es", "http://192.168.1.149:9200", "elasticsearch url")
	flags.StringVar(&c.GeoServer, "geoserver", "http://192.168.1.179:8181/grm;http://192.168.1.189:8181/grm", "geo server endpoint")
	flags.StringVar(&c.GeoUser, "geouser", "admin", "geoserver user")
	flags.StringVar(&c.GeoPwd, "geopwd", "geoserver", "geoserver user password")
	flags.StringVar(&c.OfficeServer, "officeserver", "http://106.74.18.39:8313", "office server")
	flags.StringVar(&c.OfficeUser, "officeuser", "grm@otitan.com", "offie user")
	flags.StringVar(&c.OfficePwd, "officepwd", "123456", "offie user password")

	if err := flags.Parse(args); err != nil {
		c.Ui.Error(c.Help())
		return 1
	}
	service := service.NewService(service.DataImporterService, "v2")
	service.Init(&c.Meta)

	// 初始化数据库
	sysDB, err := pg.ConnectSystemDB(c.SysDB, c.DBuser, c.DBpwd)
	if err != nil {
		fmt.Println("Faile to connect system db:", err, c.SysDB, c.DBuser, c.DBpwd)
		return 1
	}
	defer sysDB.DisConnect()

	metaDB, err := pg.ConnectMetaDB(c.MetaDB, c.DBuser, c.DBpwd)
	if err != nil {
		fmt.Println("Faile to connect meta db:", err, c.MetaDB, c.DBuser, c.DBpwd)
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
		fmt.Printf("Faile to connect mq(%s):%s\n", c.MqUrl, err)
		return 1
	}
	defer mQueue.Close()

	// geoserver
	geoUtil := geoserver.GeoserverUtil{strings.Split(c.GeoServer, ";"), c.GeoUser, c.GeoPwd}

	// office server
	officeUtil := &OnlyOffice{Endpoints: c.OfficeServer, UserName: c.OfficeUser, Password: c.OfficePwd}
	if err := officeUtil.Login(); err != nil {
		fmt.Printf("Faile to login office(%s):%s\n", c.OfficeServer, err)
		return 1
	}

	// 初始化es连接
	es := MetaEs{es.DynamicES{Endpoints: []string{c.EsUrl}}}
	if err := es.Connect(); err != nil {
		fmt.Println("Faile to connect es:", err)
		return 1
	}

	// router
	wc := restful.NewContainer()
	importer_svc := ImporterSvc{
		SysDB:        &SystemDB{sysDB},
		MetaDB:       &MetaDB{metaDB},
		ConfigDir:    c.ConfigDir,
		DynamicDB:    &dynamic,
		MsgQueue:     &mQueue,
		GeoServer:    &geoUtil,
		OfficeServer: officeUtil,
		EsServer:     &es,
		EsUrl:        c.EsUrl,
	}
	wc.Add(importer_svc.WebService())

	// tasks
	tasks_sv := TasksSvc{
		DynamicDB: &dynamic,
		MetaDB:    &MetaDB{metaDB},
		ConfigDir: c.ConfigDir,
	}
	wc.Add(tasks_sv.WebService())

	service.Handle("/", wc)
	service.Run()
	return 0
}
