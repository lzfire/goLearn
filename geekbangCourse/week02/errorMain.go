package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root:lzabcxxx@(127.0.0.1:3306)/lzsql")
	if err != nil {
		fmt.Println("open db failed")
		return
	}
	defer db.Close()
	err = db.Ping()
	if err != nil {
		fmt.Println("db connection failed", err)
		return
	}
	fmt.Println("db connection success")
	//操作一：执行数据操作语句
	/*
		sql := "insert into stu_db values (2,'berry','beijing')"
		result, _ := db.Exec(sql)     //执行SQL语句
		n, _ := result.RowsAffected() //获取受影响的记录数
		fmt.Println("受影响的记录数是", n)
	*/

	//操作二：执行预处理
	/*
		stu := [2][2]string{{"ketty", "shanghai"}, {"rose", "changsha"}}
		stmt, _ := db.Prepare("insert into stu_db values (?,?,?)") //获取预处理语句对象
		for idx, s := range stu {
			stmt.Exec(idx, s[0], s[1]) //调用预处理语句
		}
	*/

	//操作三：单行查询
	/*
		var id int
		var name, address string
		rows := db.QueryRow("select * from stu_db where id=2") //获取一行数据
		rows.Scan(&id, &name, &address)                        //将rows中的数据存到id,name中
		fmt.Println(id, "-", name, "-", address)
	*/

	//操作四：多行查询
	rows, err := db.Query("select * from stu_db;") //获取所有数据
	if err != nil {
		fmt.Println("db query nothing", err)
		return
	}
	var id int
	var address, name string
	for rows.Next() { //循环显示所有的数据
		rows.Scan(&id, &name, &address)
		fmt.Println(id, "-", name, "-", address)
	}
}
