package command

import (
	"fmt"
	"strings"

	"github.com/emicklei/go-restful"

	"grm-service/command"
	"grm-service/dbcentral/etcd"
	"grm-service/dbcentral/pg"
	"grm-service/geoserver"
	"grm-service/service"

	. "storage-manager/dbcentral/etcd"
	. "storage-manager/dbcentral/pg"
	. "storage-manager/storage"
)

type StorageCommand struct {
	command.Meta

	SystemDB     string
	DBuser       string
	Dbpwd        string
	EtcdEndpoint string
	GeoServer    string
	Geouser      string
	Geopwd       string
}

func (c *StorageCommand) Help() string {
	helpText := `
Usage: titan-grm storage-manager [registry_address] [server_address] [server_namespace] [data_dir] [config_dir]
Example: titan-grm storage-manager -registry_address consul:8500 -server_address :8080 -server_namespace titangrm
						-data_dir /opt/titangrm/data -config_dir /opt/titangrm/config
`
	return strings.TrimSpace(helpText)
}

func (c *StorageCommand) Synopsis() string {
	return "GRM Storage Manager"
}

func (c *StorageCommand) Run(args []string) int {
	flags := c.Meta.FlagSet(service.StorageManagerService, command.FlagSetDefault)
	flags.StringVar(&c.SystemDB, "sysdb", "192.168.1.149:31771", "postgresql server address and port")
	flags.StringVar(&c.DBuser, "dbuser", "postgres", "postgresql user")
	flags.StringVar(&c.Dbpwd, "dbpwd", "otitan123", "postgresql user password")
	flags.StringVar(&c.EtcdEndpoint, "etcd endpoint", "192.168.1.149:31686", "etcd endpoint")
	flags.StringVar(&c.GeoServer, "Geo Server", "http://192.168.1.179:8181/grm;http://192.168.1.189:8181/grm", "geo server endpoint")
	flags.StringVar(&c.Geouser, "geouser", "admin", "geoserver user")
	flags.StringVar(&c.Geopwd, "geopwd", "geoserver", "geoserver user password")

	if err := flags.Parse(args); err != nil {
		c.Ui.Error(c.Help())
		return 1
	}
	service := service.NewService(service.StorageManagerService, "v2")
	service.Init(&c.Meta)

	// 初始化数据库
	sysDB, err := pg.ConnectSystemDB(c.SystemDB, c.DBuser, c.Dbpwd)
	if err != nil {
		fmt.Println("Faile to connect system db:", err, c.SystemDB, c.DBuser, c.Dbpwd)
		return 1
	}
	defer sysDB.DisConnect()

	// 初始化etcd连接
	dynamic := DynamicDB{etcd.DynamicEtcd{Endpoints: strings.Split(c.EtcdEndpoint, ";")}}
	if err := dynamic.Connect(); err != nil {
		fmt.Println("Faile to connect etcd v3:", err)
		return 1
	}
	defer dynamic.DisConnect()

	// geoserver
	geoUtil := &geoserver.GeoserverUtil{strings.Split(c.GeoServer, ";"), c.Geouser, c.Geopwd}

	// 添加服务路由
	wc := restful.NewContainer()
	svc := StorageSvc{SysDB: &SystemDB{sysDB}, DynamicDB: &dynamic, GeoServer: geoUtil}
	wc.Add(svc.WebService())

	go svc.DeviceLoop()

	service.Handle("/", wc)
	service.Run()
	return 0
}
