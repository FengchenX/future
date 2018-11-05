package user

import (
	"sort"

	"titan-auth/types"
)

// 用户排序
type UserlList []*types.User

func (s UserlList) Len() int { return len(s) }

func (s UserlList) Swap(i, j int) { s[i], s[j] = s[j], s[i] }

func userSort(user []*types.User, sortArgs, order string) {
	switch sortArgs {
	case "name":
		{
			if order == "desc" {
				sort.Sort(ByUserNameDesc{user})
			} else {
				sort.Sort(sort.Reverse(ByUserNameDesc{user}))
			}
		}
	case "email":
		{
			if order == "desc" {
				sort.Sort(ByUserEmailDesc{user})
			} else {
				sort.Sort(sort.Reverse(ByUserEmailDesc{user}))
			}
		}
	case "type":
		{
			if order == "desc" {
				sort.Sort(ByUserTypeDesc{user})
			} else {
				sort.Sort(sort.Reverse(ByUserTypeDesc{user}))
			}
		}
	case "status":
		{
			if order == "desc" {
				sort.Sort(ByUserStatusDesc{user})
			} else {
				sort.Sort(sort.Reverse(ByUserStatusDesc{user}))
			}
		}
	case "create_time":
		{
			if order == "desc" {
				sort.Sort(ByRegTimeDesc{user})
			} else {
				sort.Sort(sort.Reverse(ByRegTimeDesc{user}))
			}
		}
	case "last_login":
		{
			if order == "desc" {
				sort.Sort(ByLoginTimeDesc{user})
			} else {
				sort.Sort(sort.Reverse(ByLoginTimeDesc{user}))
			}
		}
	}
}

// 用户名
type ByUserNameDesc struct{ UserlList }

func (s ByUserNameDesc) Less(i, j int) bool {
	return s.UserlList[i].Name > s.UserlList[j].Name
}

// 邮箱
type ByUserEmailDesc struct{ UserlList }

func (s ByUserEmailDesc) Less(i, j int) bool {
	return s.UserlList[i].Email > s.UserlList[j].Email
}

// 类型
type ByUserTypeDesc struct{ UserlList }

func (s ByUserTypeDesc) Less(i, j int) bool {
	return s.UserlList[i].Type > s.UserlList[j].Type
}

// 状态
type ByUserStatusDesc struct{ UserlList }

func (s ByUserStatusDesc) Less(i, j int) bool {
	return s.UserlList[i].Status > s.UserlList[j].Status
}

// 注册时间
type ByRegTimeDesc struct{ UserlList }

func (s ByRegTimeDesc) Less(i, j int) bool {
	return s.UserlList[i].CreateTime > s.UserlList[j].CreateTime
}

// 登录时间
type ByLoginTimeDesc struct{ UserlList }

func (s ByLoginTimeDesc) Less(i, j int) bool {
	return s.UserlList[i].LastLogin > s.UserlList[j].LastLogin
}
