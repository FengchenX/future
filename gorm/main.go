package main

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
)



func main() {

	db, err := gorm.Open("mysql", "launch:root@tcp(192.168.83.79:3306)/mytest?charset=utf8&parseTime=true&loc=Local")
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
	// db.AutoMigrate(&User{}, &Profile{})
	// profiles := []Profile {
	// 	{
	// 		Name: "uio",
	// 		UserID: 1,
	// 	},
	// 	{
	// 		Name: "qwe",
	// 		UserID: 1,
	// 	},
	// }
	// user := User{
	// 	//Refer: 1,
	// 	Profiles: profiles,
	// }
	// db.Create(&user)
	// user := User{}
	// db.Create(&user)
	// profiles := []Profile {
	// 	{
	// 		Name: "wer",
	// 		UserID: user.ID,
	// 	},
	// 	{
	// 		Name: "tyuu",
	// 		UserID: user.ID,
	// 	},
	// }
	// for _, v := range profiles {
	// 	db.Create(&v)
	// 
	
	//排序分页
	//db.CreateTable(&TestD{})

	// var td []TestD
	// db.Order("order_time desc").Find(&td)
	// fmt.Println("td*****************", td)

	//切片插入 不支持
	// var p = []Persons{
	// 	{
	// 		Name: "5555555555",
	// 		Age: 90,
	// 	},
	// 	{
	// 		Name: "5555555555",
	// 		Age: 90,
	// 	},
	// }
	// if mydb := db.Create(&p); mydb.Error != nil {
	// 	log.Fatal(mydb.Error)
	// }
	// var ps []Persons
	// db.Find(&ps)
	// tx := db.Begin()
	// for _, p := range ps {
	// 	if err := tx.Model(&p).Update("read", 1).Error; err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// tx.Commit()
	// db.AutoMigrate(&Persons{})

	//gorm select * 测试
	// var p []Per
	// mydb := db.Select("*").
	// 		Table("persons").
	// 		Find(&p)
	// if mydb.Error != nil {
	// 	log.Fatal(err)
	// }
	// fmt.Println(p)

	//gorm delete
	// db.AutoMigrate(&Persons{})
	// if mydb := db.Delete(Persons{}); mydb.Error != nil {
	// 	log.Fatal(mydb.Error)
	// }

	//gorm 增
	db.AutoMigrate(&TestA{})
	a := TestA {
		Name: "xya",
		Age: 12,
		OrderType: 0,
		Rflag: true,
	}
	if isNew := db.NewRecord(a); isNew {
		fmt.Println("是新纪录")
	} else {
		fmt.Println("不是新纪录")
	}
	if mydb := db.Create(&a); mydb.Error != nil {
		log.Fatal(mydb.Error)
	}
	if isNew := db.NewRecord(a); isNew {
		fmt.Println("是新纪录")
	} else {
		fmt.Println("不是新纪录")
	}
}



//TestA 主表
type TestA struct {
	ID        uint `gorm:"primary_key" json:"-"`
	Name      string
	Age       int
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


type Profile struct {
	gorm.Model
	//ID uint
	Name   string
	UserID uint
}
  
type User struct {
	gorm.Model
	//ID uint
	//Refer   uint
	Profiles []Profile `gorm:"ForeignKey:UserID;AssociationForeignKey:ID"`
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

type Persons struct {
	Id uint `gorm:"primary_key" json:"-"`
	Name string
	Age int
	Read bool
}

type Per struct {
	Persons
}