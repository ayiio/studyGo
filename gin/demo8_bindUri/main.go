package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginForm struct {
	User     string `form:"username" json:"user" uri:"user" xml:"user"`
	Password string `form:"password" json:"password" uri:"password" xml:"password"`
}

func main() {
	r := gin.Default()

	//curl http://localhost:8000/root/admin
	r.GET("/uritest/:user/:password", func(ctx *gin.Context) {
		var login LoginForm
		if err := ctx.ShouldBindUri(&login); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if login.User != "root" || login.Password != "admin" {
			ctx.JSON(http.StatusBadRequest, gin.H{"status": "304"})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":8000")
}
