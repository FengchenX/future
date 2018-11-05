package collector

import (
	"fmt"
	"strconv"

	"grm-service/log"

	"github.com/emicklei/go-restful"

	. "grm-service/util"

	"applications/data-collection/types"
	data "data-manager/types"
)

var (
	defaultFileds = []types.Field{
		types.Field{"gid", "number"},
		types.Field{"名称", "text"},
		types.Field{"geom", "geometry"},
	}
)

func (s DataCollectionSvc) addCollection(req *restful.Request, res *restful.Response) {
	var args addCollection
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	if len(args.Name) == 0 || len(args.Type) == 0 ||
		len(args.StartTime) == 0 || len(args.EndTime) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid collection info")))
		return
	}

	for i, field := range args.Fields {
		if field.Name == "名称" {
			args.Fields = append(args.Fields[:i], args.Fields[i+1:]...)
			break
		}
	}

	coll := types.Collection{
		Id:          NewUUID(),
		Name:        args.Name,
		StartTime:   args.StartTime,
		EndTime:     args.EndTime,
		Type:        args.Type,
		Description: args.Description,
		Fields:      args.Fields,
		User:        req.QueryParameter("id"),
	}
	coll.Fields = append(coll.Fields, defaultFileds...)

	// 创建数据记录，创建data表
	if err := s.DataDB.CreateDataTable(&coll); err != nil {
		ResWriteError(res, err)
		return
	}

	// 创建meta_object记录
	device := strconv.Itoa(s.Device.Id)
	shpType := "POINT"
	if coll.Type == types.PolygonType {
		shpType = "POLYGON"
	}
	// 添加调查任务
	_, err := s.SysDB.AddCollection(&coll)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	if err := s.MetaDB.AddDataObject(coll.Id, coll.Name, s.DataSet, device, coll.User, shpType); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, &coll)
}

func (s DataCollectionSvc) getCollections(req *restful.Request, res *restful.Response) {
	ret, err := s.SysDB.GetCollections(req.QueryParameter("id"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	if len(ret) == 0 {
		ResWriteHeaderEntity(res, nil)
	} else {
		ResWriteHeaderEntity(res, &ret)
	}
}

func (s DataCollectionSvc) delCollection(req *restful.Request, res *restful.Response) {
	colId := req.PathParameter("collection-id")
	err := s.SysDB.DelCollection(req.QueryParameter("id"), colId)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// TODO: 移除数据集市数据
	if err := s.MetaDB.DelDataObject(colId); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

func (s DataCollectionSvc) updateCollection(req *restful.Request, res *restful.Response) {
	var args updateCollection
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}

	coll := types.Collection{
		Id:          req.PathParameter("collection-id"),
		User:        req.QueryParameter("id"),
		Name:        args.Name,
		Description: args.Description,
	}
	err := s.SysDB.UpdateCollection(&coll)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	// 修改meta中记录的name
	if len(args.Name) != 0 {
		meta := types.Meta{
			Name:  "name",
			Value: args.Name,
		}
		request := types.UpdateMetaRequest{
			DataId: coll.Id,
			Group:  "Basic Information",
			Metas:  []*types.Meta{&meta},
			//DisplayField: "name",
		}
		err = s.MetaDB.UpdateDataMeta(coll.Id, &request)
		if err != nil {
			ResWriteError(res, err)
			return
		}
	}
	ResWriteHeaderEntity(res, nil)
}

// 添加记录
func (s DataCollectionSvc) addColData(req *restful.Request, res *restful.Response) {
	var args addColDataReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	colId := req.PathParameter("collection-id")
	user := req.QueryParameter("id")
	err := s.DataDB.AddTableData(colId, &args.Datas)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	func() {
		for _, field := range args.Datas {
			if field.Name == "gid" {
				if field.Value == "1" {
					// 发布图层
					data := types.DataPub{DataId: colId, User: user}
					if err := s.dataPub(&data); err != nil {
						log.Error("DataPub:", err)
						return
					}
				}
				break
			}
		}
	}()
	ResWriteHeaderEntity(res, &args.Datas)
}

// 数据发布
func (s DataCollectionSvc) dataPub(data *types.DataPub) error {
	layer := fmt.Sprintf("ftr-%s", data.DataId)
	srs, wmsUrl, wmtsUrl, wms, wfs, wmts, err := s.GeoServer.AddShpLayerWithBBox(
		"TitanGRM", s.Device.DBGeoStore, layer, "", 116.4, 116.9, 40.2, 40.5, "EPSG:4326")
	if err != nil {
		fmt.Println("method", "GeoserverUtil.AddShpLayer", "error", err.Error())
		return err
	}
	data.Srs = srs
	data.PubUrl = wmsUrl
	data.WMS = wms
	data.Wmts = wmts
	data.Wfs = wfs
	data.WmsUrl = wmsUrl
	data.WmtsUrl = wmtsUrl
	data.IsCached = false

	// TODO: 强制修改样式 1.0
	typeStr, err := s.SysDB.GetCollectionType(data.DataId)
	layerId := fmt.Sprintf("ftr-%s", data.DataId)
	if err := s.GeoServer.SetLayerStyle(layerId, typeStr); err != nil {
		fmt.Println("Failed to set layer style:", err.Error())
		return err
	}

	// 注册发布数据
	data.Style = typeStr
	if _, err := s.SysDB.AddDataPub(data); err != nil {
		fmt.Println("Failed to add published data:", err.Error())
		return err
	}
	return nil
}

// 移除记录
func (s DataCollectionSvc) delColData(req *restful.Request, res *restful.Response) {
	if err := s.DataDB.DelTableData(req.PathParameter("collection-id"),
		req.PathParameter("data-id")); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

//
func (s DataCollectionSvc) updateColData(req *restful.Request, res *restful.Response) {
	var args addColDataReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	ret, err := s.DataDB.UpdateTableData(req.PathParameter("collection-id"), req.PathParameter("data-id"), &args.Datas)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, ret)
}

//
func (s DataCollectionSvc) getColData(req *restful.Request, res *restful.Response) {
	ret, err := s.DataDB.GetTableData(&data.DataSearch{
		DataId: req.PathParameter("collection-id"),
		Limit:  req.QueryParameter("limit"),
		Offset: req.QueryParameter("offset"),
		Sort:   req.QueryParameter("sort"),
		Order:  req.QueryParameter("order"),
	})
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, ret)
}

//
func (s DataCollectionSvc) searchColData(req *restful.Request, res *restful.Response) {
	ret, err := s.DataDB.GetTableData(&data.DataSearch{
		DataId: req.PathParameter("collection-id"),
		Key:    req.PathParameter("key-word"),
		Limit:  req.QueryParameter("limit"),
		Offset: req.QueryParameter("offset"),
		Sort:   req.QueryParameter("sort"),
		Order:  req.QueryParameter("order"),
	})
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, ret)
}
