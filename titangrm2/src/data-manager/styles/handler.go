package styles

import (
	"fmt"

	"grm-service/common"
	//"grm-service/geoserver"
	. "grm-service/util"

	//"data-manager/types"

	"github.com/emicklei/go-restful"
	//"grm-service/log"
)

func (s StyleSvc) addStyle(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	var args addStyleReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	if len(args.Name) == 0 || len(args.Sld) == 0 || len(userId) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid style info")))
		return
	}

	// 添加记录
	style := common.LayerStyle{
		User:        userId,
		Name:        args.Name,
		Type:        args.Type,
		Description: args.Description,
	}
	ret, err := s.SysDB.AddStyle(&style)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// GeoServer添加样式
	sldName := userId + "_" + style.Id
	if err := s.GeoServer.AddStyle(sldName, args.Sld); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, ret)
}

func (s StyleSvc) getStyles(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ret, err := s.SysDB.GetStyles(userId, req.QueryParameter("style-type"))
	if err != nil {
		ResWriteError(res, err)
		return
	}

	for i, val := range ret {
		sldName := userId + "_" + val.Id
		switch val.Id {
		case "1":
			sldName = "point"
		case "2":
			sldName = "line"
		case "3":
			sldName = "polygon"
		}
		sld, err := s.GeoServer.GetStyle(sldName)
		if err != nil {
			ResWriteError(res, err)
			return
		}
		ret[i].Sld = sld
	}
	ResWriteEntity(res, &ret)
}

func (s StyleSvc) getStyle(req *restful.Request, res *restful.Response) {
	style := req.PathParameter("style-id")
	ret, err := s.SysDB.GetStyle(style)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 获取sld
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	sldName := userId + "_" + style
	switch style {
	case "1":
		sldName = "point"
	case "2":
		sldName = "line"
	case "3":
		sldName = "polygon"
	}
	sld, err := s.GeoServer.GetStyle(sldName)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	ret.Sld = sld
	ResWriteEntity(res, &ret)
}

func (s StyleSvc) delStyle(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	style := req.PathParameter("style-id")
	_, err = s.SysDB.DelStyle(style, userId)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	// 移除geoserver样式
	sldName := userId + "_" + style
	if err := s.GeoServer.DeleteStyle(sldName); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, nil)
}

// 更新样式
func (s StyleSvc) updateStyle(req *restful.Request, res *restful.Response) {
	userId, err := s.DynamicDB.GetUserId(req.HeaderParameter("auth-session"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	style := req.PathParameter("style-id")

	var args addStyleReq
	if err := req.ReadEntity(&args); err != nil {
		ResWriteError(res, err)
		return
	}
	if len(args.Name) == 0 || len(args.Sld) == 0 || len(userId) == 0 {
		ResWriteError(res, fmt.Errorf(TR("Invalid style info")))
		return
	}

	// GeoServer修改样式
	sldName := userId + "_" + style
	if err := s.GeoServer.EditStyle(sldName, args.Sld); err != nil {
		ResWriteError(res, err)
		return
	}

	// 修改记录
	layerStyle := common.LayerStyle{
		Id:          style,
		User:        userId,
		Name:        args.Name,
		Description: args.Description,
	}
	if err := s.SysDB.UpdateStyle(&layerStyle); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, nil)
}
