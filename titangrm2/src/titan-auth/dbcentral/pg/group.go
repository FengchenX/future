package pg

import (
	"fmt"
	"grm-service/log"
	"grm-service/util"
	"time"
	"titan-auth/types"
)

// 创建组织
func (db AuthDB) CreateGroup(group types.Group) (types.Group, error) {
	sql := fmt.Sprintf(`insert into sys_group(id, name,parent_id,layer,level,description,create_time) 
							values ('%s','%s','%s','%s',%d,'%s',%s)`,
		group.Id, group.Name, group.ParentID, group.Layer,
		group.Level, group.Description, util.GetTimeNowDB())
	_, err := db.Conn.Exec(sql)
	if err != nil {
		return group, err
	}

	// 添加组织用户
	if len(group.Users) > 0 {
		for _, user := range group.Users {
			sql := fmt.Sprintf(`insert into ref_group_user(user_id,group_id,is_admin,join_time) 
							values ('%s','%s','%t',%s)`,
				user.UserID, group.Id, user.IsAdmin, util.GetTimeNowDB())
			if _, err := db.Conn.Exec(sql); err != nil {
				return group, err
			}
		}
	}
	return group, nil
}

// 获取所有组织
func (db AuthDB) QueryGroups() ([]types.Group, error) {
	// 根节点组织
	var rootGroups []types.Group
	sql := fmt.Sprintf(`select id, name, parent_id, layer, level, description,
								create_time from sys_group where level = 1 and parent_id = '0'
								order by create_time`)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return rootGroups, err
	}
	defer rows.Close()

	var id, parent_id, name, layer, description string
	var create_time time.Time
	var level int
	for rows.Next() {
		err = rows.Scan(&id, &name, &parent_id, &layer, &level, &description, &create_time)
		if err != nil {
			log.Printf("GetRootGroups: rows.Scan error: %s\n", err.Error())
			continue
		} else {
			group := types.Group{
				Id:          id,
				Name:        name,
				ParentID:    parent_id,
				Path:        name,
				Layer:       layer,
				Level:       level,
				CreateTime:  util.GetTimeStd(create_time),
				Description: description,
			}
			rootGroups = append(rootGroups, group)
		}
	}
	if err := rows.Err(); err != nil {
		return rootGroups, err
	}

	// 获取所有根节点下所有节点
	for index, _ := range rootGroups {
		if err := db.GetGroupChildren(&rootGroups[index]); err != nil {
			return rootGroups, err
		}
	}
	return rootGroups, nil
}

// 修改组织
func (db AuthDB) UpdateGroup(group *types.Group) error {
	comma := ""
	sql := `update sys_group set `
	if group.Name != "" {
		sql = fmt.Sprintf(`%s name = '%s'`, sql, group.Name)
		comma = ","
	}
	if group.Description != "" {
		sql = fmt.Sprintf(`%s %s description = '%s'`, sql, comma, group.Description)
		comma = ","
	}
	if comma == "" {
		return nil
	}
	sql = fmt.Sprintf("%s where id = '%s'", sql, group.Id)
	_, err := db.Conn.Exec(sql)
	return err
}

// 删除组织
func (db AuthDB) DelGroup(groupId string) error {
	sql := fmt.Sprintf(`delete from sys_group where id = '%s'`, groupId)
	_, err := db.Conn.Exec(sql)
	return err
}

func (db AuthDB) GetGroupInfo(groupId string) (*types.Group, error) {
	var group types.Group
	sql := fmt.Sprintf(`select name, parent_id, layer, level, description,
								create_time from sys_group where id = '%s'`, groupId)
	//fmt.Println(sql)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var parent_id, name, layer, description string
	var create_time time.Time
	var level int
	if rows.Next() {
		err = rows.Scan(&name, &parent_id, &layer, &level, &description, &create_time)
		if err != nil {
			log.Printf("GetGroupInfo: rows.Scan error: %s\n", err.Error())
			return nil, err
		} else {
			group = types.Group{
				Id:          groupId,
				Name:        name,
				ParentID:    parent_id,
				Path:        "",
				Layer:       layer,
				Level:       level,
				Description: description,
				CreateTime:  util.GetTimeStd(create_time),
				Children:    nil,
			}
		}
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}
	return &group, nil
}

func (db AuthDB) GetGroupChildren(group *types.Group) error {
	format := fmt.Sprintf("%s%s-", group.Layer, group.Id)
	//level 限制只取下一级
	sql := fmt.Sprintf(`select id,name,parent_id,layer,level,description,
								create_time from sys_group where level = %d and layer like '%s' 
								order by create_time`,
		group.Level+1, format+"%")
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()
	var id, parent_id, name, layer, description string
	var create_time time.Time
	var level int
	for rows.Next() {
		err = rows.Scan(&id, &name, &parent_id, &layer, &level, &description, &create_time)
		if err != nil {
			log.Printf("GetGroupChildren: rows.Scan error: %s\n", err.Error())
			continue
		} else {
			child := types.Group{
				Id:          id,
				Name:        name,
				ParentID:    parent_id,
				Path:        fmt.Sprintf("%s/%s", group.Path, name),
				Layer:       layer,
				Level:       level,
				Description: description,
				CreateTime:  util.GetTimeStd(create_time),
				Children:    nil,
				Users:       nil,
			}
			group.Children = append(group.Children, child)
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}

	// 递归
	if len(group.Children) > 0 {
		for index, _ := range group.Children {
			if err := db.GetGroupChildren(&group.Children[index]); err != nil {
				return err
			}
		}
	}
	return nil
}

func (db AuthDB) GetGroupChild(group *types.Group) error {
	format := fmt.Sprintf("%s%s-", group.Layer, group.Id)
	sql := fmt.Sprintf(`select id,name,parent_id,layer,level,description,
								create_time from sys_group where level = %d and layer like '%s' 
								order by create_time`,
		group.Level+1, format+"%")
	fmt.Println(sql)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return err
	}
	defer rows.Close()

	var id, parent_id, name, layer, description string
	var create_time time.Time
	var level int

	for rows.Next() {
		err = rows.Scan(&id, &name, &parent_id, &layer, &level, &description, &create_time)
		if err != nil {
			log.Printf("GetGroupChildren: rows.Scan error: %s\n", err.Error())
			continue
		} else {
			child := types.Group{
				Id:          id,
				Name:        name,
				ParentID:    parent_id,
				Path:        fmt.Sprintf("%s/%s", group.Path, name),
				Layer:       layer,
				Level:       level,
				CreateTime:  util.GetTimeStd(create_time),
				Description: description,
			}
			group.Children = append(group.Children, child)
		}
	}
	if err := rows.Err(); err != nil {
		return err
	}
	return nil
}
