package routes

import (
	"gorest/src/controller"

	"github.com/gin-gonic/gin"
)

func Api(router *gin.Engine) {
	router.GET("/api/books", controller.GetAllBooks)
	router.POST("/api/books", controller.CreateBook)
	router.GET("/api/books/:id", controller.GetBook)
	router.PUT("/api/books/:id", controller.UpdateBook)
	router.DELETE("/api/books/:id", controller.DeleteBook)
}
