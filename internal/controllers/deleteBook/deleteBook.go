package deleteBook

import (
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
	"github.com/GromoVlad/go_microsrv_books/support/localContext"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Endpoint - Удалить запись о книге
// UpdateBook godoc
// @Summary      Удалить запись о книге
// @Tags         Books
// @Produce      json
// @Param        bookId  path  int  true  "Идентификатор книги"
// @Success      200  {object}  deleteBook.Response
// @Router       /book/{bookId} [delete]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}

	bookId, err := strconv.Atoi(context.Context.Param("bookId"))
	context.BadRequestError(err)
	bookRepository.DeleteBook(context, bookId)

	context.StatusCreated(gin.H{"success": true})
}
