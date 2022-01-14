package routes

import (
	"github.com/gin-gonic/gin"
	"gorest/src/controller"
)

func Api() {
	router := gin.Default()

	router.GET("/api/books", controller.GetAllBooks)
	router.POST("/api/books", controller.CreateBook)
	router.GET("/api/books/:id", controller.GetBook)
	router.PUT("/api/books/:id", controller.UpdateBook)
	router.DELETE("/api/books/:id", controller.DeleteBook)

	if err := router.Run("0.0.0.0:80"); err != nil {
		panic(err)
	}
}
