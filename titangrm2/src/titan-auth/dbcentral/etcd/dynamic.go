package etcd

import (
	"context"
	"fmt"
	"strings"
	"time"

	"grm-service/crypto"

	"github.com/coreos/etcd/clientv3"

	. "grm-service/dbcentral/etcd"
	"grm-service/log"
	"grm-service/util"

	. "titan-auth/types"
)

type DynamicDB struct {
	DynamicEtcd
}

// 创建用户session
func (e DynamicDB) CreateUserSession(user string, ttl int64) (string, error) {
	var err error
	session := util.NewUUID()
	if ttl > 0 {
		var leaseResp *clientv3.LeaseGrantResponse
		leaseResp, err = e.Cli.Grant(context.TODO(), ttl)
		if err != nil {
			log.Error("Failed to grant user session ttl:", err)
			return "", err
		}

		if _, err = e.Cli.Put(context.Background(), KEY_SESSION+"/"+session, user,
			clientv3.WithLease(leaseResp.ID)); err != nil {
			log.Error("Failed to create user session:", err)
			return "", err
		}
	} else {
		if _, err = e.Cli.Put(context.Background(), KEY_SESSION+"/"+session, user); err != nil {
			log.Error("Failed to create user session:", err)
			return "", err
		}
	}
	return session, nil
}

// 移除用户session
func (e DynamicDB) DelUserSession(session string) error {
	if _, err := e.Cli.Delete(context.Background(), KEY_SESSION+"/"+session); err != nil {
		return err
	}
	return nil
}

// 注册用户
func (e DynamicDB) AddUser(user *User) (*User, error) {
	if len(user.Id) == 0 {
		user.Id = util.NewUUID()
	}
	if len(user.CreateTime) == 0 {
		user.CreateTime = time.Now().Format("2006-01-02 15:04:05")
	}
	user.Status = User_UnActive
	user.Type = User_Type_Individual
	key := fmt.Sprintf("%s/%s", KEY_USERS, user.Id)

	infoMap := make(map[string]string, 20)
	infoMap["user"] = user.User
	infoMap["name"] = user.Name
	pwd, err := crypto.Md5Encrypt(user.Password)
	if err != nil {
		return nil, err
	}
	infoMap["password"] = pwd
	infoMap["email"] = user.Email
	infoMap["type"] = user.Type
	infoMap["profile"] = user.Profile
	infoMap["create_time"] = user.CreateTime
	infoMap["status"] = user.Status

	for k, v := range infoMap {
		if _, err := e.Cli.Put(context.Background(), key+"/"+k, v); err != nil {
			log.Error("Failed to create user :", err)
			return nil, err
		}
	}
	return user, nil
}

// 修改用户状态
func (e DynamicDB) UpdateUserStatus(user string) error {
	key := fmt.Sprintf("%s/%s", KEY_USERS, user)
	if _, err := e.Cli.Put(context.Background(), key+"/status", User_Active); err != nil {
		return err
	}
	return nil
}

// 修改用户登录时间
func (e DynamicDB) UpdateUserTime(user string) error {
	key := fmt.Sprintf("%s/%s", KEY_USERS, user)
	if _, err := e.Cli.Put(context.Background(), key+"/last_login", time.Now().Format("2006-01-02 15:04:05")); err != nil {
		return err
	}
	return nil
}

// 获取系统所有用户
func (e DynamicDB) GetUserList() (map[string]*User, error) {
	key := fmt.Sprintf("%s", KEY_USERS)
	resp, err := e.Cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	nodes := make(map[string]*User, 100)
	for _, kv := range resp.Kvs {
		value := string(kv.Value)
		keys := strings.Split(string(kv.Key), KEY_SPLIT)
		userId := keys[len(keys)-2]
		if user, ok := nodes[userId]; !ok || user == nil {
			nodes[userId] = &User{Id: userId}
		}

		switch keys[len(keys)-1] {
		case "user":
			nodes[userId].User = value
		case "name":
			nodes[userId].Name = value
		case "password":
			nodes[userId].Password = value
		case "email":
			nodes[userId].Email = value
		case "status":
			nodes[userId].Status = value
		case "type":
			nodes[userId].Type = value
		case "profile":
			nodes[userId].Profile = value
		case "create_time":
			nodes[userId].CreateTime = value
		}
	}
	return nodes, nil
}

// 用户登录
func (e DynamicDB) UserLogin(user, password string) (*User, error) {
	pwd, err := crypto.Md5Encrypt(password)
	if err != nil {
		return nil, err
	}

	users, err := e.GetUserList()
	if err != nil {
		return nil, err
	}

	for _, userInfo := range users {
		if userInfo.User == user || userInfo.Email == user {
			if userInfo.Password != pwd || userInfo.Status != User_Active {
				return nil, fmt.Errorf(util.TR("Invalid user password or user is not active"))
			}
			return userInfo, nil
		}
	}
	return nil, fmt.Errorf(util.TR("Invalid user name for login"))
}
