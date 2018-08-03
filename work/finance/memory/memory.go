package memory

import (
	"container/list"
	"errors"
	"fmt"
	"sync"
	"time"

	"sub_account_service/finance/session"
)

//SessionStore session对象，存储用户信息
type SessionStore struct {
	sid          string                      //session id 唯一标示
	timeAccessed time.Time                   //最后访问时间
	value        map[interface{}]interface{} //session 里面存储的值
}

//Set 设置用户信息的key 和 value eg: Set("username", "我妻由乃")
func (st *SessionStore) Set(key, value interface{}) error {
	st.value[key] = value
	pder.SessionUpdate(st.sid)
	return nil
}

//Get 从session中取出用户信息 eg: Get("username")
func (st *SessionStore) Get(key interface{}) interface{} {
	pder.SessionUpdate(st.sid)
	if v, ok := st.value[key]; ok {
		return v
	} else {
		return nil
	}
}

//Delete 从session中删除用户某项信息
func (st *SessionStore) Delete(key interface{}) error {
	delete(st.value, key)
	pder.SessionUpdate(st.sid)
	return nil
}

//SessionID 取出session ID
func (st *SessionStore) SessionID() string {
	return st.sid
}

//Provider session提供者
type Provider struct {
	lock     sync.Mutex               //用来锁
	sessions map[string]*list.Element //用来存储在内存
	list     *list.List               //用来做 gc
}

//SessionInit 生成session 并添加到session池中
func (provider *Provider) SessionInit(sid string) (session.Session, error) {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	v := make(map[interface{}]interface{}, 0)
	newsess := &SessionStore{sid: sid, timeAccessed: time.Now(), value: v}
	element := provider.list.PushBack(newsess)
	provider.sessions[sid] = element
	return newsess, nil
}

//SessionRead 从session池中取出session
func (provider *Provider) SessionRead(sid string) (session.Session, error) {
	if element, ok := provider.sessions[sid]; ok {
		return element.Value.(*SessionStore), nil
	} else {
		//sess, err := provider.SessionInit(sid)
		err := errors.New("token is error")
		return nil, err
	}
}

//SessionDestroy session销毁
func (provider *Provider) SessionDestroy(sid string) error {
	if element, ok := provider.sessions[sid]; ok {
		delete(provider.sessions, sid)
		provider.list.Remove(element)
		return nil
	}
	return nil
}

//SessionGC session GC
func (provider *Provider) SessionGC(maxLifeTime int64) {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	for {
		element := provider.list.Back()
		if element == nil {
			break
		}
		if (element.Value.(*SessionStore).timeAccessed.Unix() + maxLifeTime) <
			time.Now().Unix() {
			provider.list.Remove(element)
			delete(provider.sessions, element.Value.(*SessionStore).sid)
		} else {
			break
		}
	}
}

//SessionUpdate 更新session
func (provider *Provider) SessionUpdate(sid string) error {
	provider.lock.Lock()
	defer provider.lock.Unlock()
	if element, ok := provider.sessions[sid]; ok {
		element.Value.(*SessionStore).timeAccessed = time.Now()
		provider.list.MoveToFront(element)
		return nil
	}
	return nil
}

var pder = &Provider{list: list.New()}

func init() {
	pder.sessions = make(map[string]*list.Element, 0)
	session.Register("memory", pder)
	fmt.Println("wzz")
}
