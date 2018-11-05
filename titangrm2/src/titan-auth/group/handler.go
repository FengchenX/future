package group

import (
	"fmt"

	"github.com/emicklei/go-restful"

	"grm-service/log"
	. "grm-service/util"

	"titan-auth/types"
)

// 创建组织
func (s GroupSvc) createGroup(req *restful.Request, res *restful.Response) {
	var request CreateGroupReq
	session := req.HeaderParameter("auth-session")
	if err := req.ReadEntity(&request); err != nil {
		ResWriteError(res, err)
		return
	}
	if !authCheck(req, session) {
		ResWriteError(res, fmt.Errorf(TR("User authority is not enough")))
		return
	}
	group := types.Group{
		Id:          NewUUID(),
		Name:        request.Name,
		ParentID:    request.Parent,
		Path:        "",
		Layer:       "",
		Level:       0,
		Description: request.Description,
		CreateTime:  "",
		Children:    nil,
	}

	var parentGroup *types.Group
	var err error
	if group.ParentID != "" && group.ParentID != "0" && group.ParentID != "-1" {
		// 获取指定父节点layer和level
		parentGroup, err = s.AuthDB.GetGroupInfo(group.ParentID)
		if err != nil {
			log.Printf("GetGroupInfo: %s\n", err.Error())
			ResWriteError(res, err)
			return
		}
	}

	if parentGroup != nil && len(parentGroup.Name) > 0 {
		group.Layer = fmt.Sprintf("%s%s-", parentGroup.Layer, parentGroup.Id)
		group.Level = parentGroup.Level + 1
	} else {
		group.ParentID = "0"
		group.Layer = ""
		group.Level = 1
	}

	group, err = s.AuthDB.CreateGroup(group)
	if err != nil {
		ResWriteError(res, err)
		return
	}

	var resp CreateGroupResp
	resp.Group = group
	ResWriteEntity(res, &resp)
}

// 查询组织
func (s GroupSvc) queryGroups(req *restful.Request, res *restful.Response) {
	var resp QueryGroupsResp
	groupList, err := s.AuthDB.QueryGroups()
	if err != nil {
		ResWriteError(res, err)
		return
	}
	resp.Groups = groupList
	ResWriteEntity(res, &resp)
}

// 更新组织
func (s GroupSvc) updateGroup(req *restful.Request, res *restful.Response) {
	var request UpdateGroupReq
	session := req.HeaderParameter("auth-session")
	if !authCheck(req, session) {
		ResWriteError(res, fmt.Errorf(TR("User authority is not enough")))
		return
	}
	if err := req.ReadEntity(&request); err != nil {
		ResWriteError(res, err)
		return
	}
	group := types.Group{
		Id:          req.PathParameter("group-id"),
		Name:        request.Name,
		ParentID:    "",
		Path:        "",
		Layer:       "",
		Level:       0,
		Description: request.Description,
		CreateTime:  "",
		Children:    nil,
	}
	if err := s.AuthDB.UpdateGroup(&group); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, nil)
}

// 删除组织
func (s GroupSvc) delGroup(req *restful.Request, res *restful.Response) {
	session := req.HeaderParameter("auth-session")
	if !authCheck(req, session) {
		ResWriteError(res, fmt.Errorf("User authority is not enough"))
		return
	}
	groupId := req.PathParameter("group-id")
	info, err := s.AuthDB.GetGroupInfo(groupId)
	if err != nil {
		ResWriteError(res, err)
	}
	if err := s.AuthDB.GetGroupChild(info); err != nil {
		ResWriteError(res, err)
		return
	}
	if len(info.Children) > 0 {
		ResWriteError(res, fmt.Errorf(TR("Group's child is not empty")))
		return
	}
	if err := s.AuthDB.DelGroup(groupId); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, nil)
}

// 权限检查
func authCheck(req *restful.Request, session string) bool {
	//todo
	return true
}

// 添加组织成员
func (s GroupSvc) addGroupUsers(req *restful.Request, res *restful.Response) {
	var request AddGroupUsersReq
	var resp AddGroupUsersResp
	if err := req.ReadEntity(&request); err != nil {
		log.Errorf("addGroupUsers ReadEntity err: %v", err)
		ResWriteError(res, err)
		return
	}
	groupId := req.PathParameter("group-id")

	// 获取用户名
	for index, user := range request.Users {
		info, err := s.DynamicDB.GetUserById(user.UserID)
		if err != nil {
			log.Errorf("addGroupUsers GetUserById id: %s", user.UserID)
			ResWriteError(res, err)
			return
		}
		//request.Users[index].Name = info.Name
		//request.Users[index].Type = info.Type
		fmt.Println(index, info)

		// 设置用户为组织角色
		roleInfo, err := s.AuthDB.GetRoleInfoByName(types.RoleGroupUser)
		if err != nil || roleInfo.Id == "" {
			log.Println("Failed to get role group user info")
			ResWriteError(res, err)
			return
		}
		userRole := types.UserRole{
			User: info,
			Roles: []types.Role{
				roleInfo,
			},
		}

		if err := s.AuthDB.SetUserRole(userRole); err != nil {
			ResWriteError(res, err)
			return
		}

		// 设置用户type
		//todo 考虑使用int类型然后使用大小判断是否需要修改用户类型
		if info.Type == types.User_Type_Individual {
			if err := s.DynamicDB.UpdateUserType(user.UserID, types.User_Type_Member); err != nil {
				ResWriteError(res, err)
				return
			}
		}
	}
	var users []types.GroupUser
	for _, user := range request.Users {
		users = append(users, types.GroupUser{
			UserID:  user.UserID,
			IsAdmin: user.IsAdmin,
		})
	}

	if err := s.AuthDB.AddGroupUsers(groupId, users); err != nil {
		ResWriteError(res, err)
		return
	}
	ResWriteEntity(res, &resp)
}

// 查询组织成员
func (s GroupSvc) queryGroupUsers(req *restful.Request, res *restful.Response) {
	var request QueryGroupUsersReq
	var resp QueryGroupUsersResp

	request.Group = req.PathParameter("group-id")
	datas, err := s.AuthDB.GetGroupUsers(request.Group)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	for index, data := range datas {
		info, err := s.DynamicDB.GetUserById(data.UserID)
		if err != nil {
			ResWriteError(res, err)
			return
		}
		datas[index].UserID = info.User
		datas[index].Name = info.Name
		datas[index].Email = info.Email
		datas[index].Profile = info.Profile
		datas[index].LastLogin = info.LastLogin
		datas[index].Status = info.Status
		datas[index].CreateTime = info.CreateTime
		datas[index].Type = info.Type

		if roles, err := s.AuthDB.GetUserRole(data.UserID); err != nil {
			datas[index].Roles = roles
		}
	}
	resp.Users = datas
	ResWriteEntity(res, &resp)
}

// 移除组织成员
func (s GroupSvc) delGroupUsers(req *restful.Request, res *restful.Response) {
	var request DelGroupUserReq
	request.Group = req.PathParameter("group-id")
	request.User = req.PathParameter("user-id")
	if err := s.AuthDB.DelGroupUser(request.Group, request.User); err != nil {
		ResWriteError(res, err)
		return
	}
	groups, err := s.AuthDB.GetUserGroup(request.User)
	if err != nil {
		ResWriteError(res, err)
		return
	}
	hasOther := false
	for _, group := range groups {
		if group.Id != request.Group {
			hasOther = true
			break
		}
	}

	// 取消组织角色
	if !hasOther {
		roleInfo, err := s.AuthDB.GetRoleInfoByName(types.RoleGroupUser)
		if err != nil || roleInfo.Id == "" {
			log.Println("Failed to get role group info")
			ResWriteError(res, err)
			return
		}
		if err := s.AuthDB.DelUserRole(request.User, roleInfo.Id); err != nil {
			ResWriteError(res, err)
			return
		}

		// 取消用户member类型
		if err := s.DynamicDB.UpdateUserType(request.User, types.User_Type_Individual); err != nil {
			ResWriteError(res, err)
			return
		}
	}
	ResWriteEntity(res, nil)
}
