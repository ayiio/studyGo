package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.GET("cookie", func(c *gin.Context) {
		cookie, err := c.Cookie("cookie_key")
		if err != nil {
			fmt.Println("cookie: notSet")
			c.SetCookie("cookie_key", "cookie_value",
				60, "/", "localhost", false, true)
		}
		fmt.Printf("cookie=%s\n", cookie)
	})
	r.Run(":8000")
}
