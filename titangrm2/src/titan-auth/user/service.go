package user

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	. "grm-service/util"
	"titan-auth/dbcentral/etcd"
	"titan-auth/dbcentral/pg"
	. "titan-auth/types"
)

type UserSvc struct {
	AuthDB    *pg.AuthDB
	DynamicDB *etcd.DynamicDB
}

// WebService creates a new service that can handle REST requests for resources.
func (s UserSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("/").Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("auth management")}

	// 获取验证码
	ws.Route(ws.POST("/captcha").To(s.getCaptcha).Doc(TR("get captcha")).
		Metadata(api.KeyOpenAPITags, tags).Writes(captchaPic{}).Reads(captchaRequest{}))

	// 用户登录
	ws.Route(ws.PUT("/login").To(s.login).Doc(TR("user login")).
		Metadata(api.KeyOpenAPITags, tags).Writes(User{}).Reads(userlogin{}))

	// 用户登出
	ws.Route(ws.DELETE("/logout").To(s.logout).Doc(TR("user logout")).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Metadata(api.KeyOpenAPITags, tags))

	// 用户注册
	ws.Route(ws.POST("/users").To(s.userRegistry).Doc(TR("user registry")).
		Metadata(api.KeyOpenAPITags, tags).Writes(User{}).Reads(userRegistry{}))

	// 用户激活
	ws.Route(ws.PUT("/users/{user-id}/status").To(s.userActive).Doc(TR("user activation")).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("user-id", "user id").DataType("string").Required(true)).
		Metadata(api.KeyOpenAPITags, tags))

	// 获取用户列表
	ws.Route(ws.GET("/users").To(s.userList).Doc(TR("get user list")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.QueryParameter("limit", "limit").DataType("string").Required(false)).
		Param(ws.QueryParameter("offset", "offset").DataType("string").Required(false)).
		Param(ws.QueryParameter("sort", "sort").DataType("string").Required(false)).
		Param(ws.QueryParameter("order", "order").DataType("string").Required(false)).
		Writes(UserList{}))

	// TODO:修改用户信息

	// TODO:修改用户头像

	return ws
}
