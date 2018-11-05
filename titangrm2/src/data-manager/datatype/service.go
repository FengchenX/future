package datatype

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	. "grm-service/util"

	"data-manager/dbcentral/pg"
	. "data-manager/types"
)

type DataTypeSvc struct {
	SysDB *pg.SystemDB
}

// WebService creates a new service that can handle REST requests for resources.
func (s DataTypeSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/types").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("data type")}

	// 获取所有数据类型集合
	ws.Route(ws.GET("/").To(s.getDatatypeList).
		Doc(TR("get all data types")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes([]DataType{}))

	// 编辑类型基本信息
	ws.Route(ws.PUT("/{type-name}").To(s.updateTypeInfo).
		Doc(TR("update information of data type")).
		Param(ws.PathParameter("type-name", "type name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(updateTypeInfoReq{}))

	// 获取类型元元数据
	ws.Route(ws.GET("/meta/{type-name}").To(s.getTypeMeta).
		Doc(TR("get meta of data type")).
		Param(ws.PathParameter("type-name", "type name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(Meta{}))

	// 修改元元数据信息字段
	ws.Route(ws.PUT("/meta/{type-name}").To(s.updateMetaField).
		Doc(TR("update meta field")).
		Param(ws.PathParameter("type-name", "type name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(MetaFieldReq{}))

	// 添加元元数据信息字段
	ws.Route(ws.POST("/meta/{type-name}").To(s.addMetaField).
		Doc(TR("add meta field")).
		Param(ws.PathParameter("type-name", "type name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(MetaFieldReq{}))

	// 删除元元数据信息字段
	ws.Route(ws.DELETE("/meta/{type-name}/{group}/{field}").To(s.delMetaField).
		Doc(TR("delete meta field")).
		Param(ws.PathParameter("type-name", "type name").DataType("string")).
		Param(ws.PathParameter("group", "group name").DataType("string")).
		Param(ws.PathParameter("field", "field name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags))

	return ws
}
