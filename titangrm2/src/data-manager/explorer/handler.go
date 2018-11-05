package explorer

import (
	"data-manager/types"
	"encoding/json"
	"grm-service/log"
	"strings"

	"fmt"
	"github.com/emicklei/go-restful"
	//. "data-manager/types"
	. "grm-service/util"
)

// GET http://localhost:8080/explorer?mart=true
func (s ExplorerSvc) getExplorer(req *restful.Request, res *restful.Response) {
	tag := "market"
	mart := strings.ToLower(req.QueryParameter("mart"))
	if mart != "true" {
		userIdD, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
		if err != nil {
			ResWriteError(res, err)
			return
		}
		tag = userIdD
	}
	if len(tag) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Failed to get user explorer, maybe user session is timeout")))
		return
	}

	path, err := s.DynamicDB.GetExplorer(tag)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	if len(path) == 0 {
		ResWriteEntity(res, nil)
		return
	}
	fmt.Println(path)
	var dat []map[string]interface{}
	if err := json.Unmarshal([]byte(path), &dat); err != nil {
		log.Error("Failed to parser explorer path:", err)
		ResWriteEntity(res, nil)
		return
	}
	ResWriteEntity(res, &dat)
}

// PUT http://localhost:8080/explorer?mart=true
func (s ExplorerSvc) updateExplorer(req *restful.Request, res *restful.Response) {
	tag := "market"
	mart := strings.ToLower(req.QueryParameter("mart"))
	if mart != "true" {
		userIdD, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
		if err != nil {
			ResWriteError(res, err)
			return
		}
		tag = userIdD
	}
	if len(tag) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Failed to get user explorer, maybe user session is timeout")))
		return
	}

	exp := explorer{}
	err := req.ReadEntity(&exp)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	class := "user"
	if tag == "market" {
		class = "market"
	}

	// 新建数据集
	var datasets types.DataSetList
	for _, dataSet := range exp.DataSets {
		dataset := types.DataSet{
			ID:          dataSet.Id,
			User:        dataSet.UserId,
			Name:        dataSet.Name,
			Type:        dataSet.Type,
			Class:       class,
			Description: dataSet.Description,
		}
		id, err := s.SysDB.AddDataSet(dataset)
		if err != nil {
			ResWriteError(res, err)
			return
		}
		dataset.ID = id
		datasets = append(datasets, dataset)
	}

	// 更新explorer
	if err := s.DynamicDB.SetExplorer(tag, exp.Path); err != nil {
		ResWriteError(res, err)
		return
	}
	if len(datasets) > 0 {
		ResWriteEntity(res, &datasets)
	} else {
		ResWriteEntity(res, nil)
	}
}
