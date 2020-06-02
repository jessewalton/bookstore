package main

import (
	"bookstore/controllers"
	"bookstore/db"
	"fmt"
	"net/http"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

var wsupgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func wshandler(w http.ResponseWriter, r *http.Request) {
	wsupgrader.CheckOrigin = func(r *http.Request) bool { return true }

	var conn *websocket.Conn
	var errWS error

	conn, errWS = wsupgrader.Upgrade(w, r, nil)
	if errWS != nil {
		fmt.Println("Failed to set websocket upgrade: %+v", errWS)
		return
	}

	for {
		t, msg, err := conn.ReadMessage()
		if err != nil {
			break
		}
		fmt.Println("received: ", string(msg))

		conn.WriteMessage(t, msg)
	}
}

func main() {
	r := gin.Default()

	db.ConnectDataBase()

	r.Use(cors.Default())

	r.GET("/ws", func(c *gin.Context) {
		wshandler(c.Writer, c.Request)
	})

	r.GET("/books", controllers.FindBooks)
	r.POST("/books", controllers.CreateBook)
	r.GET("/book/:id", controllers.FindBook)
	r.PATCH("/books/:id", controllers.UpdateBook)
	r.DELETE("/book/:id", controllers.DeleteBook)

	r.Run()
}
