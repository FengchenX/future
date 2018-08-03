package utils

import "time"

//待优化
type SimpleCache struct {
	cache map[string]*element
	globalExpire int64
}

func CreateSimpleCache() SimpleCache {
	return CreateSimpleCacheExpire(defaultExpire)
}

func CreateSimpleCacheExpire(globalExpire int64) SimpleCache {
	return SimpleCache{
		cache:make(map[string]*element),
		globalExpire: defaultExpire,
	}
}

func (p *SimpleCache) Put(key string,value interface{}) {
	ele := element{
		value: value,
		expire: p.globalExpire,
		createTime: time.Now(),
	}
	p.cache[key] = &ele
}
func (p *SimpleCache) Get(key string) interface{}{
	ele := p.cache[key]
	if ele != nil {
		if ele.isExpire() {
			delete(p.cache,key)
			return nil
		}
		return ele.value
	}
	return nil
}
const (
	defaultExpire = 60*30 //半小时
)
type element struct {
	value interface{}
	expire int64
	createTime time.Time
}

func (e *element) isExpire() bool {
	return time.Now().Unix() - e.createTime.Unix() > e.expire
}
