package config

import (
	"sub_account_service/order_server/utils"
	"time"

)

// Config 配置类型
type Config struct {
	Port         string // 监听端口

	ReTransferTime int    // 定时重新转账时间（小时）
	PaySwitch      int    // 是否转账开关
	AliPayAddress  string // 收款方账号
	AliPayName     string // 收款方用户名

	TimeTicker2      time.Duration     // 从数据库获取 添加到编号系统失败 的交易流水，重新添加

	MysqlStr    string
	AutoMigrate bool
}

// 配置
var config Config

// Opts 获取配置
func GetConfig() Config {
	return config
}

func init() {
	utils.LoadToml("./conf/config.toml", &config)
}
