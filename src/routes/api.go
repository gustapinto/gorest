package routes

import (
	"github.com/gin-gonic/gin"
	"gorest/src/controller"
)

func Api() {
	router := gin.Default()

	router.POST("/api/books", controller.CreateBook)
	router.GET("/api/books", controller.GetAllBooks)

	router.Run("0.0.0.0:80")
}
