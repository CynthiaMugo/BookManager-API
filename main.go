package main

import (
	"net/http"
	"strconv"

	"github.com/CynthiaMugo/BookManager-API/books"
	"github.com/gin-gonic/gin"

	// Swagger imports
	_ "github.com/CynthiaMugo/BookManager-API/docs"
	ginSwagger "github.com/swaggo/gin-swagger"
	swaggerFiles "github.com/swaggo/files"
)

// @title Book Manager API
// @version 1.0
// @description This is a simple Book Manager API built with Go + Gin.
// @host localhost:8080
// @BasePath /
func main() {
	router := gin.Default()

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	list := LoadBooks()

	// ---------------- GET ----------------
	// @Summary Get all books
	// @Description Returns the full list of books
	// @Produce json
	// @Success 200 {array} books.Book
	// @Router /books [get]
	router.GET("/books", func(c *gin.Context) {
		c.JSON(http.StatusOK, list)
	})

	// ---------------- POST ----------------
	// @Summary Add a new book
	// @Description Add a new book with title and author
	// @Accept json
	// @Produce json
	// @Param book body map[string]string true "Book Data"
	// @Success 201 {array} books.Book
	// @Router /books [post]
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

	// ---------------- PUT ----------------
	// @Summary Update a book
	// @Description Update the title and author of a specific book
	// @Accept json
	// @Produce json
	// @Param id path int true "Book ID"
	// @Param book body map[string]string true "Updated Book"
	// @Success 200 {array} books.Book
	// @Failure 404 {object} map[string]string
	// @Router /books/{id} [put]
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

	// ---------------- DELETE ----------------
	// @Summary Delete a book
	// @Description Delete a book by ID
	// @Produce json
	// @Param id path int true "Book ID"
	// @Success 200 {array} books.Book
	// @Failure 404 {object} map[string]string
	// @Router /books/{id} [delete]
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
