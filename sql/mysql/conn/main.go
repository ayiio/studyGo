package main

/*
下载插件：go get -u github.com/go-sql-driver/mysql

*/

// 导入mysql包
import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	// 数据库信息, data source name
	dsn := "用户名:密码@tcp(127.0.0.1:3306)/test"
	// 连接数据库，Open不会校验用户名和密码，只会查看dsn格式是否正确
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("dsn:%s invalid, err:%v\n", dsn, err)
		return
	}
	defer db.Close()
	// Ping会校验用户名和密码，尝试连接数据库
	err = db.Ping()
	if err != nil {
		log.Fatalf("open %s failed, err:%v\n", dsn, err)
		return
	}
	fmt.Println("连接数据库成功")
}
