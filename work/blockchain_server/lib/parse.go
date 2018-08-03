package lib

import (
	"encoding/json"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
)

//ParseReq 解析请求
func ParseReq(c *gin.Context, funcName string, obj interface{}) error {
	glog.Info(funcName + "***************ParseReq,")
	body, err := ioutil.ReadAll(c.Request.Body)
	glog.Infoln("request body:",string(body))
	if err != nil {
		glog.Errorln(funcName + "************************读request错误")
		c.JSON(http.StatusOK, "解析错误")
		err := errors.New("解析错误")
		return err
	}
	err = json.Unmarshal(body, obj)
	if err != nil {
		glog.Errorln(funcName + "********************json 解码错误")
		c.JSON(http.StatusOK, "json 解码错误")
		err := errors.New("json 解码错误")
		return err
	}
	return nil
}
