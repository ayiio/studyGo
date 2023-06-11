package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.POST("/form", func(c *gin.Context) {
		//表单参数设置默认值
		type1 := c.DefaultPostForm("type", "alter")
		//接受其他值
		username := c.PostForm("username")
		password := c.PostForm("password")
		//多选框
		hobbys := c.PostFormArray("hobby")
		c.String(http.StatusOK,
			fmt.Sprintf("type is %s, username is %s, password is %s, hobbys is %v",
				type1, username, password, hobbys))
	})

	r.Run(":8000")
}
