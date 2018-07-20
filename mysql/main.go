package main

import (
	//"strings"
	//"time"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

/**
CREATE TABLE `userinfo` (
	`uid` INT(10) NOT NULL AUTO_INCREMENT,
	`username` VARCHAR(64) NULL DEFAULT NULL,
	`departname` VARCHAR(64) NULL DEFAULT NULL,
	`created` DATE NULL DEFAULT NULL,
	PRIMARY KEY (`uid`)
)
CREATE TABLE `userdetail` (
	`uid` INT(10) NOT NULL DEFAULT '0',
	`intro` TEXT NULL,
	`profile` TEXT NULL,
	PRIMARY KEY (`uid`)
)
--创建测试表
CREATE TABLE `timestampTest` (
  	`id` int(11) NOT NULL AUTO_INCREMENT,
   	`name` varchar(20) DEFAULT NULL,
  	`create_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
	`last_modify_time` timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
	PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8

10 --检测默认值，插入测试数据
11 INSERT INTO timestampTest (name) VALUES ('aa'),('bb'),('cc');
12 
13 --检测自动更新，更新某条数据
14 UPDATE timestampTest SET name = 'ab' WHERE id = 1;

*/

/*
select *
from expenses_bills
-- join user_bills on user_bills.bill_id = expenses_bills.id
-- where
-- order by column desc 
limit 2,1


select * 
from persons p
where p.name in (
select name
from persons
group by name 
having count(name)>1
)


//查询大于一次的数据
select * 
from persons 
group by name 
HAVING COUNT(name)>1;
*/
/**增删改查sql语句
insert into userinfo (uid,username,department,created) values(?,?,?,?)
delete from userinfo where uid = ?
update userinfo set username = ? where uid =?
select * from userinfo where uid = ?


select [column], [column]
form [table]
where [column] = ?
*/

func main() {

/**  普通增删改查
	db, err := sql.Open("mysql", "root:feng@/test?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res, err := stmt.Exec("astaxie", "研发部门", "2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("astaxieupdate", id)
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err = rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}

	//删除数据
	stmt, err = db.Prepare("delete from userinfo where uid=?")
	checkErr(err)
	res, err = stmt.Exec(id)
	checkErr(err)
	affect, err = res.RowsAffected()
	checkErr(err)
	fmt.Println(affect)
	db.Close()
	
*/

	/**更新时间示例
	update,err := db.Prepare("update userinfo set created=? where uid=?")
	checkErr(err)
	t:=time.Now()
	//ss:=strings.SplitAfter(t.String()," ")[0]

	_, err = update.Exec(t,1)
	checkErr(err)

	rows, err := db.Query("SELECT * FROM userinfo")
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err=rows.Scan(&uid,&username,&department,&created)
		checkErr(err)
		fmt.Println(uid)
		fmt.Println(username)
		fmt.Println(department)
		fmt.Println(created)
	}
	db.Close()
	*/

	/**
	   for {
	   	//准备一个删除状态
	   	del,err:=db.Prepare("truncate table userinfo");
	   	if err != nil {
	   		fmt.Println(err)
	   		return
	   	}
	   	//准备一个任务
	   	if tx, err := db.Begin(); err != nil {
	   		fmt.Println(err)
	   		del.Close()
	   	} else {
	   		tx.Stmt(del).Exec()
	   		insert,err := db.Prepare("insert into userinfo (username,departname,created) values(?,?,?)");
	   		if err != nil {
	   			fmt.Println(err)
	   			return
	   		}
	   		for _,row:=range []struct{
	   			username string
	   			department string
	   			created string}{
	   				{"feng","yafa","2014-5-6"},
	   				{"chen","faya","2018-2-3"},
	   				{"xx","fyd","2017-6-25"},
	   			} {
	   			tx.Stmt(insert).Exec(row.username,row.department,row.created)
	   		}
	   		insert.Close()

	   		tx.Commit()
	   		del.Close()
	   	}
	   }
	*/

	db, err := sql.Open("mysql", "root:root@tcp(39.108.80.66:3306)/finance?charset=utf8&parseTime=true&loc=Local")
	checkErr(err)

	stmt, err := db.Prepare("select * from users where user_name = adimn")
	checkErr(err)
	rows, err := stmt.Query()
	checkErr(err)
	for rows.Next() {
		// columns, err := rows.Columns()
		// checkErr(err)
		// fmt.Println(columns)
		var id uint
		var user_name string
		var password string
		var authority string
		rows.Scan(&id, &user_name, &password, &authority)
		fmt.Println(id, user_name, password, authority)
	}

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**连接数据库方式还有
user:password@tcp(localhost:5555)/dbname?charset=utf8
*/


/**标准库操作数据库实例
updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
...
tx, err := db.Begin()
...
res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)

*/
