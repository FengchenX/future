package etcd

import (
	"context"
	"fmt"

	//"github.com/coreos/etcd/clientv3"
	//"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"

	. "grm-service/dbcentral/etcd"
	"grm-service/log"
	. "grm-service/util"
)

type DynamicDB struct {
	DynamicEtcd
}

// 获取explorer
func (e DynamicDB) GetExplorer(explorer string) (string, error) {
	if len(explorer) == 0 {
		return "", fmt.Errorf(TR("Failed to get explorer: ", explorer))
	}

	key := fmt.Sprintf("%s/%s/path", KEY_GRM_EXPLORER, explorer)
	log.Println("explorer key:", key)
	resp, err := e.Cli.Get(context.Background(), key)
	if err != nil {
		log.Println("Failed to get explorer (%s):%v", key, err)
		return "", err
	}
	if len(resp.Kvs) == 0 {
		//return "", fmt.Errorf(TR("Failed to get explorer:%s", explorer))
		return "", nil
	}
	return string(resp.Kvs[0].Value), nil
}

// 更新explorer
func (e DynamicDB) SetExplorer(explorer, path string) error {
	if len(explorer) == 0 {
		return fmt.Errorf(TR("Failed to set explorer: ", explorer))
	}

	key := fmt.Sprintf("%s/%s/path", KEY_GRM_EXPLORER, explorer)
	log.Println("explorer key:", key)
	_, err := e.Cli.Put(context.Background(), key, path)
	if err != nil {
		log.Println("Failed to get explorer (%s): %v", key, err)
		return err
	}
	return nil
}
