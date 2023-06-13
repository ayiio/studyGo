package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginParm struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user"`
	Password string `form:"password" json:"password" uri:"password" xml:"password"`
}

func main() {
	r := gin.Default()

	r.POST("/loginForm", func(c *gin.Context) {
		var login LoginParm

		//Bind()默认解析并绑定form格式
		//根据请求头中的content-type自动推断
		if err := c.Bind(&login); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "root" || login.Password != "admin" {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.Run(":8000")
}
