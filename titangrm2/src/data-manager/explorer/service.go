package explorer

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	. "grm-service/util"

	"data-manager/dbcentral/etcd"
	"data-manager/dbcentral/pg"
)

type ExplorerSvc struct {
	SysDB     *pg.SystemDB
	DynamicDB *etcd.DynamicDB
}

// WebService creates a new service that can handle REST requests for resources.
func (s ExplorerSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/explorer").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("data explorer")}

	// 获取资源管理
	ws.Route(ws.GET("/").To(s.getExplorer).
		Doc(TR("get explorer of market/user")).
		Param(ws.QueryParameter("mart", "is datamart: true or false").DataType("string").Required(true)).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Metadata(api.KeyOpenAPITags, tags))

	// 更新资源管理
	ws.Route(ws.PUT("/").To(s.updateExplorer).
		Doc(TR("update explorer of market/user")).
		Param(ws.QueryParameter("mart", "is datamart: true or false").DataType("string").Required(true)).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(explorer{}))

	return ws
}
