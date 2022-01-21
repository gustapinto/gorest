package main

import (
	"gorest/src/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	routes.Api(router)

	if err := router.Run("0.0.0.0:80"); err != nil {
		panic(err)
	}
}
