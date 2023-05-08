package main

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*
sqlx: 获取 go get github.com/jmoiron/sqlx
	一定程度上简化了sql Scan的操作
*/

var db *sqlx.DB

type user struct {
	ID   int
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "用户名:密码@tcp(127.0.0.1:3306)/sql_test"
	db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		return
	}
	db.SetMaxOpenConns(6)
	db.SetMaxIdleConns(4)
	return
}

// sqlx.Get 单行
// 可设置值(传入地址)，反射方式(字段可见)
func queryTest(id int) {
	sqlStr := `select id, name, age from user where id =?`
	var u user
	err := db.Get(&u, sqlStr, id)
	if err != nil {
		fmt.Printf("Get from table failed, err=%v\n", err)
		return
	}
	fmt.Printf("user.%d=%#v\n", id, u)
}

// sqlx.Select 多行
func queryManyTest(id int) {
	sqlStr := `select id, name, age from user where id > ?`
	var userList []user
	err := db.Select(&userList, sqlStr, id)
	if err != nil {
		fmt.Printf("Select many from table failed, err=%v\n", err)
		return
	}
	fmt.Printf("userliset: %#v\n", userList)
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("open db failed, err=%v\n", err)
		return
	}
	defer db.Close()
	queryTest(1)
	queryManyTest(0)
}
