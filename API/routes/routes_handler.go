package routes

import (
	"API/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.GET("/books", controllers.GetAllBooks())            // Get all books
	router.POST("/books", controllers.AddBook())               // Add new book
	router.GET("/books/:bookID", controllers.GetBook())        // Get book by ID
	router.PUT("/read/:bookID", controllers.StartReading())    // Start reading a book
	router.PUT("/finish/:bookID", controllers.FinishReading()) // Finish reading a book
	router.DELETE("/books/:bookID", controllers.DeleteBook())  // Delete a book
}
