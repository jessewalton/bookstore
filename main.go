package main

import (
	// "bookstore/controllers"
	"bookstore/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDataBase()

	r.Use(cors.Default())

	hub := newHub()
	go hub.run()

	r.GET("/ws", func(c *gin.Context) {
		serveWs(hub, c.Writer, c.Request)
	})

	r.GET("/books", FindBooks)
	r.POST("/books", func(c *gin.Context) {
		CreateBook(c, hub)
	})
	r.GET("/book/:id", FindBook)
	r.PATCH("/books/:id", UpdateBook)
	r.DELETE("/book/:id", DeleteBook)

	r.Run()
}
