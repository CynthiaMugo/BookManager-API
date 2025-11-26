package main

import (
	"net/http"
	"strconv"

	"github.com/CynthiaMugo/BookManager-API/books"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	list := LoadBooks()

	router.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, list)
	})

	router.POST("/books", func(c *gin.Context) {
		var input struct {
			Title  string `json:"title"`
			Author string `json:"author"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		list = books.AddBook(list, input.Title, input.Author)
		SaveBooks(list)
		c.JSON(http.StatusCreated, list)
	})

	router.PUT("/books/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}

		var input struct {
			Title  string `json:"title"`
			Author string `json:"author"`
		}
		if err := c.ShouldBindJSON(&input); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}

		updatedList, ok := books.UpdateBook(list, id, input.Title, input.Author)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		list = updatedList
		SaveBooks(list)
		c.JSON(http.StatusOK, list)
	})

	router.DELETE("/books/:id", func(c *gin.Context) {
		idParam := c.Param("id")
		id, err := strconv.Atoi(idParam)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid ID"})
			return
		}
		updatedList, ok := books.DeleteBook(list, id)
		if !ok {
			c.JSON(http.StatusNotFound, gin.H{"error": "Book not found"})
			return
		}
		list = updatedList
		SaveBooks(list)
		c.JSON(http.StatusOK, list)
	})

	router.Run(":8080")
}
