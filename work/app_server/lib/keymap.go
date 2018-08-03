package lib

import (
	"bytes"
	"encoding/json"
	"github.com/golang/glog"
	"sync"
	"time"
)

var KeyMap *keyMap

type keyMap struct {
	m    map[string]*keyInfo
	lock sync.RWMutex
}

type keyInfo struct {
	k          string
	createTime time.Time
}

func NewKeyMap() {
	KeyMap = &keyMap{
		m: make(map[string]*keyInfo),
	}

	for {
		time.Sleep(20 * time.Second)
		KeyMap.lock.Lock()
		for k, v := range KeyMap.m {
			if time.Now().Sub(v.createTime) >= 5 {
				glog.Infoln("delete", k)
				delete(KeyMap.m, k)
			}
		}
		KeyMap.lock.Unlock()
	}
}

// TODO:锁有点大
func (this *keyMap) Calc(user_desp, user_pass string) string {
	this.lock.Lock()
	defer this.lock.Unlock()

	mp := make(map[string]interface{})
	json.Unmarshal(bytes.NewBufferString(user_desp).Bytes(), &mp)

	addr, ok := mp["address"].(string)
	if !ok {
		glog.Errorln(Loger("keyMap", "can not find user address"))
		return ""
	}

	if _, exist := this.m[addr]; !exist {
		k := ParseKeyStore(user_desp, user_pass)
		if k == "" {
			glog.Errorln("k == nil")
			return ""
		} else {
			this.m[addr] = &keyInfo{
				k:          k,
				createTime: time.Now(),
			}
			return k
		}
	} else {
		return this.m[addr].k
	}
}
