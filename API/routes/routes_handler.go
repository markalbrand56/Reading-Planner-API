package routes

import (
	"API/controllers"
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.POST("/books", controllers.AddBook())        // Add new book
	router.GET("/books/:bookID", controllers.GetBook()) // Get book by ID
}
