package command

import (
	"fmt"
	"strings"

	"github.com/emicklei/go-restful"

	. "applications/data-collection/collector"
	. "applications/data-collection/dbcentral/pg"

	"grm-service/command"
	"grm-service/dbcentral/pg"
	"grm-service/geoserver"
	"grm-service/service"
)

type CollectorCommand struct {
	command.Meta

	SystemDB     string
	SystemDBUser string
	SystemDBPwd  string

	DataSet    string
	MetaDB     string
	MetaDBUser string
	MetaDBPwd  string

	GeoServer string
	GeoUser   string
	GeoPwd    string

	//EtcdEndpoint string
}

func (c *CollectorCommand) Help() string {
	helpText := `
Usage: titan-grm data-collection [registry_address] [server_address] [server_namespace] [data_dir] [config_dir]
Example: titan-grm data-collection -registry_address consul:8500 -server_address :8080 -server_namespace titangrm
						-data_dir /opt/titangrm/data -config_dir /opt/titangrm/config
`
	return strings.TrimSpace(helpText)
}

func (c *CollectorCommand) Synopsis() string {
	return "GRM Data Collection"
}

func (c *CollectorCommand) Run(args []string) int {
	flags := c.Meta.FlagSet(service.DataCollection, command.FlagSetDefault)
	flags.StringVar(&c.SystemDB, "sysdb", "192.168.1.189:5432", "postgresql server address and port")
	flags.StringVar(&c.SystemDBUser, "sysdbuser", "postgres", "postgresql user")
	flags.StringVar(&c.SystemDBPwd, "sysdbpwd", "123456", "postgresql user password")
	flags.StringVar(&c.DataSet, "dataset", "FFDCF439BD96C0A4C023BEA15A9D264A", "collection dataset")

	flags.StringVar(&c.MetaDB, "metadb", "192.168.1.189:5432", "postgresql server address and port")
	flags.StringVar(&c.MetaDBUser, "metadbuser", "postgres", "postgresql user")
	flags.StringVar(&c.MetaDBPwd, "metadbpwd", "123456", "postgresql user password")

	//flags.StringVar(&c.EtcdEndpoint, "etcd endpoint", "192.168.1.149:31686", "etcd endpoint")
	flags.StringVar(&c.GeoServer, "geoserver", "http://192.168.1.179:8181/grm;http://192.168.1.189:8181/grm", "geo server endpoint")
	flags.StringVar(&c.GeoUser, "geouser", "admin", "geoserver user")
	flags.StringVar(&c.GeoPwd, "geopwd", "geoserver", "geoserver user password")

	if err := flags.Parse(args); err != nil {
		c.Ui.Error(c.Help())
		return 1
	}
	service := service.NewService(service.DataCollection, "v2")
	service.Init(&c.Meta)

	if len(c.DataSet) == 0 {
		fmt.Println("Invalid dataset for data collection")
		return 1
	}

	// geoserver
	geoUtil := geoserver.GeoserverUtil{strings.Split(c.GeoServer, ";"), c.GeoUser, c.GeoPwd}

	// 初始化数据库
	sysDB, err := pg.ConnectSystemDB(c.SystemDB, c.SystemDBUser, c.SystemDBPwd)
	if err != nil {
		fmt.Println("Faile to connect system db:", err, c.SystemDB, c.SystemDBUser, c.SystemDBPwd)
		return 1
	}
	defer sysDB.DisConnect()

	// meta
	metaDB, err := pg.ConnectMetaDB(c.MetaDB, c.MetaDBUser, c.MetaDBPwd)
	if err != nil {
		fmt.Println("Faile to connect system db:", err, c.MetaDB, c.MetaDBUser, c.MetaDBPwd)
		return 1
	}
	defer metaDB.DisConnect()

	// data
	sysDb := &SystemDB{sysDB}
	dev, err := sysDb.GetDataDevice()
	if err != nil {
		fmt.Println("Faile to get data db:", err)
		return 1
	}

	host := dev.IpAddress + ":" + dev.DBPort
	fmt.Println("data db: ", host, dev.DBUser, dev.DBPwd)
	fmt.Println("data dataset: ", c.DataSet)
	dataDB, err := pg.ConnectDataDB(host, dev.DBUser, dev.DBPwd)
	if err != nil {
		fmt.Println("Faile to connect data db:", err, host, dev.DBUser, dev.DBPwd)
		return 1
	}
	defer dataDB.DisConnect()

	// 添加服务路由
	wc := restful.NewContainer()
	svc := DataCollectionSvc{
		SysDB:     sysDb,
		MetaDB:    &MetaDB{metaDB},
		DataDB:    &DataDB{dataDB},
		Device:    dev,
		DataSet:   c.DataSet,
		GeoServer: &geoUtil,
	}
	wc.Add(svc.WebService())

	service.Handle("/", wc)
	service.Run()
	return 0
}
