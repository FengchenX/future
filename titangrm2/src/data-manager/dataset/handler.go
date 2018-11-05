package dataset

import (
	"fmt"
	"strings"

	"grm-service/log"

	"github.com/emicklei/go-restful"

	. "data-manager/types"
	. "grm-service/util"
)

// POST http://localhost:8080/datasets
func (s *DataSetSvc) createDataset(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	var args addDatasetRequest
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	if len(args.Name) == 0 || len(userId) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid dataset name or user id")))
		return
	}

	class := "user"
	if args.IsMarket {
		class = "market"
	}
	dataset := DataSet{
		User:        userId,
		Name:        args.Name,
		Type:        args.Type,
		Class:       class,
		Description: args.Description,
	}
	id, err := s.SysDB.AddDataSet(dataset)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	dataset.ID = id
	ResWriteEntity(res, &dataset)
}

// GET http://localhost:8080/datasets
func (s DataSetSvc) getDatasets(req *restful.Request, res *restful.Response) {
	var userId string
	var err error
	mart := strings.ToLower(req.QueryParameter("mart"))
	if mart != "true" {
		userId, err = s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
		if err != nil {
			ResWriteError(res, err)
			return
		}
	}

	sets, err := s.SysDB.GetDataSets(userId)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, &sets)
}

// DELETE http://localhost:8080/datasets/{dataset-id}
func (s DataSetSvc) delDataset(req *restful.Request, res *restful.Response) {
	err := s.SysDB.DelDataSet(req.PathParameter("dataset-id"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

// PUT http://localhost:8080/datasets/{dataset-id}
func (s DataSetSvc) updateDataset(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("dataset-id")
	var args updateDatasetRequest
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}

	err := s.SysDB.UpdateDataSet(id, args.Name, args.Description)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

// DELETE http://localhost:8080/datasets/{dataset-id}/datas
func (s DataSetSvc) truncateDataset(req *restful.Request, res *restful.Response) {
	err := s.SysDB.TruncateDataSet(req.PathParameter("dataset-id"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}

// DELETE http://localhost:8080/datasets/{dataset-id}/data/{data-id}
func (s DataSetSvc) delDatasetData(req *restful.Request, res *restful.Response) {
	dataset := req.PathParameter("dataset-id")
	var datas addDatas
	if err := req.ReadEntity(&datas); err != nil {
		ResWriteError(res, err)
		return
	}
	for _, data := range datas.Datas {
		// TODO: 判断是否是数据拥有者，修改数据状态变为已删除
		err := s.SysDB.DelDataSetData(dataset, data)
		if err != nil {
			ResWriteError(res, err)
			return
		}
	}
	ResWriteHeaderEntity(res, nil)
}

// PUT http://localhost:8080/datasets/{dataset-id}/data/{data-id}
func (s DataSetSvc) addDatasetData(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	dataset := req.PathParameter("dataset-id")
	var datas addDatas
	if err := req.ReadEntity(&datas); err != nil {
		ResWriteError(res, err)
		return
	}

	for _, data := range datas.Datas {
		err = s.SysDB.AddDataSetData(dataset, data, userId, DataSubscribed)
		if err != nil {
			log.Error(err)
			ResWriteError(res, fmt.Errorf(TR("Failed to add data, maybe already exists.")))
			return
		}
	}
	ResWriteHeaderEntity(res, nil)
}

// PUT http://localhost:8080/datasets/{dataset-id}/datas/{src-dataset}
func (s DataSetSvc) addDatasetDatas(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	dataset := req.PathParameter("dataset-id")
	srcDataSet := req.PathParameter("src-dataset")
	err = s.SysDB.AddDataSetDatas(dataset, userId, DataSubscribed, srcDataSet)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteHeaderEntity(res, nil)
}
