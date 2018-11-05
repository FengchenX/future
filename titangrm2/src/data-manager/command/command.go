package command

import (
	"fmt"
	"strings"

	"github.com/emicklei/go-restful"

	. "data-manager/datas"
	. "data-manager/dataset"
	. "data-manager/datatype"
	. "data-manager/explorer"
	. "data-manager/layers"
	. "data-manager/styles"

	. "data-manager/dbcentral/es"
	. "data-manager/dbcentral/etcd"
	. "data-manager/dbcentral/pg"

	"grm-service/command"
	"grm-service/dbcentral/es"
	"grm-service/dbcentral/etcd"
	"grm-service/dbcentral/pg"
	"grm-service/geoserver"
	"grm-service/service"
)

type DataMgrCommand struct {
	command.Meta

	SysDB        string
	MetaDB       string
	DBuser       string
	Dbpwd        string
	EtcdEndpoint string
	EsUrl        string

	GeoServer string
	Geouser   string
	Geopwd    string
}

func (c *DataMgrCommand) Help() string {
	helpText := `
Usage: titan-grm data-manager [registry_address] [server_address] [server_namespace] [data_dir] [config_dir]
Example: titan-grm data-manager -registry_address consul:8500 -server_address :8080 -server_namespace titangrm
						-data_dir /opt/titangrm/data -config_dir /opt/titangrm/config
`
	return strings.TrimSpace(helpText)
}

func (c *DataMgrCommand) Synopsis() string {
	return "GRM Data Manager Service"
}

func (c *DataMgrCommand) Run(args []string) int {
	flags := c.Meta.FlagSet(service.DataManagerService, command.FlagSetDefault)

	flags.StringVar(&c.SysDB, "sysdb", "192.168.1.149:31771", "postgresql server address and port")
	flags.StringVar(&c.MetaDB, "metadb", "192.168.1.149:31771", "postgresql server address and port")
	flags.StringVar(&c.DBuser, "dbuser", "postgres", "postgresql user")
	flags.StringVar(&c.Dbpwd, "dbpwd", "otitan123", "postgresql user password")
	flags.StringVar(&c.EtcdEndpoint, "etcd endpoint", "192.168.1.149:31686", "etcd endpoint")
	flags.StringVar(&c.EsUrl, "es", "http://192.168.1.149:9200", "elasticsearch url")
	flags.StringVar(&c.GeoServer, "Geo Server", "http://192.168.1.179:8181/grm;http://192.168.1.189:8181/grm", "geo server endpoint")
	flags.StringVar(&c.Geouser, "geouser", "admin", "geoserver user")
	flags.StringVar(&c.Geopwd, "geopwd", "geoserver", "geoserver user password")

	if err := flags.Parse(args); err != nil {
		c.Ui.Error(c.Help())
		return 1
	}
	service := service.NewService(service.DataManagerService, "v2")
	service.Init(&c.Meta)

	// 初始化数据库
	sysDB, err := pg.ConnectSystemDB(c.SysDB, c.DBuser, c.Dbpwd)
	if err != nil {
		fmt.Println("Faile to connect system db:", err, c.SysDB, c.DBuser, c.Dbpwd)
		return 1
	}
	defer sysDB.DisConnect()

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

	// 初始化es连接
	//fmt.Println("es url: ", c.EsUrl)
	es := MetaEs{es.DynamicES{Endpoints: []string{c.EsUrl}}}
	if err := es.Connect(); err != nil {
		fmt.Println("Faile to connect es:", err)
		return 1
	}

	// geoserver
	geoUtil := geoserver.GeoserverUtil{strings.Split(c.GeoServer, ";"), c.Geouser, c.Geopwd}

	// router
	wc := restful.NewContainer()
	dataset_svc := DataSetSvc{SysDB: &SystemDB{sysDB}, DynamicDB: &dynamic}
	wc.Add(dataset_svc.WebService())

	datatype_svc := DataTypeSvc{SysDB: &SystemDB{sysDB}}
	wc.Add(datatype_svc.WebService())

	explor_svc := ExplorerSvc{SysDB: &SystemDB{sysDB}, DynamicDB: &dynamic}
	wc.Add(explor_svc.WebService())

	data_svc := DataObjectSvc{SysDB: &SystemDB{sysDB}, MetaDB: &MetaDB{metaDB}, DynamicDB: &dynamic, EsCli: &es, ConfigDir: c.ConfigDir}
	wc.Add(data_svc.WebService())

	layer_svc := DataLayerSvc{SysDB: &SystemDB{sysDB}, MetaDB: &MetaDB{metaDB}, DynamicDB: &dynamic, GeoServer: &geoUtil, ConfigDir: c.ConfigDir}
	wc.Add(layer_svc.WebService())

	style_sv := StyleSvc{SysDB: &SystemDB{sysDB}, DynamicDB: &dynamic, GeoServer: &geoUtil}
	wc.Add(style_sv.WebService())

	service.Handle("/", wc)
	service.Run()
	return 0
}
