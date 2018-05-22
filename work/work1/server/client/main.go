package main

import (
	"flag"
	"znfz/server/config"
)

func main() {
	var path string
	flag.StringVar(&path, "config", "./config.toml", "config path")
	flag.Parse()
	config.ParseToml(path) // 初始化配置
	run()
}