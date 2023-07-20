package db

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var (
	DB *sqlx.DB
)

//初始化
func Init(dns string) (err error) {
	DB, err = sqlx.Open("mysql", dns)
	if err != nil {
		return
	}
	//查看是否连接成功
	err = DB.Ping()
	if err != nil {
		return
	}

	//设置最大连接和最大空闲
	DB.SetMaxOpenConns(100)
	DB.SetMaxIdleConns(16)

	return
}
