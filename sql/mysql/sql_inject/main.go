package main

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

/*任何时候都不应该自己拼接sql语句*/

var db *sqlx.DB

type user struct {
	ID   int
	Name string
	Age  int
}

func initDB() (err error) {
	dsn := "用户名:密码@tcp(127.0.0.1:3306)/sql_test"
	db, err = sqlx.Open("mysql", dsn)
	if err != nil {
		fmt.Printf("sqlx Open failed, err=%v\n", err)
		return
	}
	return
}

// SQL注入示例
func sqlInjectDemo(name string) {
	sqlStr := fmt.Sprintf("select id, name, age from user where name='%s'", name)
	fmt.Printf("SQL: %s\n", sqlStr)

	var users []user
	err := db.Select(&users, sqlStr)
	if err != nil {
		fmt.Printf("Select failed, err=%v\n", err)
		return
	}
	for _, u := range users {
		fmt.Printf("user: %#v\n", u)
	}
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initDB err=%v\n", err)
		return
	}
	defer db.Close()
	// 输入字符串将引发SQL注入问题
	// #注释了原拼接sql的最后一个'
	sqlInjectDemo("xxx' or 1=1#")                             // select id, name, age from user where name='xxx' or 1=1#'
	sqlInjectDemo("xxx' union select * from user #")          // select id, name, age from user where name='xxx' union select * from user #'
	sqlInjectDemo("xxx' or (select count(*) from user)<10 #") // 用户数量猜测：select id, name, age from user where name='xxx' or (select count(*) from user) < 10 #'
}
