package config

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"grm-labelmgr/dbcentral/etcd"
	. "grm-labelmgr/types"
	. "grm-service/util"
)

type ConfigSvc struct {
	DynamicDB *etcd.DynamicDB

	DataDir   string
	ConfigDir string
}

// WebService creates a new service that can handle REST requests for resources.
func (s ConfigSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("user data management")}

	// 创建
	ws.Route(ws.POST("/user/{user}/data").To(s.createLayerConfig).
		Param(ws.PathParameter("user", "identifier of the user").DataType("string")).
		Doc(TR("add user data")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(UserData{}))

	// 获取
	ws.Route(ws.GET("/user/{user}/data/{data_type}/{data_name}").To(s.getLayerConfig).
		Doc(TR("get user data")).
		Param(ws.PathParameter("user", "identifier of the user").DataType("string")).
		Param(ws.PathParameter("data_type", "data type").DataType("string")).
		Param(ws.PathParameter("data_name", "data name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(UserData{}))

	ws.Route(ws.GET("/user/{user}/data/{data_type}").To(s.getUserDatas).
		Doc(TR("get user data list by type")).
		Param(ws.PathParameter("user", "identifier of the user").DataType("string")).
		Param(ws.PathParameter("data_type", "data type").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(UserDataList{}))

	ws.Route(ws.DELETE("/user/{user}/data/{data_type}/{data_name}").To(s.deleteLayerConfig).
		Doc(TR("delete user data")).
		Param(ws.PathParameter("user", "identifier of the user").DataType("string")).
		Param(ws.PathParameter("data_type", "data type").DataType("string")).
		Param(ws.PathParameter("data_name", "data name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags))
	return ws
}
