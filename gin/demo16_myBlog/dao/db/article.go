package db

import (
	"demo16_myblog/model"

	_ "github.com/go-sql-driver/mysql"
)

//插入文章
func NewArticle(article *model.ArticleDetail) (articleId int64, err error) {
	//加验证
	if article == nil {
		return
	}
	sqlStr := `insert into 
	 			 article(
					content
				   ,summary
				   ,title
				   ,username
				   ,category_id
				   ,view_count
				   ,comment_count
				 ) values(
					?
				   ,?
				   ,?
				   ,?
				   ,?
				   ,?
				   ,?
				 )`
	res, err := DB.Exec(sqlStr, article.Content, article.Summary, article.Title,
		article.Title, article.UserName, article.CategoryId,
		article.ViewCount, article.CommentCount)
	if err != nil {
		return
	}
	articleId, err = res.LastInsertId()
	return
}

//获取文章列表，分页
func GetArticleList(pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	//加验证
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	//时间降序
	sqlStr := `select id, summary, title, view_count, create_time, comment_count, username, category_id
	    	     from article
				where status = 1
				order by create_time desc
				limit ?, ?`
	err = DB.Select(&articleList, sqlStr, pageNum, pageSize)
	return
}

//根据文章ID查询单个文章
func GetArticleDetailByID(articleId int64) (articleDetail *model.ArticleDetail, err error) {
	if articleId < 0 {
		return
	}
	sqlStr := `select id, summary, title,content, view_count, create_time, comment_count, username, category_id
	             from article
				 where id = ?
				   and status = 1`
	err = DB.Get(&articleDetail, sqlStr, articleId)
	return
}

//根据分类ID查询一类文章
func GetArticleListByCategoryId(categoryId, pageNum, pageSize int) (articleList []*model.ArticleInfo, err error) {
	if pageNum < 0 || pageSize <= 0 {
		return
	}
	sqlStr := `select id, summary, title,content, view_count, create_time, comment_count, username, category_id
				 from article
				where category_id = ?
	  			  and status = 1
			   order by create_time desc
			   limit ?, ?`
	err = DB.Select(&articleList, sqlStr, categoryId, pageNum, pageSize)
	return
}
