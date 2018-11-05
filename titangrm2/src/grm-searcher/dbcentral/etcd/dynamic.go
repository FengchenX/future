package etcd

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/coreos/etcd/clientv3"
	//"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"

	. "grm-searcher/types"
	. "grm-service/dbcentral/etcd"
	"grm-service/log"
	"grm-service/util"
)

type DynamicDB struct {
	DynamicEtcd
}

func (e DynamicDB) DeleteUserHistory(user, id string) error {
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_SEARCHER_HISTORY, user, id)
	log.Println("history key:", key)
	_, err := e.Cli.Delete(context.Background(), key)
	if err != nil {
		log.Println("Failed to delete explorer (%s):%v", key, err)
		return err
	}
	return nil
}

func (e DynamicDB) GetUserHistory(user, id string) (History, error) {
	var history History
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_SEARCHER_HISTORY, user, id)
	log.Println("history key:", key)
	resp, err := e.Cli.Get(context.Background(), key)
	if err != nil {
		log.Println("Failed to get explorer (%s):%v", key, err)
		return history, err
	}
	if len(resp.Kvs) == 0 {
		return history, nil
	}

	err = json.Unmarshal(resp.Kvs[0].Value, &history)
	if err != nil {
		log.Println("Failed to get historys:", err)
		return history, err
	}
	return history, nil
}

func (e DynamicDB) GetUserHistorys(user string) (HistoryList, error) {
	key := fmt.Sprintf("%s/%s", KEY_GRM_SEARCHER_HISTORY, user)
	log.Println("history user key:", key)
	resp, err := e.Cli.Get(context.Background(), key, clientv3.WithPrefix())
	if err != nil {
		log.Println("Failed to get historys (%s):%v", key, err)
		return HistoryList{}, err
	}

	list := HistoryList{}
	for _, ev := range resp.Kvs {
		var history History
		err := json.Unmarshal(ev.Value, &history)
		if err != nil {
			log.Println("Failed to get historys (%s):%v", string(ev.Key), err)
		} else {
			if history.SearchInfo.Geometry != nil && len(history.SearchInfo.Geometry.Type) == 0 {
				history.SearchInfo.Geometry = nil
			}

			list = append(list, history)
		}
	}
	return list, nil
}

func (e DynamicDB) SetHistory(history History) error {
	if len(history.Path) == 0 {
		return fmt.Errorf(util.TR("Failed to set history: ", history))
	}

	history.Id = util.NewUUID()
	key := fmt.Sprintf("%s/%s/%s", KEY_GRM_SEARCHER_HISTORY, history.Userid, history.Id)
	log.Println("searcher history key:", key)

	data, err := json.Marshal(history)
	if err != nil {
		return err
	}

	_, err = e.Cli.Put(context.Background(), key, string(data))
	if err != nil {
		log.Println("Failed to set explorer (%s): %v", key, err)
		return err
	}
	return nil
}
