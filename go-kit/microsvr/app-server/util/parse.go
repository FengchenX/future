package util

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

//ParseReq 解析请求
func ParseReq(c *gin.Context, funcName string, obj interface{}) error {
	logrus.Infoln(funcName + "************************start")
	body, err := ioutil.ReadAll(c.Request.Body)
	fmt.Println("req data*******************:", string(body))
	if err != nil {
		logrus.Errorln(funcName + "************************读request错误")
		c.JSON(http.StatusOK, "解析错误")
		err := errors.New("解析错误")
		return err
	}

	err = json.Unmarshal(body, obj)
	if err != nil {
		logrus.Errorln(funcName+"********************json 解码错误", err)
		c.JSON(http.StatusOK, "json 解码错误")
		err := errors.New("json 解码错误")
		return err
	}
	return nil
}
