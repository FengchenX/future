package main

import (
	"sub_account_service/process_monitor/monitor"
	"sub_account_service/process_monitor/router"
	"common-utilities/utilities"
	"time"
)

func main() {
	go monitor.StartMonitor()
	utilities.InitLogrus("./log","process_monitor", 24 * 365 * time.Hour, 24 * time.Hour)
	router.Start()
}

