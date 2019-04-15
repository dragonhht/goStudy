package main

import (
	_ "github.com/go-sql-driver/mysql"
	"database/sql"
	"fmt"
)

func chechErr(err error) {
	if err != nil {
		fmt.Println(err)
		panic(err)
	}
}

func main() {
	// 连接数据库
	// user:password@tcp(localhost:5555)/dbname?charset=utf8
	db, err := sql.Open("mysql", "root:123@tcp(localhost:3306)/gostudy?charset=utf8")
	chechErr(err)

	// 插入数据
	stmt, err := db.Prepare("insert student(name, age) value (?, ?)")
	chechErr(err)
	res, err := stmt.Exec("dragonhht", 18)
	chechErr(err)

	id, err := res.LastInsertId()
	chechErr(err)

	fmt.Println(id)

	// 更新数据
	stmt, err = db.Prepare("update student set age = ? where id = ?")
	chechErr(err)
	res, err = stmt.Exec(20, 1)
	chechErr(err)

	affect, err := res.RowsAffected()
	chechErr(err)

	fmt.Println(affect)

	// 查询数据
	rows, err := db.Query("select * from student")
	chechErr(err)

	for rows.Next() {
		var stuid int
		var stuname string
		var age int
		err = rows.Scan(&stuid, &stuname, &age)
		chechErr(err)
		fmt.Printf("student is [ %v, %v, %v ]\n", stuid, stuname, age)
	}

	// 删除
	stmt, err = db.Prepare("delete from student where id = ?")
	chechErr(err)
	res, err = stmt.Exec(3)

	affect, err = res.RowsAffected()
	chechErr(err)

	fmt.Println(affect)

	// 关闭连接
	db.Close()
}
