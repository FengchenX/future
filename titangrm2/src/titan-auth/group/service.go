package group

import (
	"github.com/emicklei/go-restful"
	api "github.com/emicklei/go-restful-openapi"

	. "grm-service/util"
	"titan-auth/dbcentral/etcd"
	"titan-auth/dbcentral/pg"
)

type GroupSvc struct {
	AuthDB    *pg.AuthDB
	DynamicDB *etcd.DynamicDB
}

// WebService creates a new service that can handle REST requests for resources.
func (s GroupSvc) WebService() *restful.WebService {
	ws := new(restful.WebService)
	ws.Path("groups").
		Produces(restful.MIME_JSON, restful.MIME_JSON)

	tags := []string{TR("group management")}

	// 添加组织
	ws.Route(ws.POST("/").To(s.createGroup).Doc(TR("create group")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Writes(CreateGroupResp{}).Reads(CreateGroupReq{}))

	// 查询组织
	ws.Route(ws.GET("/").To(s.queryGroups).Doc(TR("query groups")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Writes(QueryGroupsResp{}))

	// 修改组织信息
	ws.Route(ws.PUT("/{group-id}").To(s.updateGroup).Doc(TR("update group info")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("group-id", TR("group id")).DataType("string").Required(true)).
		Reads(UpdateGroupReq{}))

	// 移除组织
	ws.Route(ws.DELETE("/{group-id}").To(s.delGroup).Doc("delete group").
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.HeaderParameter("auth-session", "user session").DataType("string").Required(true)).
		Param(ws.PathParameter("group-id", TR("group id")).DataType("string").Required(true)))

	// 添加组织成员
	ws.Route(ws.POST("/{group-id}/users").To(s.addGroupUsers).Doc(TR("add group users")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.PathParameter("group-id", TR("group id")).DataType("string").Required(true)).
		Writes(AddGroupUsersResp{}).
		Reads(AddGroupUsersReq{}))

	// 查询组织成员
	ws.Route(ws.GET("/{group-id}/users").To(s.queryGroupUsers).Doc(TR("query group users")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.PathParameter("group-id", TR("group id")).DataType("string").Required(true)).
		Writes(QueryGroupUsersResp{}))

	// 移除组织成员
	ws.Route(ws.DELETE("/{group-id}/users/{user-id}").To(s.delGroupUsers).Doc(TR("delete group users")).
		Metadata(api.KeyOpenAPITags, tags).
		Param(ws.PathParameter("group-id", TR("group id")).DataType("string").Required(true)).
		Param(ws.PathParameter("user-id", TR("user id")).DataType("string").Required(true)))

	return ws
}
