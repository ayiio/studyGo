package book

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func InitDB(database string) {
	if db == nil {
		dsn := "xxx:xxx@tcp(localhost:3306)/" + database
		var err error
		db, err = sql.Open("mysql", dsn)
		if err != nil {
			fmt.Printf("Open DB failed, err=%v\n", err)
			return
		}
	}
}

func SaveBook(b *Book) {
	sqlStr := "insert into book(title, price) values(?, ?)"
	sqlstmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare sql failed, err=%v\n", err)
		return
	}
	res, err := sqlstmt.Exec(b.Title, b.Price)
	if err != nil {
		fmt.Printf("Save new book failed, err=%v\n", err)
		return
	}
	insertID, _ := res.LastInsertId()
	fmt.Printf("Save book success, id=%v\n", insertID)
}

func DeleteBook(id int) {
	sqlStr := "delete from book where id=?"
	sqlstmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare sql failed, err=%v\n", err)
		return
	}
	res, err := sqlstmt.Exec(id)
	if err != nil {
		fmt.Printf("Save new book failed, err=%v\n", err)
		return
	}
	delID, _ := res.RowsAffected()
	fmt.Printf("delete book success, id=%v\n", delID)
}

func ListBook() (listBook []Book) {
	var b Book
	sqlStr := "select id, title, price from book;"
	sqlstmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Printf("Prepare sql failed, err=%v\n", err)
		return
	}
	rows, err := sqlstmt.Query()
	if err != nil {
		fmt.Printf("query books failed, err=%v\n", err)
		return
	}
	for rows.Next() {
		rows.Scan(&b.ID, &b.Title, &b.Price)
		listBook = append(listBook, b)
	}
	return
}
