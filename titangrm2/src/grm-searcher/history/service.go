package history

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"grm-searcher/dbcentral/etcd"
	. "grm-searcher/types"
	. "grm-service/util"
)

type HistorySvc struct {
	DynamicDB *etcd.DynamicDB

	DataDir   string
	ConfigDir string
}

// WebService creates a new service that can handle REST requests for resources.
func (s HistorySvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/history").
		Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("search history management")}

	//history  历史查询，列表，新增，删除  TODO
	// 创建
	ws.Route(ws.POST("").To(s.createHistory).
		Doc(TR("add search history")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(History{}))

	// 获取
	ws.Route(ws.GET("").To(s.getUserHistorys).
		Doc(TR("get user search history list")).
		Param(ws.QueryParameter("user", "identifier of the user").DataType("string").Required(true)).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(HistoryList{}))

	ws.Route(ws.GET("/{id}").To(s.getUserHistory).
		Doc(TR("get one user search history")).
		Param(ws.QueryParameter("user", "identifier of the user").DataType("string").Required(true)).
		Param(ws.PathParameter("id", "identifier of the search history").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(History{}))

	ws.Route(ws.DELETE("/{id}").To(s.deleteUserHistory).
		Doc(TR("delete user dashboard")).
		Param(ws.QueryParameter("user", "identifier of the user").DataType("string").Required(true)).
		Param(ws.PathParameter("id", "identifier of the search history").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags))

	return ws
}
