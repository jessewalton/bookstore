package main

import (
	"encoding/json"
	"net/http"

	"bookstore/db"
	"bookstore/models"

	"github.com/gin-gonic/gin"
)

type CreateBookInput struct {
	Title  string `json:"title" binding:"required"`
	Author string `json:"author" binding:"required"`
	URL    string `json:"url"`
}

type UpdateBookInput struct {
	Title  string `json:"title"`
	Author string `json:"author"`
	URL    string `json:"url"`
}

func UpdateBook(c *gin.Context) {
	db := db.DB

	// get model if exist
	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	db.Model(&book).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": book})
}

// func CreateBook(c *gin.Context, ws websocket.Conn) {
func CreateBook(c *gin.Context, h *Hub) {
	db := db.DB

	// Validate input
	var input CreateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// Create book
	book := models.Book{Title: input.Title, Author: input.Author, URL: input.URL}
	db.Create(&book)

	// send http response
	c.JSON(http.StatusOK, gin.H{"data": book})

	// write to websocket
	bookJSON, err := json.Marshal(book)
	if err != nil {
		return
	}
	h.broadcast <- []byte(bookJSON)

}

func FindBook(c *gin.Context) {
	db := db.DB

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": book})
}

func FindBooks(c *gin.Context) {
	db := db.DB

	var books []models.Book
	db.Find(&books)

	c.JSON(http.StatusOK, gin.H{"data": books})
}

func DeleteBook(c *gin.Context) {
	db := db.DB

	var book models.Book
	if err := db.Where("id = ?", c.Param("id")).First(&book).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	db.Delete(&book)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
