package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//限定最大文件大小为8M, 默认32M
	r.MaxMultipartMemory = 8 << 20

	r.POST("/uploads", func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
		}
		//获取所有图片
		files := form.File["files"]
		//遍历所有图片
		for _, file := range files {
			if err := c.SaveUploadedFile(file, file.Filename); err != nil {
				c.String(http.StatusBadRequest, fmt.Sprintf("upload failed, err: %s", err.Error()))
				return
			}
		}
		c.String(http.StatusOK, fmt.Sprintf("upload success, %d files", len(files)))
	})

	r.Run(":8000")

}
