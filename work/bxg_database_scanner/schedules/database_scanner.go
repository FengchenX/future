package schedules

import (
	"time"
	"sub_account_service/bxg_database_scanner/config"
	"fmt"
	"sub_account_service/bxg_database_scanner/utils"
	"strconv"
	"io/ioutil"
	"sub_account_service/bxg_database_scanner/service"
	"github.com/golang/glog"
	"encoding/json"
	"bytes"
	"sub_account_service/bxg_database_scanner/models"
	"strings"
	"sync"
	"os"
)

// scanDataBaseSchedules 数据库扫描任务，单线程扫描
func scanDataBaseSchedules() {
	timer :=time.NewTicker(time.Duration(config.GetInstance().ScanInterval * int64(time.Second)))
	for {
		select {
		case <- timer.C:
			func() {
				defer func() {
					if err := recover(); err != nil {
						glog.Error("scanDataBaseSchedules panic:",err)
					}
				}()
				latestOrderNo,getSuccess := getLatestOrderNo()
				if !getSuccess { //query failed
					return
				}
				if latestOrderNo == "" {//从线上获取成功，且为空
					latestOrderNo = INIT_ORDER_ID
				}
				failedOrderNos := getErrorOrderNo()
				failedOrders,err := service.QueryFailedOrders(&failedOrderNos)
				orders,err := service.QueryOrderDetail(latestOrderNo,10)
				orders = append(orders,failedOrders...)
				if err != nil {
					glog.Errorln("query order detail error!",err)
					return
				}
				if len(orders) == 0 {
					return
				}
				save(orders)
			}()
		}
	}
}
var INIT_ORDER_ID = ""
// 获取最后插入的记录
func getLatestOrderNo() (string,bool) {
	latestOrderNo,err := getLatestOrderNoFromServer()
	if err == nil {//from the orderServer get latestOrderNo is blank，retry from local database
		if len(latestOrderNo) == 0 && len(INIT_ORDER_ID) == 0{ //init
			INIT_ORDER_ID,err = service.QueryLatestOrderNo()
			if err != nil {
				INIT_ORDER_ID = ""
				return "", false
			}
			return "", true
		} else {
			return latestOrderNo, true
		}
	} else {
		glog.Errorln("getLatestOrderNo from orderServer err:",err)
		return "",false
	}
}
//getLatestOrderNoFromServer from orderServer get the latest order no
func getLatestOrderNoFromServer() (string,error){
	reqUrl := getRequestUrl(config.GetInstance().GetLatestOrderNoUrl)
	item := map[string]string{
		"Company":config.GetInstance().Company,
		"BranchShop":config.GetInstance().BranchShop,
	}
	out,err := json.Marshal(item)
	if err != nil {
		glog.Errorln("parse getLatestOrderNo req param err:",err,",req param:",item)
		return "", err
	}
	resp,err := utils.Post(reqUrl,bytes.NewReader(out),nil,nil)
	if err == nil { //访问没有错误
		result := models.BatchAddResp{}
		if err := json.Unmarshal([]byte(resp), &result); err == nil {
			if result.Code == 0 {
				return result.Data, nil
			}
		}else { //解码错误
			return "", err
			glog.Errorln("batch add error!")
		}
	}
	return "", err
}

var errFileLock = sync.Mutex{}
var hasErrorOrderNo = false
var errorOrderNoFile = "./conf/error"
func getErrorOrderNo() []string {
	errFileLock.Lock()
	defer errFileLock.Unlock()
	if !hasErrorOrderNo {
		return make([]string,0)
	}
	bytes1, _ := ioutil.ReadFile(errorOrderNoFile)
	str := string(bytes1)
	str = strings.Trim(str, ",")
	hasErrorOrderNo = false
	return strings.Split(str,",")
}
func recordErrorOrderNo(orderNo *[]string) {
	errFileLock.Lock()
	defer errFileLock.Unlock()
	str := strings.Join(*orderNo,",")
	err := ioutil.WriteFile(errorOrderNoFile,[]byte(","+str),os.ModeSticky)
	hasErrorOrderNo = true
	if err != nil {
		glog.Error("record to errFile fail",err)
	}
}

func getRequestUrl(url string) string{
	timeStamp := time.Now().Unix() * 1000
	sign := utils.MD5(strconv.Itoa(int(timeStamp)) + config.GetInstance().DevelopKey)
	return fmt.Sprintf("%v?AppId=%v&Timestamp=%v&Sign=%v",
		url,config.GetInstance().AppId, timeStamp, sign)
}
//保存
func save(orders []*models.BillOrderSaveReq) {
	failedOrderNo := make([]string,0) // failed order no
	defer func() {
		if err := recover(); err != nil {
		}
		if len(failedOrderNo) != len(orders) { //means some failed，record failed values
			recordErrorOrderNo(&failedOrderNo)
		}
	}()
	reqUrl := getRequestUrl(config.GetInstance().AddOrderUrl)
	for _,item := range orders {
		out,err := json.Marshal(item)
		if err != nil {
			glog.Errorln("parse order error",err)
			return
		}
		resp,err := utils.Post(reqUrl,bytes.NewReader(out),nil,nil)
		if err == nil { //访问没有错误
			result := models.BatchAddResp{}
			if err := json.Unmarshal([]byte(resp), &result); err == nil {
				if result.Code == 0 || result.Code == 10001 { //处理成功,记录处理到哪个记录
				} else if result.Code == -1 { //记录处理错误的订单编号
					failedOrderNo = append(failedOrderNo,item.OrderNo)
				}
			}else { //解码错误
				failedOrderNo = append(failedOrderNo,item.OrderNo)
				glog.Errorln("batch add error!")
			}
		}else {
			failedOrderNo = append(failedOrderNo,item.OrderNo)
		}
	}
}