package command

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"grm-service/command"
	"grm-service/dbcentral/etcd"
	"grm-service/dbcentral/pg"
	"grm-service/service"
	"strings"
	. "tile-manager/dbcentral/etcd"
	. "tile-manager/dbcentral/pg"
	"tile-manager/theme"
)

type TileMgrCommand struct {
	command.Meta

	SysDB  string
	DBuser string
	DBpwd  string

	AuthDB   string
	Authuser string
	Authpwd  string

	EtcdEndpoint string
}

func (c *TileMgrCommand) Help() string {
	helpText := `
Usage: titan-grm tile-manager [registry_address] [server_address] [server_namespace] [data_dir] [config_dir]
Example: titan-grm tile-manager -registry_address consul:8500 -server_address :8080 -server_namespace titangrm
						-data_dir /opt/titangrm/data -config_dir /opt/titangrm/config
`
	return strings.TrimSpace(helpText)
}

func (c *TileMgrCommand) Synopsis() string {
	return "Tile Manager Service"
}

var BaseDir string

func (c *TileMgrCommand) Run(args []string) int {
	flags := c.Meta.FlagSet(service.TileManagerService, command.FlagSetDefault)
	flags.StringVar(&c.SysDB, "sysdb", "192.168.1.149:31771", "postgresql server address and port")
	flags.StringVar(&c.DBuser, "dbuser", "postgres", "postgresql user")
	flags.StringVar(&c.DBpwd, "dbpwd", "otitan123", "postgresql user password")

	flags.StringVar(&c.AuthDB, "authdb", "192.168.1.149:31771", "postgresql server address and port")
	flags.StringVar(&c.Authuser, "authuser", "postgres", "postgresql user")
	flags.StringVar(&c.Authpwd, "authpwd", "otitan123", "postgresql user password")

	flags.StringVar(&c.EtcdEndpoint, "etcd endpoint", "192.168.1.149:31686", "etcd endpoint")

	flags.StringVar(&BaseDir, "basedir", "", "base directory path")

	if err := flags.Parse(args); err != nil {
		c.Ui.Error(c.Help())
		return 1
	}

	service := service.NewService(service.TileManagerService, "v2")
	service.Init(&c.Meta)

	// 初始化数据库
	sysDB, err := pg.ConnectSystemDB(c.SysDB, c.DBuser, c.DBpwd)
	if err != nil {
		fmt.Println("Faile to connect system db:", err, c.SysDB, c.DBuser, c.DBpwd)
		return 1
	}
	defer sysDB.DisConnect()

	// 初始化auth数据库
	authDB, err := pg.ConnectAuthDB(c.AuthDB, c.Authuser, c.Authpwd)
	if err != nil {
		fmt.Println("Fail to connect auth db:", err, c.AuthDB, c.Authuser, c.Authpwd)
	}

	// 初始化etcd连接
	dynamic := DynamicDB{etcd.DynamicEtcd{Endpoints: strings.Split(c.EtcdEndpoint, ";")}}
	if err := dynamic.Connect(); err != nil {
		fmt.Println("Faile to connect etcd v3:", err)
		return 1
	}
	defer dynamic.DisConnect()

	// 添加服务路由
	wc := restful.NewContainer()
	themeSvc := theme.ThemeSvc{
		SysDB:     &SystemDB{sysDB},
		DynamicDB: &dynamic,
		AuthDB:    &AuthDB{authDB},
		BaseDir:   BaseDir,
	}
	wc.Add(themeSvc.WebService())

	service.Handle("/", wc)
	service.Run()
	return 0
}
