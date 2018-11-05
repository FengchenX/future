package pg

import (
	"fmt"
	"github.com/jackc/pgx"
	"grm-service/log"
	"grm-service/util"
	"time"
	"titan-auth/types"
)

// 添加组织用户
func (db AuthDB) AddGroupUsers(groupId string, users []types.GroupUser) error {
	for _, user := range users {
		sql := fmt.Sprintf(`insert into ref_group_user(user_id, group_id, is_admin, join_time)
							values('%s', '%s', %t, %s)`, user.UserID, groupId, user.IsAdmin, util.GetTimeNowDB())
		if _, err := db.Conn.Exec(sql); err != nil {
			if pgErr, ok := err.(pgx.PgError); ok {
				if pgErr.Code == "23505" {
					continue
				}
			}
			return err
		}
	}
	return nil
}

func (db AuthDB) GetGroupUsers(groupId string) ([]types.GroupUser, error) {
	var users []types.GroupUser
	sql := fmt.Sprintf(`select user_id, group_id, is_admin, join_time
						from ref_group_user where group_id = '%s' order by is_admin desc`, groupId)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return users, err
	}
	defer rows.Close()

	var user_id, group_id string
	var create_time time.Time
	var is_admin bool

	for rows.Next() {
		err = rows.Scan(&user_id, &group_id, &is_admin, &create_time)
		if err != nil {
			log.Printf("GetGroupUsers: rows.Scan error: %s\n", err.Error())
			continue
		} else {
			user := types.GroupUser{
				GroupID:    group_id,
				UserID:     user_id,
				IsAdmin:    is_admin,
				JoinTime:   util.GetTimeStd(create_time),
				Name:       "",
				Type:       "",
				CreateTime: "",
				Email:      "",
				Profile:    "",
				LastLogin:  "",
				Status:     "",
				Roles:      nil,
			}
			users = append(users, user)
		}
	}
	if err := rows.Err(); err != nil {
		return users, err
	}
	return users, nil
}

func (db AuthDB) DelGroupUser(groupId, userId string) error {
	sql := fmt.Sprintf(`delete from ref_group_user where user_id = '%s' and group_id = '%s'`, userId, groupId)
	_, err := db.Conn.Exec(sql)
	return err
}

func (db AuthDB) GetUserGroup(userId string) ([]types.Group, error) {
	var groups []types.Group
	sql := fmt.Sprintf(`select ref_group_user.group_id, sys_group.name
						from ref_group_user inner join sys_group on ref_group_user.group_id = sys_group.id
						where ref_group_user.user_id = '%s'`, userId)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return groups, err
	}
	defer rows.Close()

	var group_id, name string
	for rows.Next() {
		err = rows.Scan(&group_id, &name)
		if err != nil {
			log.Printf("GetRoles: rows.Scan error: %s\n", err.Error())
			continue
		} else {
			group := types.Group{
				Id:          group_id,
				Name:        name,
				ParentID:    "",
				Path:        "",
				Layer:       "",
				Level:       0,
				Description: "",
				CreateTime:  "",
				Children:    nil,
				Users:       nil,
			}
			groups = append(groups, group)
		}
	}

	if err := rows.Err(); err != nil {
		return groups, err
	}
	return groups, nil
}
