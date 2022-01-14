package controller

import (
	"github.com/gin-gonic/gin"
	"gorest/src/repository"
	"net/http"
)

func CreateBook(ctx *gin.Context) {
	var newBook repository.Book

	if err := ctx.BindJSON(&newBook); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, err.Error())
	} else {
		repository.AddBook(&newBook)
		ctx.IndentedJSON(http.StatusCreated, newBook)
	}
}

func GetAllBooks(ctx *gin.Context) {
	books := repository.GetAllBooks()

	ctx.IndentedJSON(http.StatusOK, books)
}

func GetBook(ctx *gin.Context) {
	bookId := ctx.Param("id")

	if book, err := repository.GetBookById(bookId); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
	} else if book == nil {
		ctx.IndentedJSON(http.StatusNotFound, gin.H{"Message": "Book not found"})
	} else {
		ctx.IndentedJSON(http.StatusOK, book)
	}
}

func UpdateBook(ctx *gin.Context) {
	var book repository.Book

	bookId := ctx.Param("id")

	if err := ctx.BindJSON(&book); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
	} else {
		if err := repository.UpdateBook(bookId, &book); err != nil {
			ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
		}

		ctx.IndentedJSON(http.StatusOK, book)
	}
}

func DeleteBook(ctx *gin.Context) {
	bookId := ctx.Param("id")

	if err := repository.DeleteBook(bookId); err != nil {
		ctx.IndentedJSON(http.StatusInternalServerError, gin.H{"Message": err.Error()})
	} else {
		ctx.IndentedJSON(http.StatusNoContent, gin.H{})
	}
}
