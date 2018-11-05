package statistics

import (
	"fmt"
	"github.com/emicklei/go-restful"
	//	. "grm-searcher/dbcentral/pg"
	//	. "grm-searcher/types"
	//	"grm-service/dbcentral/pg"
	//	"grm-service/log"
	"grm-service/util"
)

var (
	volUnit = map[string]string{"K": "KB", "M": "MB", "G": "GB", "T": "TB"}

	META_INDEX = "datameta"
)

func (svc *StatSvc) dataTypeStat(req *restful.Request, res *restful.Response) {
	info, err := svc.MetaDB.DataTypeStat()
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteHeaderEntity(res, info)
}

func (svc *StatSvc) userDataTypeStat(req *restful.Request, res *restful.Response) {
	userId := req.PathParameter("user_id")
	if len(userId) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid user id: %s", userId)))
		return
	}

	ids, err := svc.SysDB.GetDataIds(userId)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	if len(ids) > 0 {
		info, err := svc.MetaDB.UserDataTypeStat(ids)
		if err != nil {
			util.ResWriteError(res, err)
			return
		}
		util.ResWriteHeaderEntity(res, info)
	}
}

func (svc *StatSvc) subTypeStat(req *restful.Request, res *restful.Response) {
	typeName := req.PathParameter("type")
	if len(typeName) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid type name: %s", typeName)))
		return
	}

	info, err := svc.MetaDB.SubTypeStat(typeName)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteHeaderEntity(res, info)
}

func (svc *StatSvc) userSubTypeStat(req *restful.Request, res *restful.Response) {
	typeName := req.PathParameter("type")
	if len(typeName) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid type name: %s", typeName)))
		return
	}

	userId := req.PathParameter("user_id")
	if len(userId) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid user id: %s", userId)))
		return
	}

	ids, err := svc.SysDB.GetDataIds(userId)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	if len(ids) > 0 {
		info, err := svc.MetaDB.UserSubTypeStat(typeName, ids)
		if err != nil {
			util.ResWriteError(res, err)
			return
		}
		util.ResWriteHeaderEntity(res, info)
	}
}

func (svc *StatSvc) apiStat(req *restful.Request, res *restful.Response) {
	//	ids, err := svc.SysDB.GetMarketsDataIds()
	//	if err != nil {
	//		util.ResWriteError(res, err)
	//	}
	//
	//	infos, total, err := svc.MetaDB.SearchByGeo(ids, r)
	//	if err != nil {
	//		util.ResWriteError(res, err)
	//		return
	//	}
	//	util.ResWriteHeaderEntity(res, MetaInfosTotalReply{total, infos})
}
