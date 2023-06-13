package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

//客户端传参，后端接收并解析到结构体

type Person struct {
	//binding 修饰为必选字段
	Name string `form:"username" json:"user" uri:"user" xml:"user" binding:"required"`
	Age  int    `form:"age" json:"age" uri:"age" xml:"age" binding:"required"`
}

func main() {
	r := gin.Default()

	//curl http://localhost:8000/add -H 'content-type:application/json' -d "{\"user\":\"root\", \"age\":1}"" -X POST
	r.POST("add", func(c *gin.Context) {
		//声明接收的结构体
		var p Person
		//将request中的数据按照json格式解析到结构体
		if err := c.ShouldBindJSON(&p); err != nil {
			//gin.H生成json数据
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if p.Name != "root" || p.Age != 1 {
			c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		c.JSON(http.StatusOK, gin.H{"status": "200"})
	})

	r.Run(":8000")
}
