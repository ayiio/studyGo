package db

import (
	"demo16_myblog/model"
	"testing"
	"time"
)

func init() {
	//parseTime=true 将mysql中的时间类型自动解析为go结构体中的时间类型
	dns := "xxx:xxx@tcp(localhost:3306)/xxx?parseTime=true"
	err := Init(dns)
	if err != nil {
		panic(err)
	}
}

func TestNewArticle(t *testing.T) {
	//构建对象
	article := &model.ArticleDetail{}
	article.ArticleInfo.CategoryId = 1
	article.ArticleInfo.CommentCount = 0
	article.Content = "test for new article"
	article.ArticleInfo.CreateTime = time.Now()
	article.ArticleInfo.Title = "test"
	article.ArticleInfo.UserName = "tester"
	article.ArticleInfo.Summary = "test for test"
	article.ArticleInfo.ViewCount = 1
	articleId, err := NewArticle(article)
	if err != nil {
		return
	}
	t.Logf("article id=%d\n", articleId)
}

func TestGetArticleList(t *testing.T) {
	articleList, err := GetArticleList(1, 15)
	if err != nil {
		return
	}
	for _, v := range articleList {
		t.Logf("article id = %d, article=%#v\n", v.Id, v)
	}
}

func TestGetArticleDetailByID(t *testing.T) {
	article, err := GetArticleDetailByID(1)
	if err != nil {
		return
	}
	t.Logf("article=%#v\n", article)
}

func TestGetArticleListByCategoryId(t *testing.T) {
	articleList, err := GetArticleListByCategoryId(1, 1, 15)
	if err != nil {
		return
	}
	for _, v := range articleList {
		t.Logf("article id = %d, article=%#v\n", v.Id, v)
	}
}
