//author xinbing
//time 2018/8/31 14:32
package caches

import (
	"common-utilities/utilities"
	"testing"
	"fmt"
	"sync"
	"time"
)

var testingSimpleCache = CreateSimpleCacheExpire(time.Minute * 5)
func TestSimpleCache(t *testing.T) {
	//r := rand.New(rand.NewSource(time.Now().UnixNano()))
	key := "zhangs"
	for j:=0;j<1000;j++ {
		if j%2 == 0 {
			go testingSimpleCache.Get(key)
		} else if j%3 == 0 {
			go testingSimpleCache.Put(key, j)
		} else if j%4 == 0 {
			go testingSimpleCache.Expire(key)
		} else if j%17 == 0 {
			go testingSimpleCache.Delete(key)
		}
	}
	fmt.Println(testingSimpleCache.Get(key))
}
func TestSimpleCache2(t *testing.T) {
	for i:=0; i<10; i++ {
		go func() {
			value := testingSimpleCache.Get("zhangs")
			fmt.Println(value)
		} ()
		go func(z int) {
			testingSimpleCache.Put("zhangs",z)
		}(i)
	}
}

func TestGarbageCollection(t *testing.T) {
	cache := CreateSimpleCache()

	for i:=0;i<100;i++ {
		cache.Put(utilities.GetRandomNumStr(4),"1")
	}
	fmt.Println("begin cache size:",len(cache.cache))
	time.Sleep(30 * time.Second)
	fmt.Println("end cache size:",len(cache.cache))
}
func TestRwMutex(t *testing.T) {
	rw := sync.RWMutex{}
	// ---deadlock，write lock lock write lock
	//rw.Lock()
	//rw.Lock()
	//rw.Unlock()
	//rw.Unlock()
	// --deadlock，write lock lock read lock
	//rw.Lock()
	//rw.RLock()
	//rw.RUnlock()
	//rw.Unlock()
	// --deadlock read lock lock write lock
	//rw.RLock()
	//rw.Lock()
	//rw.Unlock()
	//rw.RUnlock()

	// --normal
	rw.RLock()
	rw.RLock()
	rw.RUnlock()
	rw.RUnlock()
}

func TestRWMutex(t *testing.T) {

}