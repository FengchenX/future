package config

import (
	"fmt"

	"github.com/emicklei/go-restful"
	. "grm-labelmgr/types"
	"grm-service/util"
	//	"time"
)

func (s *ConfigSvc) createLayerConfig(req *restful.Request, res *restful.Response) {
	user := req.PathParameter("user")
	if len(user) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid user id: %s", user)))
		return
	}

	var data UserData
	err := req.ReadEntity(&data)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	if len(data.Data) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid input")))
		return
	}
	data.UserId = user

	err = s.DynamicDB.SetUserData(data)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
}

func (s ConfigSvc) getLayerConfig(req *restful.Request, res *restful.Response) {
	user := req.PathParameter("user")
	if len(user) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid user id: %s", user)))
		return
	}

	data_type := req.PathParameter("data_type")
	if len(data_type) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid data type: %s", data_type)))
		return
	}

	data_name := req.PathParameter("data_name")
	if len(data_name) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid data name: %s", data_name)))
		return
	}

	v, err := s.DynamicDB.GetUserData(user, data_type, data_name)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteHeaderEntity(res, &v)
}

func (s ConfigSvc) deleteLayerConfig(req *restful.Request, res *restful.Response) {
	user := req.PathParameter("user")
	if len(user) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid user id: %s", user)))
		return
	}

	data_type := req.PathParameter("data_type")
	if len(data_type) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid data type: %s", data_type)))
		return
	}

	data_name := req.PathParameter("data_name")
	if len(data_name) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid data name: %s", data_name)))
		return
	}

	err := s.DynamicDB.DeleteUserData(user, data_type, data_name)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
}

func (s ConfigSvc) getUserDatas(req *restful.Request, res *restful.Response) {
	user := req.PathParameter("user")
	if len(user) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid user id: %s", user)))
		return
	}

	data_type := req.PathParameter("data_type")
	if len(data_type) == 0 {
		util.ResWriteError(res, fmt.Errorf(util.TR("Invalid data type: %s", data_type)))
		return
	}

	v, err := s.DynamicDB.GetUserDatas(user, data_type)
	if err != nil {
		util.ResWriteError(res, err)
		return
	}
	util.ResWriteHeaderEntity(res, &v)
}
