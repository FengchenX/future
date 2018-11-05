package datas

import (
	"data-manager/types"

	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"data-manager/dbcentral/es"
	"data-manager/dbcentral/etcd"
	"data-manager/dbcentral/pg"
	. "grm-service/util"
)

type DataObjectSvc struct {
	SysDB     *pg.SystemDB
	MetaDB    *pg.MetaDB
	DynamicDB *etcd.DynamicDB
	EsCli     *es.MetaEs

	ConfigDir string
}

// WebService creates a new service that can handle REST requests for resources.
func (s DataObjectSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/datas").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("data object")}

	// 获取数据信息,标签，摘要，描述
	ws.Route(ws.GET("/{data-id}/info").To(s.getDataInfo).
		Doc(TR("get data info")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Writes(types.DataInfo{}))

	// 修改数据信息：标签，摘要，描述
	ws.Route(ws.PUT("/{data-id}/info").To(s.updateDataInfo).
		Doc(TR("update data info")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Reads(types.UpdateDataInfoReq{}))

	// 获取数据元数据信息
	ws.Route(ws.GET("/{data-id}/metas").To(s.getDataMeta).
		Doc(TR("get data metas")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Writes(types.Meta{}))

	// 修改数据元数据信息
	ws.Route(ws.PUT("/{data-id}/metas").To(s.updateDataMeta).
		Doc(TR("update data meta info")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Reads(types.UpdateDataMetaReq{}))

	// 修改数据快视图
	ws.Route(ws.POST("/{data-id}/snapshot").To(s.updateDataSnapShot).
		Doc(TR("update data snapshot")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Param(ws.QueryParameter("type", "update style: file/base64").DataType("string").Required(true)).
		Param(ws.FormParameter("snapshot", "snapshot").DataType("form-data")).
		Reads(types.SnapshotReq{}))

	// 移除数据
	ws.Route(ws.DELETE("/{data-id}").To(s.delData).
		Doc(TR("delete data object")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)))

	// 获取数据内容：子对象或者浏览地址
	ws.Route(ws.GET("/{data-id}/content").To(s.getDataContent).
		Doc(TR("get data content")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Param(ws.QueryParameter("limit", "limit").DataType("string").Required(false)).
		Param(ws.QueryParameter("offset", "offset").DataType("string").Required(false)).
		Param(ws.QueryParameter("sort", "sort").DataType("string").Required(false)).
		Param(ws.QueryParameter("order", "order").DataType("string").Required(false)).
		Writes(dataContent{}))

	// 添加数据评论
	ws.Route(ws.POST("{data-id}/comments").To(s.addComment).
		Doc(TR("add comment")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Reads(addCommentReq{}))

	// 获取数据评论
	ws.Route(ws.GET("{data-id}/comments").To(s.queryComments).
		Doc(TR("query comments")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Writes(commentListResp{}))

	// 删除数据评论
	ws.Route(ws.DELETE("{data-id}/comments/{comment-id}").To(s.delComment).
		Doc(TR("delete comment")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("data-id", "data id").DataType("string").Required(true)).
		Param(ws.PathParameter("comment-id", "comment id").DataType("string").Required(true)))

	return ws
}
