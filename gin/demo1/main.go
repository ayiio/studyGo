package main

import "github.com/gin-gonic/gin"

func main() {
	r := gin.Default()

	r.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "ok",
		})
	})

	r.GET("/test2/:name", func(ctx *gin.Context) {
		s := ctx.Param("name")
		ctx.String(http.StatusOK, s)
	})

	r.GET("/test3/:name/*action", func(ctx *gin.Context) {
		name := ctx.Param("name")
		action := ctx.Param("action")
		ctx.String(http.StatusOK, name+" = "+action)
	})

	r.Run(":0808")

}
