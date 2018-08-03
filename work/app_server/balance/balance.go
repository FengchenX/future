package balance

import (
	"encoding/json"
	"fmt"
	"github.com/golang/glog"
	"io"
	"io/ioutil"
	"net/http"
	"sub_account_service/app_server_v2/config"
	"sub_account_service/app_server_v2/db"
	"sub_account_service/app_server_v2/lib"
)

// 获取订单列表
func GetOrderList() ([]db.NumberOrder, error) {
	glog.Infoln("GetOrderList************start")
	//获取最新版本号
	vUrl := fmt.Sprintf("http://%s/getLatestVersion?appId=%s",
		config.ConfInst().OrderAddress, config.ConfInst().AppId)
	glog.Infoln("GetOrderList*******************url:", vUrl)

	body, err := SendHttpRequest("GET", vUrl, nil, nil, nil)
	if err != nil {
		glog.Errorln(lib.Loger("get order version err", "GetOrderListFromServer"), "", err)
		return []db.NumberOrder{}, err
	}

	var vOut map[string]interface{}
	if err := json.Unmarshal(body, &vOut); err != nil {
		glog.Errorln(lib.Loger("unmarshal order version err", "GetOrderListFromServer"), "", err)
		return []db.NumberOrder{}, err
	}

	if vOut["code"] == nil {
		glog.Errorln(lib.Loger("order version code is nil", "GetOrderListFromServer"), "", nil)
		return []db.NumberOrder{}, err
	}
	if vOut["code"].(float64) != 0 {
		glog.Errorln(lib.Loger("get order version code err", "GetOrderListFromServer"), "", fmt.Errorf("code: %v", vOut["code"].(float64)))
		return []db.NumberOrder{}, err
	}

	latestVer := vOut["data"].(map[string]interface{})["latestVersion"].(string)
	if latestVer == "" {
		glog.Infoln("[GetOrderListFromServer]:  latest version is empty")
		return []db.NumberOrder{}, nil
	}
	glog.Infoln("GetOrderList******************version:", latestVer)

	//根据版本号获取订单列表
	orderUrl := fmt.Sprintf("http://%s/orders/batch/?version=%s&appId=%s",
		config.ConfInst().OrderAddress, latestVer, config.ConfInst().AppId)

	glog.Infoln("GetOrderList**********************orderUrl:", orderUrl)

	body, err = SendHttpRequest("GET", orderUrl, nil, nil, nil)
	if err != nil {
		glog.Errorln(lib.Log("get order list err", "", "GetOrderListFromServer"), "", err)
		return []db.NumberOrder{}, err
	}
	glog.Infoln("GetOrderList***********************body:", string(body))
	var oOut map[string]interface{}
	if err := json.Unmarshal(body, &oOut); err != nil {
		glog.Errorln(lib.Log("unmarshal order list err", "", "GetOrderListFromServer"), "", err)
		return []db.NumberOrder{}, err
	}

	if oOut["code"].(float64) != 0 {
		glog.Errorln(lib.Log("get order list code err", "", "GetOrderListFromServer"), "", fmt.Errorf("get order status code=%v", oOut["code"].(float64)))
		return []db.NumberOrder{}, err
	}
	res := []db.NumberOrder{}

	buf, err := json.Marshal(oOut["data"].(map[string]interface{})["orders"])
	if err != nil {
		glog.Errorln("GetOrderList***************json err:", err)
		return nil, nil
	}
	if err := json.Unmarshal(buf, &res); err != nil {
		glog.Errorln("GetOrderList*************************json err:", err)
		return nil, nil
	}
	glog.Infoln("GetOrderList*************************res: ", res)
	return res, nil
}

// 发送http请求
func SendHttpRequest(method, url string, body io.Reader, head, fdata map[string]string) ([]byte, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return []byte{}, err
	}

	cli := http.Client{}
	resp, err := cli.Do(req)
	if err != nil {
		return []byte{}, err
	} else if resp.StatusCode != http.StatusOK {
		return []byte{}, fmt.Errorf("http response status code: %v", resp.StatusCode)
	}
	out, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return []byte{}, err
	}

	return out, nil
}
