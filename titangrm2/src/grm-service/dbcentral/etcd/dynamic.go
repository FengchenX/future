package etcd

import (
	"context"
	"fmt"
	"time"

	"github.com/coreos/etcd/clientv3"
	//"github.com/coreos/etcd/etcdserver/api/v3rpc/rpctypes"

	"grm-service/log"
	. "grm-service/util"
)

var (
	dialTimeout    = 5 * time.Second
	requestTimeout = 10 * time.Second
)

const (
	KEY_SPLIT                = "/"
	KEY_SESSION              = "/titanauth/session"
	KEY_USERS                = "/titanauth/users"
	KEY_CAPTCHA              = "/titanauth/captcha"
	KEY_GRM_EXPLORER         = "/titangrm/explorer"
	KEY_GRM_DATAWORKER       = "/titangrm/dataworker"
	KEY_GRM_SEARCHER_HISTORY = "/titangrm/searcher/history"
	KEY_GRM_USERDATA         = "/titangrm/user_data"
	KEY_GRM_STORAGE          = "/titangrm/storage"
	KYE_GRM_COMMENT          = "/titangrm/comments"

	KEY_GRM_TYPE_AGGR = "/titangrm/aggregation"
)

type DynamicEtcd struct {
	Endpoints []string
	Cli       *clientv3.Client
}

func (e *DynamicEtcd) Connect() error {
	if e.Cli != nil {
		return nil
	}
	var err error
	e.Cli, err = clientv3.New(clientv3.Config{
		Endpoints:   e.Endpoints,
		DialTimeout: dialTimeout,
	})
	if err != nil {
		log.Error(err)
		return err
	}
	return nil
}

func (e *DynamicEtcd) IsConnect() bool {
	return e.Cli == nil
}

func (e *DynamicEtcd) DisConnect() error {
	return e.Cli.Close()
}

// 获取用户id
func (e *DynamicEtcd) GetUserId(session string) (string, error) {
	if len(session) == 0 {
		return "", fmt.Errorf(TR("Invalid user session:%s", session))
	}
	//fmt.Println("session key:", KEY_SESSION+"/"+session)
	resp, err := e.Cli.Get(context.Background(), KEY_SESSION+"/"+session)
	if err != nil {
		log.Println("Failed to get user id (%s):", session, err)
		return "", fmt.Errorf(TR("Failed to get user id (%s):", session, err))
	}
	if len(resp.Kvs) == 0 {
		return "", fmt.Errorf(TR("Session is timeout:%s", session))
	}
	return string(resp.Kvs[0].Value), nil
}

// 获取用户名
func (e *DynamicEtcd) GetUserName(id string) (string, error) {
	if len(id) == 0 {
		return "", fmt.Errorf(TR("Invalid user id:%s", id))
	}
	key := fmt.Sprintf("%s/%s", KEY_USERS, id)
	resp, err := e.Cli.Get(context.Background(), key+"/name")
	if err != nil {
		log.Printf("Failed to get user name (%s):\n", key, err)
		return "", err
	}
	if len(resp.Kvs) == 0 {
		log.Println("Failed to get user name :", key)
		return "", fmt.Errorf(TR("Invalid user id:", id))
	}
	return string(resp.Kvs[0].Value), nil
}
