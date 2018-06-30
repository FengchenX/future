package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)

/*
select *
from expenses_bills
-- join user_bills on user_bills.bill_id = expenses_bills.id
-- where 
limit 2,1

*/

func main() {

	db, err := gorm.Open("mysql", "root:root@tcp(39.108.80.66:3306)/finance?charset=utf8&parseTime=true&loc=Local")
	defer db.Close()
	if err != nil {
		log.Fatalln(err)
	}
	/*
	//db.CreateTable(&TestA{}, &TestB{})
	//分页查询
	//连表查询
	var count int
	size := 2
	num := 1
	var bs []TestB
	db.Table("test_bs").Joins("right join test_as on test_bs.a_id = test_as.id").Count(&count)
	db.Table("test_bs").Joins("right join test_as on test_bs.a_id = test_as.id").Where("test_as.order_type = ?", 0).Limit(size).Offset((num - 1) * size).Find(&bs)
	fmt.Println(bs)
	//连表查，并将两个表拼接起来
	var abs []struct {
		Name string
		Age  string
		Addr string
	}
	//连表查的时候按条件查询时一定要指定表名，不然两个表有相同字段会报错 where 的时候条件可以为 ""
	db.Table("test_as").Joins("left join test_bs on test_as.id = test_bs.a_id").Select("test_as.name, test_as.age, test_bs.addr").
		Where("test_as.name = ?", "a1").Limit(size).Offset((num - 1) * size).Find(&abs)
	fmt.Println(abs)

	var as []*TestA
	f2(db, &TestA{}, &as, "name = ?", "a1")
	for _, v := range as {
		fmt.Println(v.Age)
	}
	fmt.Println("*********************************")
	var ta TestA
	db.Where("rflag = ?", true).First(&ta)
	fmt.Println(ta)*/

	//外键知识
	//db.CreateTable(&UserTest{}, &Profile{})

	// var profile Profile
	// db.Model(&UserTest{ProfileID: 3}).Related(&profile)
	// fmt.Println(profile)

	//排序分页
	//db.CreateTable(&TestD{})

	var td []TestD
	db.Order("order_time desc").Find(&td)
	fmt.Println("td*****************", td)
}

//TestA 主表
type TestA struct {
	ID        uint `gorm:"primary_key" json:"-"`
	Name      string
	Age       string
	OrderType int64
	Rflag 	  bool
}

//TestB 详情表
type TestB struct {
	ID        uint `gorm:"primary_key" json:"-"`
	AID       uint `gorm:"not null;unique"`
	Addr      string
	OrderType int64
}

func f1(db *gorm.DB, model interface{}, out interface{}, str string, args ...interface{}) {
	db.Where(str, args).Limit(1).Offset(0).Find(out)
}

func f2(db *gorm.DB, model interface{}, out interface{}, str string, args ...interface{}) {
	//注意这个地方需要...
	f1(db, model, out, str, args...)
}


// `User`属于`Profile`, `ProfileID`为外键
type UserTest struct {
	gorm.Model
	Profile   Profile
	ProfileID int
}
  
type Profile struct {
	gorm.Model
	Name string
}


type TestC struct {
	ID uint `gorm:"primary_key" json:"-"`
	Name string
	OrderTime int
}

type TestD struct {
	ID uint `gorm:"primary_key" json:"-"`
	Name string
	OrderTime string
}