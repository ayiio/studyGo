package controller

import (
	"demo16_myblog/service"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

//访问主页的控制器
func IndexHandle(c *gin.Context) {
	//从service取数据
	//加载了文章数据
	articleRecordList, err := service.GetArticleRecordList(0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//加载分类数据
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}

//分类云
func CategoryList(c *gin.Context) {
	categoryIdStr := c.Query("category_id")
	//转int
	category_id, err := strconv.ParseInt(categoryIdStr, 10, 64)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//根据分类ID获取文章列表
	articleRecordList, err := service.GetArticleRecordListByID(int(category_id), 0, 15)
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	//再次加载分类数据，用于分类面板显示
	categoryList, err := service.GetAllCategoryList()
	if err != nil {
		c.HTML(http.StatusInternalServerError, "views/500.html", nil)
		return
	}
	c.HTML(http.StatusOK, "views/index.html", gin.H{
		"article_list":  articleRecordList,
		"category_list": categoryList,
	})
}
