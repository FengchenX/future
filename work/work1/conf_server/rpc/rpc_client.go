package rpc

import (
	"github.com/golang/glog"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"sync"
	"znfz/conf_server/config"
	"znfz/conf_server/db"
	"znfz/conf_server/protocol"
)

func ReloadDeploy(addr, acco string) {
	agents := make([]*db.RpcHistory, 0)
	db.DbClient.Client.Find(agents)
	var wg sync.WaitGroup

	for num := range agents {
		wg.Add(1)
		agent := agents[num]
		go func() {
			conn, err := grpc.Dial(agent.Address, grpc.WithInsecure())
			defer func() {
				conn.Close()
				wg.Done()
			}()

			if err != nil {
				glog.Errorln("Can't connect: "+agent.Address, err)
			}
			c := protocol.NewApiServiceClient(conn)
			c.ReloadDeploy(context.Background(), &protocol.CReloadDeploy{
				Address: addr,
				Account: acco,
			})
		}()
	}
}

func ReloadConfig() {
	agents := make([]*db.RpcHistory, 0)
	db.DbClient.Client.Find(agents)
	var wg sync.WaitGroup

	for num := range agents {
		wg.Add(1)
		agent := agents[num]
		go func() {
			conn, err := grpc.Dial(agent.Address, grpc.WithInsecure())
			defer func() {
				conn.Close()
				wg.Done()
			}()

			if err != nil {
				glog.Errorln("Can't connect: "+agent.Address, err)
			}
			c := protocol.NewApiServiceClient(conn)
			c.ReloadConfig(context.Background(), &protocol.CReloadConfig{
				OperateTimeout: uint32(config.Opts().Operate_timeout),
				LocalAddress:   config.Optional.LocalAddress,
				Port:           config.Optional.Port,
				AccAddress:     config.Optional.AccAddress,
				EthAddress:     config.Optional.EthAddress,
				IpcDir:         config.Optional.IpcDir,
				ServerId:       config.Optional.ServerId,
				ManagerKey:     config.Optional.ManagerKey,
				ManagerPhrase:  config.Optional.ManagerPhrase,
				KeyDir:         config.Optional.KeyDir,
			})
		}()
	}
}
