package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

//统计程序运行耗时的中间件
func myTime(c *gin.Context) {
	start := time.Now()
	c.Next()
	t := time.Since(start)
	fmt.Println("程序执行耗时：", t)
}

func main() {
	r := gin.Default()

	r.Use(myTime)
	g := r.Group("/test")
	{
		g.GET("/since1", sinceProc1)
		g.GET("/since2", sinceProc2)
	}

	r.Run(":8000")
}

func sinceProc1(c *gin.Context) {
	time.Sleep(5 * time.Second)
}

func sinceProc2(c *gin.Context) {
	time.Sleep(3 * time.Second)
}
