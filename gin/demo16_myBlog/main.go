package main

import (
	"demo16_myblog/controller"
	"demo16_myblog/dao/db"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	dns := "xxx:xxx@tcp(localhost:3306)/xxx?parseTime=true"
	err := db.Init(dns)
	if err != nil {
		panic(err)
	}

	//静态文件
	r.Static("/static/", "./static")
	//加载模板
	r.LoadHTMLGlob("views/*")

	r.GET("/", controller.IndexHandle)
	r.GET("/category/", controller.CategoryList)

	r.Run(":8080")
}
