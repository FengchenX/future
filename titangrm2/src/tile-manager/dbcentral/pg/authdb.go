package pg

import (
	"fmt"
	"grm-service/common"
	"grm-service/dbcentral/pg"
	"grm-service/log"
)

type AuthDB struct {
	pg.AuthCentralDB
}

func (db AuthDB) GetUserInfo(role string) ([]common.UserInfo, error) {
	sql := fmt.Sprintf(
		`SELECT
				ref_user_role.user_id, sys_role."name"
			FROM
				ref_user_role
				INNER JOIN sys_role ON ref_user_role.role_id = sys_role."id" 
			WHERE
				sys_role."name" = '%s'`, role)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return nil, err
	}
	var users []common.UserInfo
	var user_id string
	for rows.Next() {
		if err := rows.Scan(&user_id); err != nil {
			log.Errorf("rows.Scan error: %s\n", err.Error())
			continue
		}
		users = append(users, common.UserInfo{Id: user_id})
	}
	if rows.Err() != nil {
		return users, rows.Err()
	}
	return users, nil
}
