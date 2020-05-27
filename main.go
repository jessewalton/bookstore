package main

import (
	"bookstore/controllers"
	"bookstore/db"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	db.ConnectDataBase()

	r.Use(cors.Default())

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)

	r.Run()
}
