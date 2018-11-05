package etcd

import (
	"context"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	"strings"
	//"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"

	. "grm-labelmgr/types"
	. "grm-service/dbcentral/etcd"
	"grm-service/log"
	"grm-service/util"
)

type DynamicDB struct {
	DynamicEtcd
}

func (e DynamicDB) DeleteUserData(user, dataType, name string) error {
	key := fmt.Sprintf("%s/%s/%s/%s", KEY_GRM_USERDATA, user, dataType, name)
	log.Println("user data key:", key)
	_, err := e.Cli.Delete(context.Background(), key)
	if err != nil {
		log.Println("Failed to delete user data (%s):%v", key, err)
		return err
	}
	return nil
}

func (e DynamicDB) GetUserData(user, dataType, name string) (UserData, error) {
	var data UserData = UserData{}

	key := fmt.Sprintf("%s/%s/%s/%s", KEY_GRM_USERDATA, user, dataType, name)
	log.Println("user data key:", key)
	resp, err := e.Cli.Get(context.Background(), key)
	if err != nil {
		log.Println("Failed to get user data (%s):%v", key, err)
		return data, err
	}
	if len(resp.Kvs) == 0 {
		return data, nil
	}

	data.Data = string(resp.Kvs[0].Value)
	data.DataType = dataType
	data.Name = name
	data.UserId = user
	return data, nil
}

func (e DynamicDB) GetUserDatas(user, dataType string) ([]UserData, error) {
	var datas []UserData = make([]UserData, 0)
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_USERDATA, user, dataType)
	log.Println("user data key:", key)
	resp, err := e.Cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		log.Println("Failed to get user data (%s):%v", key, err)
		return datas, err
	}
	if len(resp.Kvs) == 0 {
		return datas, nil
	}

	for _, ev := range resp.Kvs {
		var data UserData = UserData{}
		data.Data = string(ev.Value)
		ks := strings.Split(string(ev.Key), "/")
		data.DataType = dataType
		data.Name = ks[len(ks)-1]
		data.UserId = user
		datas = append(datas, data)
	}

	return datas, nil
}

func (e DynamicDB) SetUserData(data UserData) error {
	if len(data.Data) == 0 {
		return fmt.Errorf(util.TR("Failed to set user data: ", data))
	}

	key := fmt.Sprintf("%s/%s/%s/%s", KEY_GRM_USERDATA, data.UserId, data.DataType, data.Name)
	log.Println("user data key:", key)

	_, err := e.Cli.Put(context.Background(), key, data.Data)
	if err != nil {
		log.Println("Failed to set user data (%s): %v", key, err)
		return err
	}
	return nil
}
