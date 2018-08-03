package test

import (
	"testing"
	"sync"
	"math/big"
	"fmt"
)
type element struct {
	c chan *big.Int
}
var cache = sync.Map{}

func TestSyncMap(t *testing.T) {
	key := "a"
	elePointer, ok := cache.Load(key)
	if !ok {
		mutex.Lock()
		defer func() {
			if err := recover();err != nil {
			}
			mutex.Unlock()
		}()
		ele := &element{
			c:make(chan *big.Int,1),
		}
		cache.Store(key,ele)
		ele.c <- big.NewInt(23)
		elePointer,_ = cache.Load(key)
	}
	ele := elePointer.(*element)
	fmt.Println(<-ele.c)
}

var mutex = sync.Mutex{}
