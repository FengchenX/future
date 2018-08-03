package api

import "sub_account_service/fabric_server/sdk"

// api服务结构体，必须实现了所有service ApiService中的方法
type ApiService struct {
	Fabric *sdk.FabricSetup
}
