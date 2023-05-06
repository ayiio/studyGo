package main

// 全局变量db
// 最大连接数
// QueryRow后调用Scan
// Query后调用Close
// Exec针对更新、插入、删除操作
// 删除后解决自增id断层的问题

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func initDB() (err error) {
	dsn := "用户名:密码@tcp(127.0.0.1:3306)/sql_test"
	db, err = sql.Open("mysql", dsn) // 全局变量db
	if err != nil {
		return
	}
	err = db.Ping()
	if err != nil {
		return
	}
	// 设置数据库连接池最大数(与数据库建立连接的最大数目)
	// 当<=0时，不会限制最大开启连接数，默认为0
	db.SetMaxOpenConns(4)

	// 设置连接池中最大闲置连接数
	// 当<=0时，不会保留闲置连接
	db.SetMaxIdleConns(2)
	return
}

// user 表对象
type user struct {
	id   int
	name string
	age  int
}

// 单行查询
// db.QueryRow(...) 执行一次查询，并期望返回最多一行结果(Row)
// QueryRow总是返回非nil的值，知道返回值的Scan方法被调用时，才会返回延迟的错误，如未找到结果
func queryOne(id int) {
	sqlStr := "select id, name, age from user where id=?"
	var lv user
	// 从连接池中拿一个连接去数据库查询单条记录
	// 确保QueryRow之后调用Scan方法，否则持有的数据库连接不会被释放
	err := db.QueryRow(sqlStr, id).Scan(&lv.id, &lv.name, &lv.age)
	if err != nil {
		fmt.Printf("scan failed, err=%v\n", err)
		return
	}
	fmt.Printf("queryOne: id:%v, name:%v, age:%v\n", lv.id, lv.name, lv.age)
}

// 单行查询，但不调用Scan方法，当超过最大连接池数时将一直处于等待状态
func queryForMaxConnTest() {
	sql_tst := `select id, name, age from user where id=?`
	for i := 1; i < 6; i++ {
		rowObj := db.QueryRow(sql_tst, 1)
		fmt.Printf("获取到第%v个连接,执行了queryrow得到%v对象\n", i, rowObj)
		// 当执行了4次QueryRow后一直等待
	}

}

// 多行查询
// db.Query(...) 执行一次查询，返回多行结果(Rows)
// 一般用于执行select命令，参数args表示query中的占位参数
func queryRows(id int) {
	sqlStr := `select id, name, age from user where id > ?`
	rows, err := db.Query(sqlStr, id)
	if err != nil {
		fmt.Printf("query failed, err=%v\n", err)
		return
	}
	// 关闭rows，释放所持有的数据库连接
	defer rows.Close()

	// 循环读取结果集中的数据
	for rows.Next() {
		var u user
		err = rows.Scan(&u.id, &u.name, &u.age)
		if err != nil {
			fmt.Printf("scan rows failed, err=%v\n", err)
			return
		}
		fmt.Printf("queryRows: id:%v, name:%v, age:%v\n", u.id, u.name, u.age)
	}
}

// 插入数据
// db.Exec(...)
// Exec执行一次命令(包括查询，删除，更新，插入等)，返回的ret是对已执行的SQL命令的总结
// 参数args表示query中的占位参数
func insertRow(name string, age int) (inID int64) {
	sqlStr := `insert into user(name, age) values (?, ?)`
	ret, err := db.Exec(sqlStr, name, age)
	if err != nil {
		fmt.Printf("insert failed, err=%v\n", err)
		return
	}
	inID, err = ret.LastInsertId() // 新插入数据的id
	if err != nil {
		fmt.Printf("get lastInsertId failed, err=%v\n", err)
		return
	}
	fmt.Printf("insert success, the id=%d\n", inID)
	return
}

// 更新数据
func updateRow(id, newAge int) {
	sqlStr := `update user set age=? where id=?`
	ret, err := db.Exec(sqlStr, newAge, id)
	if err != nil {
		fmt.Printf("update failed, err=%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get rowsAffected failed, err=%v\n", err)
		return
	}
	fmt.Printf("update success, affected rows:%d\n", n)
}

// 删除数据
// 解决id自增后断层问题：
/*
方法1：先删除id这个字段，再把id这个字段按建表时的要求添加到首位
alter table user drop id;
alter table user add id int(11) primary key auto_increment FIRST;
方法2：如果删除完还没有新增数据，即断层id还没有出现
alter table user auto_increment=1;
可以用该方法跳过某些编号，当=1时，ID默认从最大值加1开始自增
方法3：如果已经出现断层
set @auto_id=0;
update user set id=(@auto_id := @auto_id+1);
alter table user auto_increment=1;
*/
func deleteRow(id int) {
	sqlStr := `delete from user where id=?`
	ret, err := db.Exec(sqlStr, id)
	if err != nil {
		fmt.Printf("delete failed, err=%v\n", err)
		return
	}
	n, err := ret.RowsAffected() // 操作影响的行数
	if err != nil {
		fmt.Printf("get rowsAffected failed, err=%v\n", err)
		return
	}
	fmt.Printf("delete success, affected rows:%d\n", n)
}

func main() {
	err := initDB()
	if err != nil {
		log.Fatalf("open db failed, err=%v\n", err)
		return
	}
	defer db.Close()
	queryOne(1)
	// queryForMaxConnTest()
	queryRows(0)
	inID := insertRow("ww", 22)
	updateRow(int(inID), 25)
	deleteRow(int(inID))
	queryRows(0)
}
