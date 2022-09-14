package main

import (
	"API/configs"
	"errors"
	"github.com/gin-gonic/gin"
	"net/http"
)

type book struct {
	ID       string `json:"id"`
	Title    string `json:"title"`
	Author   string `json:"author"`
	Progress string `json:"progress"`
	Volume   string `json:"volume"`
}

var books = []book{
	{ID: "1", Title: "The Art of War", Author: "Sun Tzu", Progress: "Reading", Volume: "ND"},
	{ID: "2", Title: "Social Engineering - The Science of Human Hacking", Author: "Christopher Hadnagy", Progress: "Not started", Volume: "ND"},
	{ID: "3", Title: "Classroom of the Elite 2nd Year", Author: "Syougo Kinugasa", Progress: "Reading", Volume: "3"},
}

func getBooks(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, books)
}

func bookById(c *gin.Context) {
	id := c.Param("id")
	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	c.IndentedJSON(http.StatusOK, book)
}

func getBookByID(id string) (*book, error) {
	for i, b := range books {
		if b.ID == id {
			return &books[i], nil
		}
	}

	return nil, errors.New("book not found")
}

func addBook(c *gin.Context) {
	var newBook book

	if err := c.BindJSON(&newBook); err != nil {
		return
	}

	books = append(books, newBook)
	c.IndentedJSON(http.StatusCreated, newBook)
}

func startReading(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Progress == "Reading" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book is being read already"})
		return
	}

	// Starts reading the book
	book.Progress = "Reading"
	c.IndentedJSON(http.StatusOK, book)
}

func finishReading(c *gin.Context) {
	id, ok := c.GetQuery("id")

	if ok == false {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Missing id query parameter"})
	}

	book, err := getBookByID(id)

	if err != nil {
		c.IndentedJSON(http.StatusNotFound, gin.H{"message": "Book not found."})
		return
	}

	if book.Progress != "Reading" {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "Book was not being read"})
		return
	}

	// Finish reading the book
	book.Progress = "Finished"
	c.IndentedJSON(http.StatusOK, book)
}

func main() {
	router := gin.Default()

	router.GET("/books", getBooks) // Get all books
	//router.POST("/books", addBook)         // Add new book
	//router.GET("/books/:id", bookById)     // Search book by ID
	router.PATCH("/read", startReading)    // Start reading a book
	router.PATCH("/finish", finishReading) // Start reading a book

	configs.ConnectDB()

	router.Run("localhost:8080")
}
