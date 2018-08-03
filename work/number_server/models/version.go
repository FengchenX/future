package models

import (
	"github.com/golang/glog"
	"sub_account_service/finance_order_server/models"
	"time"
)

type QueryVersions struct {
	ID           int       `json:"id"`
	AppId        string    `json:"app_id"`         //appId
	CompanyId    int        `json:"company_id"`
	Version      string    `json:"version"`        //版本号
	BeginOrderId int       `json:"begin_order_id"` //该次查询起始的ID
	EndOrderId   int       `json:"end_order_id"`   //该次查询结束的ID
	CreateTime   time.Time `json:"create_time"`
}

type ServiceDevelopKey struct {
	KeyId      int       `json:"key_id"`
	AppId      string    `json:"app_id"`      //开发的appid
	CompanyId    int      `json:"company_id"`
	DevelopKey string    `json:"develop_key"` // 开发的key
	Service    string    `json:"service"`     //服务名称
	CreateTime time.Time `json:"create_time"` //创建时间
}

func LegalAppId(appId string) bool {
	var count int
	var models = &[]ServiceDevelopKey{}
	db.Where(map[string]interface{}{"app_id": appId}).Find(&models).Count(&count)
	if count > 0 {
		return true
	}
	return false
}

func GetLatestVersion(appId string) *QueryVersions { //获取最后一个版本号
	latestVersion := &QueryVersions{}
	db.Where("app_id = ?", appId).Last(latestVersion) //获取该appId下最后一个
	return latestVersion
}

func GetVersionByAppIdAndVersion(appId string, v string) *QueryVersions {
	currentVesion := &QueryVersions{}
	db.Where("app_id = ? AND version = ?", appId, v).First(currentVesion)
	return currentVesion
}

func GetVersionsByAppIdAndID(id int, appId string) *QueryVersions {
	nextVersion := &QueryVersions{}
	db.Where("id > ? AND app_id = ?", id, appId).First(nextVersion)
	return nextVersion
}

func IsVersionExist(appId string, version string) bool {
	var count int
	db.Model(&models.QueryVersions{}).Where(" app_id = ? and version = ?", appId, version).Count(&count)
	return count > 0
}

func CreateVersion(query *QueryVersions) bool {
	createDB := db.Create(query)
	if createDB.Error != nil {
		glog.Error("insert query version error，", createDB.Error)
		return false
	}
	return true
}
