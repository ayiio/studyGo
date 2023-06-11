package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/upload", func(c *gin.Context) {
		//表单取文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)
		//保存到项目根目录
		c.SaveUploadedFile(file, file.Filename)
		//反馈信息
		c.String(http.StatusOK, fmt.Sprintf("%s upload!", file.Filename))
	})

	r.Run(":8001")
}
