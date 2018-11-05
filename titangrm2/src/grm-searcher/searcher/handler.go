package searcher

import (
	"fmt"
	"github.com/emicklei/go-restful"
	. "grm-searcher/dbcentral/pg"
	. "grm-searcher/types"
	"grm-service/dbcentral/pg"
	"grm-service/log"
	"grm-service/util"
)

var (
	volUnit = map[string]string{"K": "KB", "M": "MB", "G": "GB", "T": "TB"}

	META_INDEX = "datameta"
)

//用于数据分发时，对传入的数据，按照范围进行过滤
func (svc *SearcherSvc) dataFilter(req *restful.Request, res *restful.Response) {
	var dfReq DataFilterRequest
	err := req.ReadEntity(&dfReq)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	if len(dfReq.UserId) == 0 || len(dfReq.GeoRegion) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid Request")))
		return
	}

	infos := make([]*MetaInfo, 0)
	var data_ids []string = make([]string, 0)
	data_ids = append(data_ids, dfReq.Datas...)

	if len(dfReq.DataSets) > 0 {
		_ds := make([]string, 0)
		for _, ds := range dfReq.DataSets {
			_ds = append(_ds, fmt.Sprintf("'%s'", ds))
		}

		data_ids, err = svc.SysDB.GetDataIds(dfReq.UserId, _ds, data_ids)
	}

	_dataids := make([]string, 0)
	for _, id := range data_ids {
		_dataids = append(_dataids, fmt.Sprintf("'%s'", id))
	}

	infos, total, err := svc.MetaDB.DataFilter(_dataids, dfReq.GeoRegion)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteHeaderEntity(res, &MetaInfosTotalReply{total, infos})
}

//通过传入的范围，来对数据本身进行查询，返回数据对象在这个查询范围内的数据信息
func (svc *SearcherSvc) dataSearch(req *restful.Request, res *restful.Response) {
	var r SearchInfo
	err := req.ReadEntity(&r)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	if len(r.DataId) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid Request")))
		return
	}

	var dataDB *DataDB
	if storage, ok := svc.DataIdConns[r.DataId]; ok {
		//存在
		var err error
		_dataDb, err := pg.ConnectDataDBUrl(storage)
		if err != nil {
			util.ResWriteError(res, err)
		}
		defer _dataDb.DisConnect()
		dataDB = &DataDB{_dataDb}
	} else {
		//不存在，则直接获取
		storage, err := svc.MetaDB.GetPgStorage(r.DataId)
		if err != nil {
			util.ResWriteError(res, err)
			return
		}

		fmt.Println("storage:", storage)

		svc.DataIdConns[r.DataId] = storage
		_dataDb, err := pg.ConnectDataDBUrl(storage)
		if err != nil {
			util.ResWriteError(res, err)
		}
		defer _dataDb.DisConnect()
		dataDB = &DataDB{_dataDb}
	}

	data, err := dataDB.GetTableData(r)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteHeaderEntity(res, &data)

}

func (svc *SearcherSvc) geoSearch(req *restful.Request, res *restful.Response) {
	var r SearchInfo
	err := req.ReadEntity(&r)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	//	if len(r.) == 0 || len(r.GeoRegion) == 0 {
	//		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid Request")))
	//		return
	//	}

	ids, err := svc.SysDB.GetMarketsDataIds()
	if err != nil {
		util.ResWriteError(res, err)
	}

	infos, total, err := svc.MetaDB.SearchByGeo(ids, r)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteHeaderEntity(res, &MetaInfosTotalReply{total, infos})
}

func (svc *SearcherSvc) typeSearch(req *restful.Request, res *restful.Response) {
	type_name := req.PathParameter("type_name")
	if len(type_name) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid type_name: %s", type_name)))
		return
	}

	//	fmt.Println("type_name:", type_name)
	//通过type_name查询数据

	var si SearchInfo
	err := req.ReadEntity(&si)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	data_ids := make([]string, 0)
	if len(si.DatasetIds) > 0 {
		ds_ids := make([]string, 0)
		for _, v := range si.DatasetIds {
			ds_ids = append(ds_ids, fmt.Sprintf("'%s'", v))
		}

		data_ids, err = svc.SysDB.GetDataIds(si.Userid, ds_ids, data_ids)
		if err != nil {
			util.ResWriteError(res, err)
			return
		}

		if len(data_ids) == 0 {
			util.ResWriteEntity(res, &MetaInfosTotalReply{0, nil})
			return
		}
	}

	//各种条件组合查询
	total, infos, err := svc.EsUtil.TypeQuery(META_INDEX, type_name, data_ids, si)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	//userid 2 username
	for _, info := range infos {
		info.LoadUser, err = svc.DynamicDB.GetUserName(info.LoadUser)
		if err != nil {
			log.Warn("GetUserName error:", err)
		}
	}

	util.ResWriteEntity(res, &MetaInfosTotalReply{total, infos})
}

func (svc *SearcherSvc) keySearch(req *restful.Request, res *restful.Response) {
	key := req.PathParameter("key")
	if len(key) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid key")))
		return
	}

	fmt.Println("key:", key)
	//通过key查询数据，es的操作
	order := req.QueryParameter("order")
	limit := req.QueryParameter("limit")
	sort := req.QueryParameter("sort")
	offset := req.QueryParameter("offset")

	total, infos, err := svc.EsUtil.QueryByKey(META_INDEX, key, order, limit, sort, offset)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	util.ResWriteEntity(res, &MetaInfosTotalReply{total, infos})
}

func (svc *SearcherSvc) metaIdSearch(req *restful.Request, res *restful.Response) {
	data_id := req.PathParameter("data_id")
	if len(data_id) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid data_id")))
		return
	}

	fmt.Println("data_id:", data_id)
	//通过key查询数据，es的操作
	info, err := svc.EsUtil.QueryById(META_INDEX, data_id)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteEntity(res, &info)
}

func (svc *SearcherSvc) marketIdSearch(req *restful.Request, res *restful.Response) {
	marketid := req.PathParameter("id")
	if len(marketid) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid market Id")))
		return
	}

	fmt.Println("marketid:", marketid)
	//通过key查询数据，es的操作
	order := req.QueryParameter("order")
	limit := req.QueryParameter("limit")
	sort := req.QueryParameter("sort")
	offset := req.QueryParameter("offset")
	key := req.QueryParameter("key")

	//dataids
	dataids, err := svc.SysDB.GetDataIdsByOne(marketid)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	if len(dataids) == 0 {
		util.ResWriteEntity(res, &MetaInfosTotalReply{0, nil})
		return
	}

	total, infos, err := svc.EsUtil.QueryByDataset(dataids, "", key, META_INDEX, order, limit, sort, offset)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	//userid 2 username
	for _, info := range infos {
		info.LoadUser, err = svc.DynamicDB.GetUserName(info.LoadUser)
		if err != nil {
			log.Warn("GetUserName error:", err)
		}
	}

	util.ResWriteEntity(res, &MetaInfosTotalReply{total, infos})
}

func (svc *SearcherSvc) datasetIdSearch(req *restful.Request, res *restful.Response) {
	dsid := req.PathParameter("id")
	if len(dsid) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid dataset Id")))
		return
	}

	fmt.Println("dataset id:", dsid)
	//通过key查询数据，es的操作
	order := req.QueryParameter("order")
	limit := req.QueryParameter("limit")
	sort := req.QueryParameter("sort")
	offset := req.QueryParameter("offset")
	key := req.QueryParameter("key")
	typeName := req.QueryParameter("type")

	//dataids
	dataids, err := svc.SysDB.GetDataIdsByOne(dsid)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	if len(dataids) == 0 {
		log.Warn("this dataset has no datas:", dsid)
		util.ResWriteEntity(res, &MetaInfosTotalReply{0, nil})
		return
	}

	fmt.Println("dataids:", dataids)
	total, infos, err := svc.EsUtil.QueryByDataset(dataids, typeName, key, META_INDEX, order, limit, sort, offset)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	//userid 2 username
	for _, info := range infos {
		info.LoadUser, err = svc.DynamicDB.GetUserName(info.LoadUser)
		if err != nil {
			log.Warn("GetUserName error:", err)
		}
	}

	util.ResWriteEntity(res, &MetaInfosTotalReply{total, infos})
}
