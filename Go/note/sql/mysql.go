package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func main()  {
	//db, err := sql.Open("mysql", "root:root@/test?charset=utf8")
	db, err := sql.Open("mysql", "root:root@tcp(127.0.0.1:3306)/test?charset=utf8")
	checkErr(err)
	//插入数据-----------------------------------------------------------
/*	stmt, err := db.Prepare("INSERT userinfo SET username=?,departname=?,created=?")
	checkErr(err)

	res, err := stmt.Exec("astaxie", "研发部门", "2017-12-03")
	checkErr(err)

	id, err := res.LastInsertId()
	checkErr(err)

	fmt.Println(id)

	//更新数据-------------------------------------------------------------------
	stmt, err = db.Prepare("update userinfo set username=? where username=?")
	checkErr(err)

	res, err = stmt.Exec("astaxieupdate", "astaxie")
	checkErr(err)

	affect, err := res.RowsAffected()
	checkErr(err)

	fmt.Println(affect)*/

	//查询数据
	rows, err := db.Query("SELECT * FROM userinfo where uid=?",3)
	checkErr(err)

	for rows.Next() {
		var uid int
		var username string
		var department string
		var created string
		err := rows.Scan(&uid, &username, &department, &created)
		checkErr(err)
		fmt.Println(uid, " ", username, " ", department, " ", created)
	}

	db.Close()
}

func checkErr(err error)  {
	if err != nil {
		panic(err)
	}
}
