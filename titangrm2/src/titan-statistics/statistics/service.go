package statistics

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	. "grm-service/dbcentral/pg"
	. "grm-service/util"
	"titan-statistics/dbcentral/etcd"
	"titan-statistics/dbcentral/pg"
	. "titan-statistics/types"
)

type StatSvc struct {
	SysDB     *pg.SystemDB
	MetaDB    *pg.MetaDB
	DynamicDB *etcd.DynamicDB

	DataDir   string
	ConfigDir string

	DataIdConns map[string]*ConnConfig
}

// WebService creates a new service that can handle REST requests for resources.
func (s StatSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/stat").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("statistic info manage")}

	//全局
	//按大类统计
	ws.Route(ws.GET("/data_type").To(s.dataTypeStat).
		Doc(TR("statistic by data type")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(TypeStatList{}))

	//按小类统计
	ws.Route(ws.GET("/data_type/{type}/sub_types").To(s.subTypeStat).
		Doc(TR("statistic by data sub type")).
		Param(ws.PathParameter("type", "data type name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(SubtypeList{}))

	//按接口统计
	//	ws.Route(ws.GET("/api").To(s.apiStat).
	//		Doc(TR("statistic by api request")).
	//		Metadata(api.KeyOpenAPITags, tags).
	//		Writes(TypeStatList{}))

	//个人
	//按大类统计
	ws.Route(ws.GET("/user/{user_id}/data_type").To(s.userDataTypeStat).
		Doc(TR("statistic by user data type")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(TypeStatList{}))

	//按小类统计
	ws.Route(ws.GET("/user/{user_id}/data_type/{type}/sub_types").To(s.userSubTypeStat).
		Doc(TR("statistic by user data sub type")).
		Param(ws.PathParameter("user_id", "user id").DataType("string")).
		Param(ws.PathParameter("type", "data type name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(SubtypeList{}))

	//数据的浏览次数

	//数据的下载次数

	return ws
}
