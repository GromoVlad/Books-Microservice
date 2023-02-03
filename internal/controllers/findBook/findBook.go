package findBook

import (
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
	"github.com/GromoVlad/go_microsrv_books/internal/response/findBook"
	"github.com/GromoVlad/go_microsrv_books/support/localContext"
	"github.com/gin-gonic/gin"
	"strconv"
)

// Endpoint - Найти книгу по идентификатору
// CreateBook godoc
// @Summary      Найти книгу
// @Tags         Books
// @Produce      json
// @Param        bookId  path  int  true  "Идентификатор пользователя"
// @Success      200  	 {object}  	findBook.Response
// @Router       /book/{bookId} [get]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	bookId, err := strconv.Atoi(ginContext.Param("bookId"))
	context.BadRequestError(err)

	book := bookRepository.FindOrFailBook(context, bookId)

	result := findBook.Response{Data: book, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}
