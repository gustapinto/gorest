package controller

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gorest/src/repository"
	"net/http"
)

func CreateBook(ctx *gin.Context) {
	var newBook repository.Book

	if err := ctx.BindJSON(&newBook); err != nil {
		fmt.Println(err)
		return
	}

	repository.AddBook(&newBook)

	ctx.IndentedJSON(http.StatusCreated, newBook)
}

func GetAllBooks(ctx *gin.Context) {
	books := repository.GetAllBooks()

	ctx.IndentedJSON(http.StatusOK, books)
}
