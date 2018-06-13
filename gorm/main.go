

package main

import (
	"fmt"
    "github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

func main() {
  	db, err := gorm.Open("mysql", "root:root@tcp(39.108.80.66:3306)/test_order?charset=utf8&parseTime=true&loc=Local")
  	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	//分页查询
	//连表查询
	var count int
	size := 2
	num := 1
	var bs []TestB
	db.Table("test_bs").Joins("right join test_as on test_bs.a_id = test_as.id").Count(&count)
	db.Table("test_bs").Joins("right join test_as on test_bs.a_id = test_as.id").Where("name = ?", "a1").Limit(size).Offset((num-1)*size).Find(&bs)
	fmt.Println(bs)
	//连表查，并将两个表拼接起来
	var abs []struct{Name string; Age string; Addr string}
	db.Table("test_as").Joins("left join test_bs on test_as.id = test_bs.a_id").Select("test_as.name, test_as.age, test_bs.addr").
	Where("name = ?", "a1").Limit(size).Offset((num-1)*size).Find(&abs)
	fmt.Println(abs)
}

//TestA 主表
type TestA struct {
	ID uint `gorm:"primary_key" json:"-"`
	Name string
	Age string
}

//TestB 详情表
type TestB struct {
	ID uint `gorm:"primary_key" json:"-"`
	AID uint `gorm:"not null;unique"`
	Addr string
}