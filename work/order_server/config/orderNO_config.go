package config

import "sub_account_service/order_server/utils"

type OrderNoConfig struct {
	Host     string //host
	Port     string //端口
	AddOrder string //添加订单
}

var orderNoConfigs OrderNoConfig

func GetOrderNoConfig() OrderNoConfig {
	return orderNoConfigs
}
func init() {
	utils.LoadToml("./conf/orderNO_config.toml", &orderNoConfigs)
	orderNoConfigs.jointRequestUrl() //初始化拼接
}

func (config *OrderNoConfig) jointRequestUrl() {
	prefix := config.Host + ":" + config.Port + "/"
	config.AddOrder = prefix + config.AddOrder
}
