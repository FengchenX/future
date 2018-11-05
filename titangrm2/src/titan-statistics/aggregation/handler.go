package aggregation

import (
	"fmt"
	"github.com/emicklei/go-restful"
	//	. "grm-searcher/dbcentral/pg"
	. "titan-statistics/types"
	//	"grm-service/dbcentral/pg"
	//	"grm-service/log"
	"grm-service/util"
)

var (
	volUnit = map[string]string{"K": "KB", "M": "MB", "G": "GB", "T": "TB"}
)

func (svc *AggrSvc) GetAggr(req *restful.Request, res *restful.Response) {
	typename := req.PathParameter("type")
	if len(typename) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid data type: %s", typename)))
		return
	}
	info, err := svc.DynamicDB.GetTypeAggr(typename)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	util.ResWriteHeaderEntity(res, info)
}

func (svc *AggrSvc) SetAggr(req *restful.Request, res *restful.Response) {
	typename := req.PathParameter("type")
	if len(typename) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid data type: %s", typename)))
		return
	}

	var data TypeAggr
	err := req.ReadEntity(&data)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	if len(data.Aggr) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid input")))
		return
	}
	data.DataType = typename

	err = svc.DynamicDB.SetTypeAggr(data)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
}

func (svc *AggrSvc) StatByAggr(req *restful.Request, res *restful.Response) {
	typename := req.PathParameter("type")
	if len(typename) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid data type: %s", typename)))
		return
	}

	field := req.PathParameter("field")
	if len(field) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid field: %s", field)))
		return
	}

	//根据type和filed来对数据库中的数据进行distinct

}
