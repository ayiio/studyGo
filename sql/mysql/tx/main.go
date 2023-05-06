package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

/*
事务：一个最小的不可再分的工作单元。
	通常一个事务对应一个完整的业务，例如银行账户转账，即最小的工作单元。
	同时这个完整的业务需要执行多次DML(insert/update/delete)语句共同联合完成

MySQL中的InnoDB支持事务，事务处理可以维护数据库的完整性，保证批量SQL要么全部执行，要么都不执行

事务的ACID: 原子性/不可分割性、一致性、隔离性/独立性、持久性
原子性：一个事务中的所有操作，要么全部完成，要么全部不完成，不会结束在中间某个环节
	  事务在执行过程中发生了错误，会回滚到事务开始前的状态
一致性：事务开始和结束后，数据库的完整性没有被破坏，表明写入的数据必须完全符合所有的预设规则
	  规则包括数据的精确度、串联性和后续数据库可以自发完成预定工作
隔离性：数据库允许多个并发事务同时对其数据进行读写和修改行为，隔离性可以防止多个事务并发执行时由于交叉执行而导致数据不一致
	  隔离级别：读未提交、读已提交、可重复读、串行化
持久性：事务处理结束后，对数据的修改是永久的，系统故障也不会造成数据丢失

相关方法：
开启事务： func (db *DB) Begin() (*Tx, error)
提交事务： func (tx *Tx) Commit() error
回滚事务： func (tx *Tx) Rollback() error
*/

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

func txDemo(id1, id2 int, trans_age int) {
	tx, err := db.Begin() // 开启事务
	if err != nil {
		if tx != nil {
			tx.Rollback() // 回滚
		}
		fmt.Printf("begin trans failed, err=%v\n", err)
		return
	}

	sqlStr := `update user set age=age-? where id = ?`
	_, err = tx.Exec(sqlStr, trans_age, id1)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql failed, err=%v\n", err)
		return
	}
	sqlStr2 := `update user set age=age+? where id = ?`
	_, err = tx.Exec(sqlStr2, trans_age, id2)
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("exec sql2 failed, err=%v\n", err)
		return
	}
	err = tx.Commit() // 提交事务
	if err != nil {
		tx.Rollback() // 回滚
		fmt.Printf("commit failed, err=%v\n", err)
		return
	}
	fmt.Println("exec trans success.")
}

func main() {
	err := initDB()
	if err != nil {
		fmt.Printf("initdb failed, err=%v\n", err)
		return
	}
	fmt.Println("befor tx")
	preQuery(0)
	txDemo(1, 2, 2)
	fmt.Println("after tx")
	preQuery(0)
}
