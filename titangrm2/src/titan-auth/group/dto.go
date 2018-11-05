package group

import "titan-auth/types"

// 创建组织请求
type CreateGroupReq struct {
	Name        string `json:"name"`
	Parent      string `json:"parent_id"`
	Description string `json:"description"`
}

// 创建组织响应
type CreateGroupResp struct {
	types.Group
}

// 查询所有组织请求
type QueryGroupsReq struct {
}

// 查询所有组织响应
type QueryGroupsResp struct {
	Groups []types.Group `json:"groups,omitempty"`
}

// 更新组织请求
type UpdateGroupReq struct {
	Name        string `json:"name,omptempty"`
	Description string `json:"description,omptempty"`
}

// 添加组织成员请求
type AddGroupUsersReq struct {
	Users []struct {
		UserID  string `json:"user_id"`
		IsAdmin bool   `json:"is_admin"`
	} `json:"users,omitempty"`
}

// 添加组织成员响应
type AddGroupUsersResp struct {
}

type QueryGroupUsersReq struct {
	Group string `json:"group"`
}
type QueryGroupUsersResp struct {
	Users []types.GroupUser `json:"users,omitempty"`
}
type DelGroupUserReq struct {
	Group string `json:"group"`
	User  string `json:"user"`
}
