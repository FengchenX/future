package pg

import (
	"fmt"
	"github.com/jackc/pgx"
	"grm-service/log"
	"grm-service/util"
	"time"
	"titan-auth/types"
)

func (db AuthDB) GetRoleInfoByName(name string) (types.Role, error) {
	var role types.Role
	sql := fmt.Sprintf(`select id, label, create_time, description from sys_role where name = '%s'`, name)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return role, err
	}
	defer rows.Close()

	var id, label, description string
	var create_time time.Time
	for rows.Next() {
		err = rows.Scan(&id, &label, &create_time, &description)
		if err != nil {
			log.Printf("GetRoles: rows.Scan error: %s\n", err.Error())
			return role, err
		} else {
			role = types.Role{
				Id:          id,
				Name:        name,
				Label:       label,
				Description: description,
				CreateTime:  util.GetTimeStd(create_time),
			}
		}
	}
	if err := rows.Err(); err != nil {
		return role, err
	}
	return role, nil
}

func (db AuthDB) SetUserRole(user types.UserRole) error {
	for _, role := range user.Roles {
		sql := fmt.Sprintf(`insert into ref_user_role(user_id, role_id, create_time)
									values('%s', '%s', %s)`, user.User.Id, role.Id, util.GetTimeNowDB())
		if _, err := db.Conn.Exec(sql); err != nil {
			if pgErr, ok := err.(pgx.PgError); ok {
				if pgErr.Code == "23505" {
					continue
				}
			}
			log.Printf("SetUserRole error: %s\n", err.Error())
			return err
		}
	}
	return nil
}

func (db AuthDB) GetUserRole(userId string) ([]types.Role, error) {
	var roles []types.Role
	sql := fmt.Sprintf(`select ref_user_role.role_id, sys_role.label, sys_role.name
						from ref_user_role inner join sys_role on ref_user_role.role_id = sys_role.id 
						where ref_user_role.user_id = '%s'`, userId)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return roles, err
	}
	defer rows.Close()

	var role_id, label, name string
	for rows.Next() {
		err = rows.Scan(&role_id, &label, &name)
		if err != nil {
			log.Printf("GetRoles: rows.Scan error: %s\n", err.Error())
			continue
		} else {
			role := types.Role{
				Id:          role_id,
				Name:        name,
				Label:       label,
				Description: "",
				CreateTime:  "",
			}
			roles = append(roles, role)
		}
	}
	if err := rows.Err(); err != nil {
		return roles, err
	}
	return roles, nil
}

func (db AuthDB) DelUserRole(userId, roleId string) error {
	sql := fmt.Sprintf(`delete from ref_user_role where user_id = '%s' and role_id = '%s'`, userId, roleId)
	_, err := db.Conn.Exec(sql)
	return err
}
