package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	//	"github.com/coreos/etcd/clientv3"
	//"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"

	. "grm-service/dbcentral/etcd"
	"grm-service/log"
	"grm-service/util"
	. "titan-statistics/types"
)

type DynamicDB struct {
	DynamicEtcd
}

func (e DynamicDB) GetTypeAggr(typeName string) (*TypeAggr, error) {
	key := fmt.Sprintf("%s/%s", KEY_GRM_TYPE_AGGR, typeName)
	log.Println("type aggr key:", key)
	resp, err := e.Cli.Get(context.Background(), key)
	if err != nil {
		log.Println("Failed to get type aggr (%s):%v", key, err)
		return nil, err
	}

	for _, ev := range resp.Kvs {
		aggr := TypeAggr{}
		aggr.DataType = typeName

		var attr []AggrAttr

		err := json.Unmarshal(ev.Value, &attr)
		if err != nil {
			log.Println("json.Unmarshal error", err)
			continue
		}
		aggr.Aggr = attr
		return &aggr, nil
	}
	return nil, nil
}

func (e DynamicDB) SetTypeAggr(aggr TypeAggr) error {
	if len(aggr.DataType) == 0 {
		return fmt.Errorf(util.TR("Failed to set type aggr: ", aggr))
	}

	key := fmt.Sprintf("%s/%s", KEY_GRM_TYPE_AGGR, aggr.DataType)
	log.Println("type aggr key:", key)

	data, err := json.Marshal(aggr.Aggr)
	if err != nil {
		log.Println("json.Marshal error", err)
		return err
	}

	_, err = e.Cli.Put(context.Background(), key, string(data))
	if err != nil {
		log.Println("Failed to set type aggr (%s): %v", key, err)
		return err
	}
	return nil
}
