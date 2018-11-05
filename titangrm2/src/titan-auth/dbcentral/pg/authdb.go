package pg

import (
	"fmt"
	"time"

	//"github.com/jackc/pgx"
	_ "github.com/jackc/pgx/stdlib"

	. "grm-service/dbcentral/pg"
	. "grm-service/time"
	. "titan-auth/types"

	"grm-service/crypto"
	"grm-service/log"
)

type AuthDB struct {
	AuthCentralDB
}

// 用户登录
func (db AuthDB) UserLogin(user, password string) (User, error) {
	var info User
	pwd, err := crypto.Md5Encrypt(password)
	if err != nil {
		return info, err
	}

	sql := fmt.Sprintf(`select user_id,user_name,name,email,type,profile,create_time,last_login,
						status from sys_user where user_name='%s' or email = '%s' 
						and status != '%s' and password = '%s'`,
		user, user, User_Obsoleted, pwd)
	log.Info(sql)
	rows, err := db.Conn.Query(sql)
	if err != nil {
		return info, err
	}
	defer rows.Close()

	var create_time, last_login time.Time
	if rows.Next() {
		err = rows.Scan(&info.Id, &info.User, &info.Name, &info.Email,
			&info.Type, &info.Profile, &create_time, &last_login, &info.Status /*, &info.Organization, &info.Department*/)
		if err != nil {
			log.Errorf("user login: rows.Scan error: %s\n", err.Error())
			return info, err
		}
		info.CreateTime = GetTimeStd(create_time)
		info.LastLogin = GetTimeStd(last_login)
	}
	if err := rows.Err(); err != nil {
		return info, err
	}
	return info, nil
}
