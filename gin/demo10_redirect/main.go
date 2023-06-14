package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("/redirect", func(c *gin.Context) {
		//支持内外部重定向
		c.Redirect(http.StatusMovedPermanently, "http://www.sogou.com/")
	})

	r.Run(":8000")
}
