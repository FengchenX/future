package captcha

import (
	"context"
	"fmt"

	"github.com/coreos/etcd/clientv3"

	. "grm-service/dbcentral/etcd"
	"grm-service/log"
)

//customizeRdsStore An object implementing Store interface
type EtcdStore struct {
	Cli *clientv3.Client
}

// customizeRdsStore implementing Set method of  Store interface
func (s *EtcdStore) Set(id string, value string) {
	key := fmt.Sprintf(`%s/%s`, KEY_CAPTCHA, id)

	var leaseResp *clientv3.LeaseGrantResponse
	leaseResp, err := s.Cli.Grant(context.TODO(), 10*60)
	if err != nil {
		log.Error("Failed to grant captcha ttl:", err)

		if _, err = s.Cli.Put(context.Background(), key, value); err != nil {
			log.Error("Failed to create captcha:", err)
		}
		return
	}

	if _, err = s.Cli.Put(context.Background(), key, value, clientv3.WithLease(leaseResp.ID)); err != nil {
		log.Error("Failed to create captcha:", err)
	}
}

// customizeRdsStore implementing Get method of  Store interface
func (s *EtcdStore) Get(id string, clear bool) (value string) {
	key := fmt.Sprintf(`%s/%s`, KEY_CAPTCHA, id)
	res, err := s.Cli.Get(context.Background(), key)
	if err != nil {
		log.Error("Failed to get captcha:", err)
		return ""
	}
	if len(res.Kvs) == 0 {
		log.Error("Failed to get captcha:", err)
		return ""
	} else {
		if clear {
			_, err := s.Cli.Delete(context.Background(), key)
			if err != nil {
				log.Error("Captcha.EtcdStore.Delete error:", err)
			}
		}
	}
	return string(res.Kvs[0].Value)
}
