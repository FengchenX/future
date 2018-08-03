package utils

import (
	"testing"
	"fmt"
	"time"
	"sub_account_service/order_server/entity"
)

func TestSimpleCache_Get(t *testing.T) {
	simpleCache := CreateSimpleCache()
	inter := simpleCache.Get("1")
	fmt.Println(inter)
	simpleCache.Put("1",122223)
	i := GetInt(&simpleCache,"1")
	fmt.Println(*i)
	time.Sleep(time.Second * 2)
	i = GetInt(&simpleCache,"1")
	if i == nil {
		fmt.Println(0)
	}
	fmt.Println(i)
	simpleCache.Put("2",&entity.Company{
		AppId:"123",
		Name:"123333333333333333",
	})
	company := GetCompany(&simpleCache,"2")
	fmt.Println(company)
}

func GetInt(cache *SimpleCache,key string) *int {
	inter := cache.Get(key)
	if inter != nil {
		i,_ := inter.(int)
		return &i
	}
	return nil
}

func GetCompany(cache *SimpleCache,key string) *entity.Company {
	inter := cache.Get(key)
	if inter != nil {
		c,ok := inter.(*entity.Company)
		fmt.Println("err:",ok)
		return c
	}
	return nil
}