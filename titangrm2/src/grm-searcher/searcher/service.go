package searcher

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"grm-searcher/dbcentral/es"
	"grm-searcher/dbcentral/etcd"
	"grm-searcher/dbcentral/pg"
	. "grm-searcher/types"
	//	. "grm-service/dbcentral/pg"
	. "grm-service/util"
)

type SearcherSvc struct {
	SysDB     *pg.SystemDB
	MetaDB    *pg.MetaDB
	DynamicDB *etcd.DynamicDB
	EsUtil    *es.ESUtil

	DataDir   string
	ConfigDir string

	DataIdConns map[string]string
}

// WebService creates a new service that can handle REST requests for resources.
func (s SearcherSvc) WebService() *restful.WebService {
	s.DataIdConns = make(map[string]string, 0)

	ws := new(restful.WebService)
	ws.Path("/").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("searcher")}

	//对传入的数据进行过滤
	ws.Route(ws.POST("/data/_filter").To(s.dataFilter).
		Doc(TR("data fileter")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(DataFilterRequest{}).
		Writes(MetaInfosTotalReply{}))

	//查询数据的内容
	ws.Route(ws.POST("/data/_search").To(s.dataSearch).
		Doc(TR("data search")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(SearchInfo{}).
		Writes(TableData{}))

	//查询元数据信息   只按照geom字段来查询
	//	ws.Route(ws.POST("/meta/geometry").To(s.geoSearch).
	//		Doc(TR("meta geometry search")).
	//		Metadata(api.KeyOpenAPITags, tags).
	//		Writes(MetaInfosTotalReply{}))

	ws.Route(ws.GET("/meta/id/{data_id}").To(s.metaIdSearch).
		Doc(TR("meta id search")).
		Param(ws.PathParameter("data_id", "data id").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(TypeMeta{}))

	//这里要用到es，会存在基于关键字的查询，同时会有空间查询
	ws.Route(ws.GET("/meta/key/{key}").To(s.keySearch).
		Doc(TR("meta key search")).
		Param(ws.PathParameter("key", "key words").DataType("string")).
		Param(ws.QueryParameter("order", "order").DataType("string")).
		Param(ws.QueryParameter("limit", "limit").DataType("string")).
		Param(ws.QueryParameter("sort", "sort").DataType("string")).
		Param(ws.QueryParameter("offset", "offset").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(MetaInfosTotalReply{}))

	ws.Route(ws.GET("/dataset/{id}/_search").To(s.datasetIdSearch).
		Doc(TR("dataset id search")).
		Param(ws.PathParameter("id", "dataset id").DataType("string")).
		Param(ws.QueryParameter("key", "key words").DataType("string")).
		Param(ws.QueryParameter("type", "type name").DataType("string")).
		Param(ws.QueryParameter("order", "order").DataType("string")).
		Param(ws.QueryParameter("limit", "limit").DataType("string")).
		Param(ws.QueryParameter("sort", "sort").DataType("string")).
		Param(ws.QueryParameter("offset", "offset").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(MetaInfosTotalReply{}))

	ws.Route(ws.GET("/marketplace/{id}/_search").To(s.marketIdSearch).
		Doc(TR("marketplace dataset id search")).
		Param(ws.PathParameter("id", "dataset id").DataType("string")).
		Param(ws.QueryParameter("key", "key words").DataType("string")).
		Param(ws.QueryParameter("order", "order").DataType("string")).
		Param(ws.QueryParameter("limit", "limit").DataType("string")).
		Param(ws.QueryParameter("sort", "sort").DataType("string")).
		Param(ws.QueryParameter("offset", "offset").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(MetaInfosTotalReply{}))

	ws.Route(ws.POST("/meta/type/{type_name}").To(s.typeSearch).
		Doc(TR("meta type search")).
		Param(ws.PathParameter("type_name", "type name").DataType("string")).
		Metadata(api.KeyOpenAPITags, tags).
		Reads(SearchInfo{}).
		Writes(MetaInfosTotalReply{}))
	return ws
}
