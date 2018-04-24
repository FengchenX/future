package data 

import (
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
func(d *Data) SearchAll(x interface{}, query string, args ...interface{}) {
	rows,err := d.Query(query,args)	
	if err != nil {
		log.Fatal(err)
	}
	p := x.([]interface{})

	for rows.Next() {
		var dest []interface{}
		rows.Scan(dest...)
		p=append(p,dest)
	}
}
