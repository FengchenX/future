package theme

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"
	. "grm-service/util"
	"tile-manager/dbcentral/etcd"
	"tile-manager/dbcentral/pg"
)

type ThemeSvc struct {
	SysDB     *pg.SystemDB
	DynamicDB *etcd.DynamicDB
	AuthDB    *pg.AuthDB

	BaseDir string
}

func (s ThemeSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/themes").
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("theme manager")}

	// 新建主题
	ws.Route(ws.POST("/").To(s.AddTheme).Doc(TR("add theme")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Reads(addThemeReq{}).Writes(addThemeResp{}))

	// 查询特定主题
	ws.Route(ws.GET("/{id}").To(s.GetTheme).Doc(TR("get theme")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("id", "theme id").DataType("string").Required(true)))

	// 更新主题
	ws.Route(ws.PUT("/{id}").To(s.UpdateTheme).Doc(TR("update theme")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("id", "theme id").DataType("string").Required(true)).
		Reads(updateThemeReq{}))

	// 删除主题
	ws.Route(ws.DELETE("/{id}").To(s.DelTheme).Doc(TR("delete theme")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("id", "theme id").DataType("string").Required(true)))

	// 编辑主题图标
	ws.Route(ws.POST("/{id}/pic").To(s.EditThemePic).Doc(TR("edit theme picture")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("id", "theme id").DataType("string").Required(true)).
		Reads(updateThemePicReq{}))

	// 查询主题列表
	ws.Route(ws.GET("/").To(s.GetThemes).Doc("get themes").
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.QueryParameter("offset", "Offset").DataType("string")).
		Param(ws.QueryParameter("limit", "limit").DataType("string")).
		Param(ws.QueryParameter("sort", "sort").DataType("string")).
		Param(ws.QueryParameter("order", "order").DataType("string")))

	return ws
}
