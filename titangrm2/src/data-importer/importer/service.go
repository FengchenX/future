package importer

import (
	"data-importer/types"

	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"data-importer/dbcentral/es"
	"data-importer/dbcentral/etcd"
	"data-importer/dbcentral/pg"
	"data-importer/office"

	"grm-service/geoserver"
	"grm-service/mq"
	. "grm-service/util"
)

type ImporterSvc struct {
	SysDB        *pg.SystemDB
	MetaDB       *pg.MetaDB
	DynamicDB    *etcd.DynamicDB
	MsgQueue     *mq.RabbitMQ
	GeoServer    *geoserver.GeoserverUtil
	OfficeServer *office.OnlyOffice
	EsServer     *es.MetaEs

	DataDir   string
	ConfigDir string
	EsUrl     string
}

// WebService creates a new service that can handle REST requests for resources.
func (s ImporterSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/")
	//ws.Consumes(restful.MIME_JSON, restful.MIME_XML)
	ws.Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("data importer")}

	// 获取系统支持扫描目录
	ws.Route(ws.POST("/sysdomain").To(s.getSysDomain).
		Doc(TR("get file system domain")).
		Param(ws.FormParameter("path", TR("file sysdomain path"))).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(sysDomains{}))

	// 数据扫描
	ws.Route(ws.POST("/scan").To(s.dataScan).
		Doc(TR("data scan")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", TR("user session")).DataType("string")).
		Reads(dataScanRequest{}).
		Writes(dataScanReply{}))

	// 获取扫描结果信息
	ws.Route(ws.GET("/scan/{task-id}").To(s.dataScanResult).
		Doc(TR("data scan result")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", TR("user session")).DataType("string")).
		Param(ws.PathParameter("task-id", TR("scan task id")).DataType("string")).
		Param(ws.QueryParameter("limit", "limit").DataType("string").Required(false)).
		Param(ws.QueryParameter("offset", "offset").DataType("string").Required(false)).
		Param(ws.QueryParameter("sort", "sort").DataType("string").Required(false)).
		Param(ws.QueryParameter("order", "order").DataType("string").Required(false)).
		Param(ws.QueryParameter("file-name", TR("file name")).DataType("string").Required(false)).
		Param(ws.QueryParameter("file-size-min", TR("file size min")).DataType("string").Required(false)).
		Param(ws.QueryParameter("file-size-max", TR("file size max")).DataType("string").Required(false)).
		Param(ws.QueryParameter("create-time-min", TR("create time min")).DataType("string").Required(false)).
		Param(ws.QueryParameter("create-time-max", TR("create time max")).DataType("string").Required(false)).
		Writes(types.ScanResults{}))

	// 数据入库
	ws.Route(ws.POST("/load").To(s.dataLoad).
		Doc(TR("data load")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string")).
		Reads(dataLoadRequest{}).
		Writes(dataLoadReply{}))

	// 数据上传
	ws.Route(ws.POST("/upload").To(s.dataUpload).
		Doc(TR("data upload")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.FormParameter("file", "file").DataType("form-data")).
		Param(ws.FormParameter("data-type", "data type").DataType("form-data").Required(true)).
		//Param(ws.FormParameter("pre-type", "pre data type").DataType("form-data").Required(true)).
		Param(ws.FormParameter("data-set", "data set").DataType("form-data").Required(true)).
		Param(ws.FormParameter("geo-type", "geo type").DataType("form-data")))

	return ws
}
