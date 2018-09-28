//author xinbing
//time 2018/8/31 14:15
//简单的内存cache，可以支持过期时间设置
package caches

import (
	"fmt"
	"time"
	"sync"
)

// 读锁不阻塞读锁
// 读锁阻塞写锁，直到所以读锁都释放
// 写锁阻塞读\写锁，直到写锁都释放
//待优化
type SimpleCache struct {
	rwMutex *sync.RWMutex //读写锁
	cache map[string]*element //存放的map
	//strategy SimpleStrategy //过期策略
	globalExpire int64 //全局的过期时间
	latestGarbageTime int64 //上一次垃圾回收时间
	hasInitGarbage	  bool
}

const (
	defaultExpire = time.Minute * 30 //半小时
)

// 创建一个SimpleCache，默认过期时间为半小时
func CreateSimpleCache() SimpleCache {
	return CreateSimpleCacheExpire(defaultExpire)
}

// 创建一个指定过期时间的SimpleCache
func CreateSimpleCacheExpire(globalDur time.Duration) SimpleCache {
	sc := SimpleCache{
		rwMutex: new(sync.RWMutex),
		cache: make(map[string]*element),
		globalExpire: int64(globalDur),
	}
	go sc.beginGarbageCollection() //开启垃圾回收,20分钟回收一次
	return sc
}

func (p *SimpleCache) Delete(key string) {
	p.rwMutex.Lock()
	delete(p.cache, key)
	p.rwMutex.Unlock()
}

func (p *SimpleCache) Put(key string,value interface{}) {
	ele := element{
		value: value,
		expire: p.globalExpire,
		createTime: time.Now(),
	}
	p.rwMutex.Lock()
	p.cache[key] = &ele
	p.rwMutex.Unlock()
}

func (p *SimpleCache) PutAssignExpire(key string, value interface{}, dur time.Duration) {
	ele := element{
		value: value,
		expire: int64(dur),
		createTime: time.Now(),
	}
	p.rwMutex.Lock()
	p.cache[key] = &ele
	p.rwMutex.Unlock()
}

func (p *SimpleCache) Get(key string) interface{}{
	p.rwMutex.RLock()
	releasedReadLock := false
	defer func(){
		if !releasedReadLock {
			p.rwMutex.RUnlock()
		}
	}()
	ele := p.cache[key]
	if ele != nil {
		if ele.isExpire() {
			releasedReadLock = true
			p.rwMutex.RUnlock()
			p.Delete(key)
			return nil
		}
		return ele.value
	}
	return nil
}

func (p *SimpleCache) Expire(key string) {
	p.rwMutex.RLock()
	ele := p.cache[key]
	if ele != nil {
		ele.expire = 0
	}
	p.rwMutex.RUnlock()
}

func (p *SimpleCache) beginGarbageCollection() {
	p.rwMutex.Lock()
	if p.hasInitGarbage {
		fmt.Println("has initial garbage!")
		return
	}
	p.hasInitGarbage = true
	p.rwMutex.Unlock()

	ticker := time.NewTicker(time.Minute * 20) //20分钟进行一次
	for {
		select {
		case <-ticker.C:
			p.garbageCollection()
		}
	}
}
//垃圾回收
func (p *SimpleCache) garbageCollection() {
	if len(p.cache) > 100 && time.Now().UnixNano() - p.latestGarbageTime >= p.globalExpire { //大于回收时间
		fmt.Println("begin garbage collection:",time.Now())
		p.rwMutex.Lock() //加锁
		defer func() {
			if err := recover(); err != nil {
				fmt.Println("garbage collection err!")
			}
			p.rwMutex.Unlock()
		}()
		if time.Now().UnixNano() - p.latestGarbageTime < p.globalExpire {
			return
		}
		p.latestGarbageTime = time.Now().UnixNano()
		for key,item := range p.cache {
			if item.isExpire() {
				delete(p.cache, key)
			}
		}
	}
}


type element struct {
	value interface{}
	expire int64
	createTime time.Time
}

func (e *element) isExpire() bool {
	return time.Now().UnixNano() - e.createTime.UnixNano() > e.expire
}
