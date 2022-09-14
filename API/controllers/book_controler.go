package controllers

import (
	"API/configs"
	"API/models"
	"API/responses"
	"context"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"net/http"
	"time"
)

var bookCollection *mongo.Collection = configs.GetCollection(configs.DB, "books")
var validate = validator.New()

func AddBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var book models.Book
		defer cancel()

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&book); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		newBook := models.Book{
			ID:       primitive.NewObjectID(),
			Title:    book.Title,
			Author:   book.Author,
			Progress: book.Progress,
			Volume:   book.Volume,
		}

		result, err := bookCollection.InsertOne(ctx, newBook)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusCreated, responses.BookResponse{Status: http.StatusCreated, Message: "success", Data: map[string]interface{}{"data": result}})
	}
}

func GetBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookID := c.Param("bookID")
		var book models.Book
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bookID)

		err := bookCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&book)
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": book}})

	}
}

func StartReading() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookID := c.Param("bookID")
		var book models.Book
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bookID)

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&book); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"title": book.Title, "author": book.Author, "progress": "Reading", "volume": book.Volume}
		result, err := bookCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		var updatedBook models.Book
		if result.MatchedCount == 1 {
			err := bookCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedBook)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedBook}})

	}
}

func FinishReading() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookID := c.Param("bookID")
		var book models.Book
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bookID)

		if err := c.BindJSON(&book); err != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if validationErr := validate.Struct(&book); validationErr != nil {
			c.JSON(http.StatusBadRequest, responses.BookResponse{Status: http.StatusBadRequest, Message: "error", Data: map[string]interface{}{"data": validationErr.Error()}})
			return
		}

		update := bson.M{"title": book.Title, "author": book.Author, "progress": "Finished", "volume": book.Volume}
		result, err := bookCollection.UpdateOne(ctx, bson.M{"id": objId}, bson.M{"$set": update})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}
		var updatedBook models.Book
		if result.MatchedCount == 1 {
			err := bookCollection.FindOne(ctx, bson.M{"id": objId}).Decode(&updatedBook)
			if err != nil {
				c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
				return
			}
		}
		c.JSON(http.StatusOK, responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": updatedBook}})

	}
}

func DeleteBook() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		bookID := c.Param("bookID")
		defer cancel()

		objId, _ := primitive.ObjectIDFromHex(bookID)

		result, err := bookCollection.DeleteOne(ctx, bson.M{"id": objId})
		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		if result.DeletedCount < 1 {
			c.JSON(http.StatusNotFound,
				responses.BookResponse{Status: http.StatusNotFound, Message: "error", Data: map[string]interface{}{"data": "User with specified ID not found!"}},
			)
			return
		}

		c.JSON(http.StatusOK,
			responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": "User successfully deleted!"}},
		)
	}
}

func GetAllBooks() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
		var books []models.Book
		defer cancel()

		results, err := bookCollection.Find(ctx, bson.M{})

		if err != nil {
			c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			return
		}

		//reading from the db in an optimal way
		defer results.Close(ctx)
		for results.Next(ctx) {
			var singleUser models.Book
			if err = results.Decode(&singleUser); err != nil {
				c.JSON(http.StatusInternalServerError, responses.BookResponse{Status: http.StatusInternalServerError, Message: "error", Data: map[string]interface{}{"data": err.Error()}})
			}

			books = append(books, singleUser)
		}

		c.JSON(http.StatusOK,
			responses.BookResponse{Status: http.StatusOK, Message: "success", Data: map[string]interface{}{"data": books}},
		)
	}
}
