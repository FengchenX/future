package dataset

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	"data-manager/dbcentral/etcd"
	"data-manager/dbcentral/pg"
	. "data-manager/types"
	. "grm-service/util"
)

type DataSetSvc struct {
	SysDB     *pg.SystemDB
	DynamicDB *etcd.DynamicDB
}

// WebService creates a new service that can handle REST requests for resources.
func (s DataSetSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/datasets").
		//Consumes(restful.MIME_JSON, restful.MIME_JSON).
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("dataset")}

	// 创建数据集
	ws.Route(ws.POST("").To(s.createDataset).
		Doc(TR("create dataset")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Reads(addDatasetRequest{}).
		Writes(DataSet{}))

	// 获取用户/数据集市所有数据集
	ws.Route(ws.GET("").To(s.getDatasets).
		Doc(TR("get dataset list")).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.QueryParameter("mart", "is datamart: true or fase").DataType("string").Required(true)).
		Metadata(api.KeyOpenAPITags, tags).
		Writes(DataSetList{}))

	// 移除数据集
	ws.Route(ws.DELETE("/{dataset-id}").To(s.delDataset).
		Doc(TR("delete dataset")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("dataset-id", "dataset id").DataType("string").Required(true)))

	// 修改数据集信息
	ws.Route(ws.PUT("/{dataset-id}").To(s.updateDataset).
		Doc(TR("update dataset")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("dataset-id", "dataset id").DataType("string").Required(true)).
		Reads(updateDatasetRequest{}))

	// 移除数据集特定数据
	ws.Route(ws.PUT("/{dataset-id}/data").To(s.delDatasetData).
		Doc(TR("remove data of dataset")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("dataset-id", "dataset id").DataType("string").Required(true)).
		Reads(addDatas{}))

	// 数据集添加数据
	ws.Route(ws.POST("/{dataset-id}/data").To(s.addDatasetData).
		Doc(TR("add data of dataset")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("dataset-id", "dataset id").DataType("string").Required(true)).
		Reads(addDatas{}))

	// 移除数据集下所有数据
	ws.Route(ws.DELETE("/{dataset-id}/datas").To(s.truncateDataset).
		Doc(TR("truncate dataset")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("dataset-id", "dataset id").DataType("string").Required(true)))

	// 拷贝指定数据集数据
	ws.Route(ws.PUT("/{dataset-id}/datas/{src-dataset}").To(s.addDatasetDatas).
		Doc(TR("copy dataset")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("dataset-id", "dataset id").DataType("string").Required(true)).
		Param(ws.PathParameter("src-dataset", "source dataset").DataType("string").Required(true)))

	return ws
}
