package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//路由组1，处理get请求
	v1 := r.Group("/v1")
	{
		//curl http://localhost:8000/v1/login?name=zzz
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	//路由组2，处理Post请求
	v2 := r.Group("/v2")
	{
		//curl http://localhost:8000/v2/login -X POST
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}

	r.Run(":8000")
}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "xxx")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "yyy")
	c.String(http.StatusOK, fmt.Sprintf("hello %s", name))
}
