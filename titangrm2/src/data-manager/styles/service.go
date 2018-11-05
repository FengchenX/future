package styles

import (
	//"data-manager/types"
	"grm-service/common"

	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"grm-service/geoserver"
	. "grm-service/util"

	"data-manager/dbcentral/etcd"
	"data-manager/dbcentral/pg"
)

type StyleSvc struct {
	SysDB     *pg.SystemDB
	DynamicDB *etcd.DynamicDB
	GeoServer *geoserver.GeoserverUtil
}

// WebService creates a new service that can handle REST requests for resources.
func (s StyleSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/styles").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("layer style")}

	// 添加样式
	ws.Route(ws.POST("").To(s.addStyle).
		Doc(TR("add user style")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Reads(addStyleReq{}).Writes(common.LayerStyle{}))

	// 获取用户样式
	ws.Route(ws.GET("").To(s.getStyles).
		Doc(TR("get user styles")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.QueryParameter("style-type", "style type").DataType("string")).
		Writes([]common.LayerStyle{}))

	// 删除样式
	ws.Route(ws.DELETE("/{style-id}").To(s.delStyle).
		Doc(TR("get user styles")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("style-id", "style id").DataType("string").Required(true)))

	// 修改样式
	ws.Route(ws.PUT("/{style-id}").To(s.updateStyle).
		Doc(TR("update user styles")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("style-id", "style id").DataType("string").Required(true)).
		Reads(addStyleReq{}).Writes(common.LayerStyle{}))

	// 获取样式sld
	ws.Route(ws.GET("/{style-id}").To(s.getStyle).
		Doc(TR("get style")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("style-id", "style id").DataType("string").Required(true)).
		Writes(common.LayerStyle{}))

	return ws
}
