package manager

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/golang/glog"
	"io/ioutil"
	"net/http"
	"time"
	"znfz/web_server/lib"
	"znfz/web_server/client"
	"znfz/web_server/db"
	"znfz/web_server/models"
)

type ReqBinder struct {
	CompanyName  string
	BranchName   string
	UseBanchName bool
	UserAddress  string
	Phone        string
}

// apply company account
func Setbindmsg(c *gin.Context, cli *client.Client) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	req := &ReqBinder{}
	err = json.Unmarshal(body, req)
	if err != nil {
		c.String(http.StatusOK, "error")
	}

	company := &models.Company{}
	// get company from mysql
	dberr := db.DbClient.Client.Where(&models.Company{
		CompanyName:      req.CompanyName,
		CompanyBanchName: req.BranchName,
	}).First(company)
	if dberr.Error != nil {
		glog.Error(dberr.Error)
	}
	glog.Infoln(lib.Log("api", "", "Setbindmsg:company"), company)

	// save it to mysql
	binder := &models.Binder{
		UserAddress:    req.UserAddress,
		CompanyName:    req.CompanyName,
		BrenchName:     req.BranchName,
		CompanyId:      company.ID,
		CompanyAddress: company.CompanyAddress,
		BindTime:       time.Now(),
	}
	dberr = db.DbClient.Client.Create(binder)
	if dberr.Error != nil {
		glog.Errorln(dberr.Error)
	}
	glog.Infoln(lib.Log("api", "", "Setbindmsg:binder"), binder)
}

type ReqBindMsg struct {
	UserAddress  string
	Phone        string
}

// apply company account
func Getbindmsg(c *gin.Context, cli *client.Client) {
	body, err := ioutil.ReadAll(c.Request.Body)
	if err != nil {
		c.String(http.StatusOK, "error")
	}
	req := &ReqBindMsg{}
	err = json.Unmarshal(body, req)
	if err != nil {
		c.String(http.StatusOK, "error")
	}

	// find it in mysql
	bind := &models.Binder{}
	dberr:=db.DbClient.Client.Where(&models.Binder{
		UserAddress: req.UserAddress,
	}).First(bind)
	if dberr.Error!=nil{
		glog.Errorln(dberr.Error)
	}
	glog.Infoln(lib.Log("api", "", "Getbindmsg"), bind)
}
