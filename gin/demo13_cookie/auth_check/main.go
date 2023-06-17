package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//中间件，检查是否有登录
func checkCookie() gin.HandlerFunc {
	return func(c *gin.Context) {
		if cookie, err := c.Cookie("root"); err == nil {
			if cookie == "admin" {
				c.Next()
				return
			}
		}
		//返回错误
		c.JSON(http.StatusUnauthorized, gin.H{"error": "StatusUnauthorized"})
		//不再调用后续代码
		c.Abort()
	}
}

func main() {
	r := gin.Default()

	r.GET("/login", func(c *gin.Context) {
		_, err := c.Cookie("root")
		if err != nil {
			//设置cookie
			c.SetCookie("root", "admin", 60, "/", "localhost", false, true)
		}
		c.String(http.StatusOK, "login successful")
	})

	r.GET("/home", checkCookie(), func(c *gin.Context) {
		cookie, _ := c.Cookie("root")
		c.JSON(http.StatusOK, gin.H{"cookie": cookie})
	})

	r.Run(":8000")
}
