package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

//定义中间件
func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("中间件开始执行")
		c.Set("request", "中间件")
		//向下继续执行函数
		c.Next()
		s := c.Writer.Status()
		fmt.Println("中间件执行结束", s)
		t2 := time.Since(t)
		fmt.Println("中间件执行耗时", t2)
	}
}

func main() {
	r := gin.Default()

	//注册全局中间件
	r.Use(MiddleWare())

	{
		r.GET("/middle", func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request:", req)
			c.JSON(http.StatusOK, gin.H{"request": req})
		})

		//局部中间件MiddleWare
		r.GET("middle2", MiddleWare(), func(c *gin.Context) {
			req, _ := c.Get("request")
			fmt.Println("request2:", req)
			c.JSON(http.StatusOK, gin.H{"request2": req})
		})
	}

	r.Run(":8000")
}
