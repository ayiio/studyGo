package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//加载模板文件
	r.LoadHTMLGlob("templates/*")
	// r.LoadHTMLFiles("templates/index.tmpl")

	r.GET("/index", func(c *gin.Context) {
		//根据文件名进行渲染
		//json将替换模板中的title字段
		c.HTML(http.StatusOK, "index.impl", gin.H{"title": "自定义标题"})
	})

	r.Run(":8000")
}
