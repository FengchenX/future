package schedules

import (
	"sub_account_service/order_server/db"
	"sub_account_service/order_server/entity"
	"sub_account_service/order_server/handlers"
	"github.com/golang/glog"
	"time"
)

//add order schedule
func StartAddOrderSchedule() {
	go func(){
		for {
			addOrderSchedule()
			time.Sleep(20 * time.Second)
		}
	}()
}

func addOrderSchedule() {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorln("addOrderSchedule error",err)
		}
	}()
	var orders []*entity.Order
	db.DbClient.Client.Where("complete = ?",0).Find(&orders)
	if orders != nil && len(orders) > 0 {
		for _,order := range orders {
			handlers.SendToNumberServer(order,)
		}
	}
}