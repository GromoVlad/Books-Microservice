package updateBook

import (
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
	updateBookRequest "github.com/GromoVlad/go_microsrv_books/internal/request/updateBook"
	"github.com/GromoVlad/go_microsrv_books/internal/response/updateBook"
	"github.com/GromoVlad/go_microsrv_books/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Обновить запись о книге
// UpdateBook godoc
// @Summary      Обновить запись о книге
// @Tags         Books
// @Produce      json
// @Param        bookId  path  int  true  "Идентификатор книги"
// @Param  		 RequestBody  body  updateBook.DTO  true	"Тело запроса"
// @Success      200  {object}  updateBook.Response
// @Router       /book/{bookId} [put]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	dto, bookId := updateBookRequest.GetRequest(context)

	book := bookRepository.UpdateBook(context, dto, bookId)

	result := updateBook.Response{Data: book, Success: true}
	context.StatusCreated(gin.H{"data": result.Data, "success": result.Success})
}
