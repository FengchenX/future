package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"grm-service/dbcentral/etcd"
	"strings"
	"titan-auth/types"
)

func (e DynamicDB) GetUserById(userId string) (types.User, error) {
	key := fmt.Sprintf("%s%s", etcd.KEY_USERS, "/"+userId)
	resp, err := e.Cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return types.User{}, err
	}

	var user types.User
	for _, ev := range resp.Kvs {
		key := string(ev.Key)
		key = key[strings.LastIndex(key, "/")+1:]
		value := string(ev.Value)

		switch key {
		case "user":
			user.User = value
		case "name":
			user.Name = value
		case "password":
			user.Password = value
		case "email":
			user.Email = value
		case "type":
			user.Type = value
		case "profile":
			user.Profile = value
		case "create_time":
			user.CreateTime = value
		case "status":
			user.Status = value
		}
	}
	return user, nil
}

//修改用户类型
func (e DynamicDB) UpdateUserType(userId, userType string) error {
	key := fmt.Sprintf("%s%s", etcd.KEY_USERS, "/"+userId)
	if _, err := e.Cli.Put(context.Background(), key+"/type", userType); err != nil {
		return err
	}
	return nil
}
