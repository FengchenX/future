



package main
import (
	//"strings"
	//"time"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	
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

*/

func main() {
	db, err := sql.Open("mysql", "root:feng@/test?charset=utf8")
	checkErr(err)

	//插入数据
	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)
	res,err := stmt.Exec("astaxie", "研发部门","2012-12-09")
	checkErr(err)
	id, err := res.LastInsertId()
	checkErr(err)
	fmt.Println(id)

	//更新数据
	stmt, err = db.Prepare("update userinfo set username=? where uid=?")
	checkErr(err)
	res, err = stmt.Exec("astaxieupdate",id)
	checkErr(err)

	affect, err:=res.RowsAffected()
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
		err=rows.Scan(&uid,&username,&department,&created)
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

/*
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
}


func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

/**连接数据库方式还有
user:password@tcp(localhost:5555)/dbname?charset=utf8
*/

/**增删改查sql语句
insert into userinfo (uid,username,department,created) values(?,?,?,?)
delet from userinfo where uid = ?
update userinfo set username = ? where uid =?
select * from userinfo where uid = ?


select [column], [column]
form [table]
where [column] = ?
*/

/**标准库操作数据库实例
updateMoney, err := db.Prepare("UPDATE balance SET money=money+? WHERE id=?")
...
tx, err := db.Begin()
...
res, err := tx.Stmt(updateMoney).Exec(123.45, 98293203)

*/
