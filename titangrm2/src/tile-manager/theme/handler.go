package theme

import (
	"fmt"
	"github.com/emicklei/go-restful"
	"github.com/sirupsen/logrus"
	. "grm-service/util"
	"os"
	"path/filepath"
	"strconv"
	"tile-manager/types"
)

// 添加主题
func (s ThemeSvc) AddTheme(req *restful.Request, res *restful.Response) {
	var request addThemeReq
	if err := req.ReadEntity(&request); err != nil {
		ResWriteError(res, err)
		return
	}
	session := req.HeaderParameter("auth-session")
	userId, err := s.DynamicDB.GetUserId(session)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	theme := types.Theme{
		Id:             0,
		Name:           request.Name,
		ImageType:      request.ImageType,
		TileSize:       request.TileSize,
		TileFormat:     request.TileFormat,
		Srs:            request.Srs,
		NoData:         request.NoData,
		TimeResolution: request.TimeResolution,
		AuthUser:       userId,
		CreateTime:     "",
		Description:    request.Description,
		Transparency:   request.Transparency,
		Thumb:          "",
		DeviceId:       request.Device,
		DeviceInfo:     types.Device{},
		Projection:     request.Projection,
		Center:         "",
		Session:        request.Session,
	}
	id, err := s.SysDB.AddTheme(&theme)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	theme.Id = id
	resp := addThemeResp{
		Id: id,
	}
	ResWriteEntity(res, &resp)
}

// 更新主题
func (s ThemeSvc) UpdateTheme(req *restful.Request, res *restful.Response) {
	themeReq := updateThemeReq{}
	if err := req.ReadEntity(&themeReq); err != nil {
		ResWriteError(res, err)
		return
	}
	id, err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	theme := types.Theme{
		Id:             id,
		Name:           themeReq.Name,
		ImageType:      themeReq.ImageType,
		TileSize:       themeReq.TileSize,
		TileFormat:     themeReq.TileFormat,
		Srs:            themeReq.Srs,
		NoData:         themeReq.NoData,
		TimeResolution: themeReq.TimeResolution,
		//AuthUser:       themeReq.AuthUser,
		CreateTime:   "",
		Description:  themeReq.Description,
		Transparency: themeReq.Transparency,
		Thumb:        "",
		DeviceId:     0,
		DeviceInfo:   types.Device{},
		Projection:   themeReq.Projection,
		Center:       "",
		Session:      "",
	}

	if err := s.SysDB.UpdateTheme(&theme); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, nil)
}

// 移除主题
func (s ThemeSvc) DelTheme(req *restful.Request, res *restful.Response) {
	id, err := strconv.Atoi(req.PathParameter("id"))
	if err != nil {
		ResWriteError(res, err)
		return
	}
	session := req.HeaderParameter("auth-session")
	userId, e := s.DynamicDB.GetUserId(session)
	if e != nil {
		ResWriteError(res, e)
		return
	}
	if err := s.SysDB.DelTheme(id, userId); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, nil)
}

// 获取特定主题
func (s ThemeSvc) GetTheme(req *restful.Request, res *restful.Response) {
	session := req.HeaderParameter("auth-session")
	userId, e := s.DynamicDB.GetUserId(session)
	if e != nil {
		ResWriteError(res, e)
		return
	}
	themeId, e := strconv.Atoi(req.PathParameter("id"))
	if e != nil {
		ResWriteError(res, e)
		return
	}
	userInfos, e := s.AuthDB.GetUserInfo("admin")
	if e != nil {
		ResWriteError(res, e)
		return
	}
	users := []string{userId}
	for _, info := range userInfos {
		users = append(users, info.Id)
	}
	_, themes, err := s.SysDB.GetThemes(users, themeId, -1, -1, "", "")
	if err != nil {
		ResWriteError(res, err)
		return
	}
	if len(themes) <= 0 {
		ResWriteEntity(res, nil)
		return
	}
	//todo Device等信息
	ResWriteEntity(res, &themes[0])
}

// 编辑主题图片
func (s ThemeSvc) EditThemePic(req *restful.Request, res *restful.Response) {
	id := req.PathParameter("id")
	name := fmt.Sprintf("theme_%s_%s.jpg", id, GetTimeNow())
	session := req.HeaderParameter("auth-session")
	userId, e := s.DynamicDB.GetUserId(session)
	if e != nil {
		ResWriteError(res, e)
		return
	}
	var request updateThemePicReq
	if err := req.ReadEntity(&request); err != nil {
		ResWriteError(res, err)
		return
	}
	baseFile := filepath.Join("theme", userId, name)
	themeFile := filepath.Join(s.BaseDir, baseFile)
	CheckFileDir(themeFile)

	f, err := os.OpenFile(themeFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		logrus.Errorln("OpenFile error:", err)
		ResWriteError(res, err)
		return
	}
	defer f.Close()
	if _, err := f.Write([]byte(request.Picture)); err != nil {
		ResWriteError(res, err)
		return
	}
	i, e := strconv.Atoi(id)
	if e != nil {
		ResWriteError(res, e)
		return
	}
	if err := s.SysDB.UpdateThemePic(i, baseFile); err != nil {
		ResWriteError(res, err)
		return
	}

	ResWriteEntity(res, nil)
}

// 获取所有主题
func (s ThemeSvc) GetThemes(req *restful.Request, res *restful.Response) {
	session := req.HeaderParameter("auth-session")
	userId, e := s.DynamicDB.GetUserId(session)
	if e != nil {
		ResWriteError(res, e)
		return
	}
	userInfos, e := s.AuthDB.GetUserInfo("admin")
	if e != nil {
		ResWriteError(res, e)
		return
	}

	var limit, offset int
	if IsNum(req.QueryParameter("limit")) {
		limit, e = strconv.Atoi(req.QueryParameter("limit"))
		if e != nil {
			ResWriteError(res, e)
			return
		}
	} else {
		limit = 0
	}

	if IsNum(req.QueryParameter("offset")) {
		offset, e = strconv.Atoi(req.QueryParameter("offset"))
		if e != nil {
			ResWriteError(res, e)
			return
		}
	} else {
		offset = 0
	}

	sort := req.QueryParameter("sort")
	order := req.QueryParameter("order")

	users := []string{userId}
	for _, info := range userInfos {
		users = append(users, info.Id)
	}

	total, themes, err := s.SysDB.GetThemes(users, -1,
		limit, offset, sort, order)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	var resp getThemesResp
	resp.Total = total
	resp.Themes = themes

	//todo Device等信息

	ResWriteEntity(res, &resp)
}
