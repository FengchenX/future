package history

import (
	"fmt"

	"github.com/emicklei/go-restful"
	. "grm-searcher/types"
	"grm-service/util"
	"time"
)

func (s *HistorySvc) createHistory(req *restful.Request, res *restful.Response) {
	var history History
	err := req.ReadEntity(&history)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	if len(history.Path) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid history input")))
		return
	}

	history.DateTime = time.Now().Format("2006-01-02 15:04:05")

	err = s.DynamicDB.SetHistory(history)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	//	util.ResWriteOK(res)
}

func (s HistorySvc) getUserHistorys(req *restful.Request, res *restful.Response) {
	user := req.QueryParameter("user")
	if len(user) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid user id: %s", user)))
		return
	}
	infos, err := s.DynamicDB.GetUserHistorys(user)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteHeaderEntity(res, &infos)
}

func (s HistorySvc) getUserHistory(req *restful.Request, res *restful.Response) {
	user := req.QueryParameter("user")
	if len(user) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid user id: %s", user)))
		return
	}

	id := req.PathParameter("id")
	if len(id) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid id: %s", id)))
		return
	}

	v, err := s.DynamicDB.GetUserHistory(user, id)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteHeaderEntity(res, &v)
}

func (s HistorySvc) deleteUserHistory(req *restful.Request, res *restful.Response) {
	user := req.QueryParameter("user")
	if len(user) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid user id: %s", user)))
		return
	}

	id := req.PathParameter("id")
	if len(id) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid id: %s", id)))
		return
	}

	err := s.DynamicDB.DeleteUserHistory(user, id)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}

	//	util.ResWriteOK(res)
}
