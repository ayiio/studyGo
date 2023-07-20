package model

import "time"

//id, category_id, content, title, view_count, comment_count, username, status, summary, create_time, update_time

//文章结构体
type ArticleInfo struct {
	Id           int64     `db:"id"`
	CategoryId   int64     `db:"category_id"`
	Summary      string    `db:"summary"`
	Title        string    `db:"title"`
	ViewCount    uint32    `db:"view_count"`
	CommentCount uint32    `db:"comment_count"`
	UserName     string    `db:"username"`
	CreateTime   time.Time `db:"create_time"`
}

//文章详情页结构体
type ArticleDetail struct {
	ArticleInfo
	//文章内容
	Content string `db:"content"`
	Category
}

//文章上下页
type ArticleRecord struct {
	ArticleInfo
	Category
}
