package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	//异步
	r.GET("/request_async", func(c *gin.Context) {
		//需要使用拷贝的Context副本
		copyContext := c.Copy()

		go func() {
			time.Sleep(3 * time.Second)
			log.Println("异步请求" + copyContext.Request.URL.Path)
		}()
	})

	//同步
	r.GET("/request_sync", func(c *gin.Context) {
		time.Sleep(3 * time.Second)
		log.Println("同步请求" + c.Request.URL.Path)
	})

	r.Run(":8000")
}
