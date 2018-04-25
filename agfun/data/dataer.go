package data 

import (
	"reflect"
	"log"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

type Data struct {
	*sql.DB
	tp string
	key string
}

func NewData(tp, key string) *Data {
	var d = Data{tp: tp, key: key}	
	db,err := sql.Open(d.tp,d.key)
	if err != nil {
		log.Fatal(err)
	}
	d.DB=db
	return &d
}

func(d *Data) Close() {
	d.Close()
}

//insert
func(d *Data) Insert(query string, args ...interface{}) {
	insert,err := d.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer insert.Close()
	tx, err := d.Begin()
	if err != nil {
		log.Fatal(err)
	}
	tx.Stmt(insert).Exec(args)
	tx.Commit()
}

//update
func(d *Data) Update(query string, args ...interface{}) {
	update,err := d.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer update.Close()
	tx, err := d.Begin()
	if err != nil {
		log.Fatal(err)
	}
	tx.Stmt(update).Exec(args)
	tx.Commit()
}

//delete
func(d *Data) Delete(query string, args ...interface{}) {
	delete, err := d.Prepare(query)
	if err != nil {
		log.Fatal(err)
	}
	defer delete.Close()
	tx, err := d.Begin()
	if err != nil {
		log.Fatal(err)
	}
	tx.Stmt(delete).Exec(args)
	tx.Commit()
}

//search
func(d *Data) SearchAll(x interface{}, query string) {
	rows,err := d.Query(query)	
	if err != nil {
		log.Fatal(err)
	}

	typ := reflect.TypeOf(x)
	val := reflect.ValueOf(x)

	strtyp := typ.Elem().Elem()
	for rows.Next() {
		oneptr := reflect.New(strtyp)
		var pp  []interface{}
		
		//下面这一部分很重要不然不能从数据库中取得出数据
		for i:=0;i<oneptr.Elem().NumField();i++ {
			switch oneptr.Elem().Field(i).Kind() {
			case reflect.Int:
				var a int
				pp = append(pp,&a)
			case reflect.String:
				var a string
				pp = append(pp,&a)
			case reflect.Bool:
				var a bool
				pp = append(pp,&a)
			}
		}
		rows.Scan(pp...)

		for i:=0;i<len(pp);i++ {
			oneptr.Elem().Field(i).Set(reflect.ValueOf(pp[i]).Elem())
		}
		temp := reflect.Append(val.Elem(),oneptr.Elem())
		val.Elem().Set(temp)
	}
}
