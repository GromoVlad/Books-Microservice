package createBook

import (
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
	createBookRequest "github.com/GromoVlad/go_microsrv_books/internal/request/createBook"
	"github.com/GromoVlad/go_microsrv_books/internal/response/createBook"
	"github.com/GromoVlad/go_microsrv_books/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Создать запись о книге
// CreateBook godoc
// @Summary      Создать запись о книге
// @Tags         Books
// @Produce      json
// @Param  		 RequestBody  body  createBook.DTO  true	"Тело запроса"
// @Success      201  {object}  createBook.Response
// @Router       /book [post]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	dto := createBookRequest.GetRequest(context)

	bookRepository.CreateBook(context, dto)

	result := createBook.Response{Success: true}
	context.StatusCreated(gin.H{"success": result.Success})
}
