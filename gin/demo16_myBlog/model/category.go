package model

//id, category_name, category_no, create_time, update_time

// 分类结构体
type Category struct {
	Id           int64  `db:"id"`
	CategoryName string `db:"category_name"`
	CategoryNo   int    `db:"category_no"`
}
