package layers

import (
	"data-manager/types"
	"grm-service/common"

	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"grm-service/geoserver"
	. "grm-service/util"

	"data-manager/dbcentral/etcd"
	"data-manager/dbcentral/pg"
	//. "data-manager/types"
)

type DataLayerSvc struct {
	SysDB     *pg.SystemDB
	MetaDB    *pg.MetaDB
	DynamicDB *etcd.DynamicDB
	GeoServer *geoserver.GeoserverUtil

	ConfigDir string
}

// WebService creates a new service that can handle REST requests for resources.
func (s DataLayerSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/layers").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("data layer")}

	// 添加数据图层
	ws.Route(ws.POST("/{data-id}").To(s.addLayer).
		Doc(TR("add layer")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Reads(addLayerReq{}).Writes(common.DataLayer{}))

	// 获取用户数据图层
	ws.Route(ws.GET("/{data-id}").To(s.getLayers).
		Doc(TR("get data layers")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Writes([]*common.DataLayer{}))

	// 移除数据图层
	ws.Route(ws.DELETE("/{data-id}/{layer-id}").To(s.delLayer).
		Doc(TR("delete user layer")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("layer-id", "layer id").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)))

	// 修改图层信息
	ws.Route(ws.PUT("/{data-id}/{layer-id}").To(s.updateLayer).
		Doc(TR("update user layer")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("layer-id", "layer id").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Reads(types.UpdateLayerReq{}))

	// 获取指定图层信息
	ws.Route(ws.GET("/{data-id}/{layer-id}").To(s.getLayer).
		Doc(TR("get user layer")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Param(ws.PathParameter("layer-id", "layer id").DataType("string").Required(true)).
		Writes(common.DataLayer{}))

	// 修改图层缩略图
	ws.Route(ws.POST("/{data-id}/{layer-id}/snapshot").To(s.updateDataSnapShot).
		Doc(TR("update layer snapshot")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Param(ws.PathParameter("layer-id", "layer id").DataType("string").Required(true)).
		Param(ws.QueryParameter("type", "update style: file/base64").DataType("string").Required(true)).
		Param(ws.FormParameter("snapshot", "snapshot").DataType("form-data")).
		Reads(types.SnapshotReq{}))

	return ws
}
