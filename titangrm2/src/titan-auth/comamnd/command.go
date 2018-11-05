package command

import (
	"fmt"
	"strings"
	"titan-auth/group"

	"github.com/emicklei/go-restful"

	"grm-service/command"
	"grm-service/dbcentral/etcd"
	"grm-service/dbcentral/pg"
	"grm-service/service"

	. "titan-auth/dbcentral/etcd"
	. "titan-auth/dbcentral/pg"

	"titan-auth/user"
)

type TitanAuthCommand struct {
	command.Meta

	AuthDB       string
	DBuser       string
	DBpwd        string
	EtcdEndpoint string
}

func (c *TitanAuthCommand) Help() string {
	helpText := `
Usage: titan-grm titan-auth [registry_address] [server_address] [server_namespace] [data_dir] [config_dir]
Example: titan-grm titan-auth -registry_address consul:8500 -server_address :8080 -server_namespace titangrm
						-data_dir /opt/titangrm/data -config_dir /opt/titangrm/config
`
	return strings.TrimSpace(helpText)
}

func (c *TitanAuthCommand) Synopsis() string {
	return "Titan Auth Service"
}

func (c *TitanAuthCommand) Run(args []string) int {
	flags := c.Meta.FlagSet(service.TitanAuthService, command.FlagSetDefault)
	flags.StringVar(&c.AuthDB, "authdb", "192.168.1.149:31771", "postgresql server address and port")
	flags.StringVar(&c.DBuser, "dbuser", "postgres", "postgresql user")
	flags.StringVar(&c.DBpwd, "dbpwd", "otitan123", "postgresql user password")
	flags.StringVar(&c.EtcdEndpoint, "etcd endpoint", "192.168.1.149:31686", "etcd endpoint")

	if err := flags.Parse(args); err != nil {
		c.Ui.Error(c.Help())
		return 1
	}
	service := service.NewService(service.TitanAuthService, "v2")
	service.Init(&c.Meta)

	// 初始化数据库
	authDB, err := pg.ConnectAuthDB(c.AuthDB, c.DBuser, c.DBpwd)
	if err != nil {
		fmt.Println("Faile to connect auth db:", err, c.AuthDB, c.DBuser, c.DBpwd)
		return 1
	}
	defer authDB.DisConnect()

	// 初始化etcd连接
	dynamic := DynamicDB{etcd.DynamicEtcd{Endpoints: strings.Split(c.EtcdEndpoint, ";")}}
	if err := dynamic.Connect(); err != nil {
		fmt.Println("Faile to connect etcd v3:", err)
		return 1
	}
	defer dynamic.DisConnect()

	// TODO: 初始化系统auth信息

	// 添加服务路由
	wc := restful.NewContainer()
	userSvc := user.UserSvc{AuthDB: &AuthDB{authDB}, DynamicDB: &dynamic}
	wc.Add(userSvc.WebService())

	groupSvc := group.GroupSvc{
		AuthDB:    &AuthDB{authDB},
		DynamicDB: &dynamic,
	}
	wc.Add(groupSvc.WebService())

	service.Handle("/", wc)
	service.Run()
	return 0
}
