package collector

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"applications/data-collection/dbcentral/pg"
	"applications/data-collection/types"
	"grm-service/geoserver"
	. "grm-service/util"
)

type DataCollectionSvc struct {
	SysDB  *pg.SystemDB
	DataDB *pg.DataDB
	MetaDB *pg.MetaDB

	Device  *types.Device
	DataSet string

	GeoServer *geoserver.GeoserverUtil

	ConfigDir string
}

// WebService creates a new service that can handle REST requests for resources.
func (s DataCollectionSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/collections").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("data collection")}

	// 添加调查
	ws.Route(ws.POST("").To(s.addCollection).
		Doc(TR("new collection")).
		Metadata(api.KeyOpenAPITags, tags).
		//Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.QueryParameter("id", "user id").DataType("string").Required(true)).
		Param(ws.QueryParameter("s", "user session").DataType("string").Required(true)).
		Reads(addCollection{}))

	// 获取调查列表
	ws.Route(ws.GET("").To(s.getCollections).
		Doc(TR("get collections")).
		Metadata(api.KeyOpenAPITags, tags).
		//Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.QueryParameter("id", "user id").DataType("string").Required(true)).
		Param(ws.QueryParameter("s", "user session").DataType("string").Required(true)).
		Writes(types.Colllections{}))

	// 移除调查列表
	ws.Route(ws.DELETE("/{collection-id}").To(s.delCollection).
		Doc(TR("delete collection")).
		Metadata(api.KeyOpenAPITags, tags).
		//Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("collection-id", "collection id").DataType("string").Required(true)).
		Param(ws.QueryParameter("id", "user id").DataType("string").Required(true)).
		Param(ws.QueryParameter("s", "user session").DataType("string").Required(true)))

	// 修改调查信息
	ws.Route(ws.PUT("/{collection-id}").To(s.updateCollection).
		Doc(TR("update collection")).
		Metadata(api.KeyOpenAPITags, tags).
		//Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("collection-id", "collection id").DataType("string").Required(true)).
		Param(ws.QueryParameter("id", "user id").DataType("string").Required(true)).
		Param(ws.QueryParameter("s", "user session").DataType("string").Required(true)).
		Reads(updateCollection{}))

	// 添加调查记录
	ws.Route(ws.PUT("/{collection-id}/datas").To(s.addColData).
		Doc(TR("add collection data")).
		Metadata(api.KeyOpenAPITags, tags).
		//Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("collection-id", "collection id").DataType("string").Required(true)).
		Param(ws.QueryParameter("id", "user id").DataType("string").Required(true)).
		Param(ws.QueryParameter("s", "user session").DataType("string").Required(true)).
		Reads(addColDataReq{}))

	// 移除调查记录
	ws.Route(ws.DELETE("/{collection-id}/datas/{data-id}").To(s.delColData).
		Doc(TR("del collection data")).
		Metadata(api.KeyOpenAPITags, tags).
		//Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("collection-id", "collection id").DataType("string").Required(true)).
		Param(ws.PathParameter("gid", "gid").DataType("string").Required(true)).
		Param(ws.QueryParameter("id", "user id").DataType("string").Required(true)).
		Param(ws.QueryParameter("s", "user session").DataType("string").Required(true)))

	// 修改调查记录
	ws.Route(ws.PUT("/{collection-id}/datas/{data-id}").To(s.updateColData).
		Doc(TR("update collection data")).
		Metadata(api.KeyOpenAPITags, tags).
		//Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("collection-id", "collection id").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Param(ws.QueryParameter("id", "user id").DataType("string").Required(true)).
		Param(ws.QueryParameter("s", "user session").DataType("string").Required(true)).
		Reads(addColDataReq{}))

	// 获取调查记录
	ws.Route(ws.GET("/{collection-id}/datas").To(s.getColData).
		Doc(TR("get collection data")).
		Metadata(api.KeyOpenAPITags, tags).
		//Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("collection-id", "collection id").DataType("string").Required(true)).
		Param(ws.QueryParameter("limit", "limit").DataType("string").Required(true)).
		Param(ws.QueryParameter("offset", "offset").DataType("string").Required(true)).
		Param(ws.QueryParameter("sort", "sort").DataType("string")).
		Param(ws.QueryParameter("order", "order").DataType("string")).
		Param(ws.QueryParameter("id", "user id").DataType("string").Required(true)).
		Param(ws.QueryParameter("s", "user session").DataType("string").Required(true)))

	// 记录搜索
	ws.Route(ws.GET("/{collection-id}/datas/{key-word}").To(s.searchColData).
		Doc(TR("search collection data")).
		Metadata(api.KeyOpenAPITags, tags).
		//Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("collection-id", "collection id").DataType("string").Required(true)).
		Param(ws.PathParameter("key-word", "key world").DataType("string").Required(true)).
		Param(ws.QueryParameter("limit", "limit").DataType("string").Required(true)).
		Param(ws.QueryParameter("offset", "offset").DataType("string").Required(true)).
		Param(ws.QueryParameter("sort", "sort").DataType("string")).
		Param(ws.QueryParameter("order", "order").DataType("string")).
		Param(ws.QueryParameter("id", "user id").DataType("string").Required(true)).
		Param(ws.QueryParameter("s", "user session").DataType("string").Required(true)))

	return ws
}
