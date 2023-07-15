package main

import (
	"fmt"
	"demo14_booklist/book"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	r.LoadHTMLGlob("book/templates/*")

	r.GET("/book/list", func(c *gin.Context) {
		data := book.ListBookFunc()
		c.HTML(http.StatusOK, "book_list.tmpl", gin.H{"data": data})
	})

	r.GET("/book/new", func(c *gin.Context) {
		c.HTML(http.StatusOK, "new_book.tmpl", nil)
	})

	r.POST("/book/new", func(c *gin.Context) {
		var b book.Book
		c.Bind(&b)
		b.NewBookFunc()
		c.Redirect(302, "/book/list")
	})

	r.GET("/book/delete", func(c *gin.Context) {
		var b book.Book
		c.Bind(&b)
		b.DelBookFunc()
		c.Redirect(302, "list")
	})

	r.Run(":8008")
}
