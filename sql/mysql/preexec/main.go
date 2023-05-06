package main

/*
MySQL预处理

  普通sql处理过程：
    1.客户端对SQL语句进行占位符替换得到完整的SQL语句
	2.客户端发送完整的SQL语句到MySQL服务器
	3.MySQL服务器端执行完整的SQL语句并将结果返回给客户端

  预处理执行过程：
    1.把SQL语句分成两部分，命令部分和数据部分
	2.先把命令部分发送给MySQL服务器，MySQL服务器端进行SQL预处理
	3.数据部分随后发送给MySQL服务器，MySQL服务器端对SQL语句进行占位符替换
	4.MySQL服务器端执行完整的SQL语句并将结果返回给客户端

  预处理目的：
    1.优化MySQL服务器重复执行SQL的方法，可以提升服务器性能
	  提前让服务器编译，一次编译多次执行，节省后续编译成本
	2.避免SQL注入问题

*/

import (
	"database/sql"
	"fmt"
	"strconv"
)

var db *sql.DB

type user struct {
	id   int
	name string
	age  int
}

func initDB() (err error) {
	dsn := "用户名:密码@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn)
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	return
}

// 预处理查询
func preQuery(id int) {
	sqlStr := `select id, name, age from user where id > ?`
	sqlStmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err=%v\n", err)
		return
	}
	defer sqlStmt.Close()
	rows, err := sqlStmt.Query(id)
	if err != nil {
		fmt.Printf("query failed, err=%v\n", err)
		return
	}
	defer rows.Close()
	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan failed, err=%v\n", err)
			return
		}
		fmt.Printf("id=%d, name=%v, age=%d\n", u.id, u.name, u.age)
	}
}

// 预处理插入(同删除，更新，调用Exec)
func preInsert(name string, age int) {
	sqlStr := `insert into user(name, age) values(?, ?)`
	sqlStmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("prepare failed, err=%v\n", err)
		return
	}
	defer sqlStmt.Close()
	for i := 1; i < 4; i++ {
		_, err = sqlStmt.Exec(name+strconv.Itoa(i), age)
		if err != nil {
			fmt.Printf("insert failed, err=%v\n", err)
			return
		}
		fmt.Println("insert success.")
	}
}

func main() {
	initDB()
	preQuery(0)
	preInsert("ll", 22)
	preInsert("gg", 27)
	preQuery(0)
}
