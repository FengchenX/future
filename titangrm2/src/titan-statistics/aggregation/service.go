package aggregation

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	. "grm-service/dbcentral/pg"
	. "grm-service/util"
	"titan-statistics/dbcentral/etcd"
	"titan-statistics/dbcentral/pg"
	. "titan-statistics/types"
)

type AggrSvc struct {
	SysDB     *pg.SystemDB
	MetaDB    *pg.MetaDB
	DynamicDB *etcd.DynamicDB

	DataDir   string
	ConfigDir string

	DataIdConns map[string]*ConnConfig
}

// WebService creates a new service that can handle REST requests for resources.
func (s AggrSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/aggr").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("aggregation info manage")}

	//聚合条件
	ws.Route(ws.POST("/data_type/{type}").To(s.SetAggr).
		Doc(TR("aggregation by data type")).
		Param(ws.PathParameter("type", "data type name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(TypeAggr{}))

	ws.Route(ws.GET("/data_type/{type}").To(s.GetAggr).
		Doc(TR("aggregation by data type")).
		Param(ws.PathParameter("type", "data type name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(TypeAggr{}))

	//获取这个类型的聚合条件下的值列表
	//	ws.Route(ws.GET("/data_type/{type}/field/{field}/stat").To(s.StatByAggr).
	//		Doc(TR("aggregation by data type")).
	//		Param(ws.PathParameter("type", "data type name").DataType("string")).
	//		Param(ws.PathParameter("field", "stat filed").DataType("string")).
	//		Metadata(api.KeyOpenAPITags, tags).
	//		Writes(DistinctValues{}))

	return ws
}
