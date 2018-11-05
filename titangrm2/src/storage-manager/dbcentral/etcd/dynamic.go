package etcd

import (
	"context"
	"fmt"
	"strings"

	"github.com/coreos/etcd/clientv3"

	. "grm-service/dbcentral/etcd"
	"grm-service/log"
	"storage-manager/types"
)

type DynamicDB struct {
	DynamicEtcd
}

// 更新存储使用情况
func (e DynamicDB) UpdateStorage(id int, used, userPercent string) error {
	key := fmt.Sprintf("%s/%d", KEY_GRM_STORAGE, id)
	if _, err := e.Cli.Put(context.Background(), key+"/used", used); err != nil {
		log.Error("Failed to update storage info(%s) : %s", key+"/used", err)
		return err
	}

	if _, err := e.Cli.Put(context.Background(), key+"/used_percent", userPercent); err != nil {
		log.Error("Failed to update storage info(%s) : %s", key+"/used_percent", err)
		return err
	}

	return nil
}

// 获取存储使用情况
func (e DynamicDB) GetStorage(dev *types.Device) error {
	key := fmt.Sprintf("%s/%d", KEY_GRM_STORAGE, dev.Id)
	resp, err := e.Cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		return err
	}

	for _, ev := range resp.Kvs {
		key := string(ev.Key)
		key = key[strings.LastIndex(key, "/")+1:]
		value := string(ev.Value)

		switch key {
		case "used":
			dev.Used = value
		case "used_percent":
			dev.UsedPercent = value
		}
	}
	return nil
}
