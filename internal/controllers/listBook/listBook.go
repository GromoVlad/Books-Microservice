package listBook

import (
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
	"github.com/GromoVlad/go_microsrv_books/internal/request/listBookRequest"
	"github.com/GromoVlad/go_microsrv_books/internal/response/listBook"
	"github.com/GromoVlad/go_microsrv_books/support/localContext"
	"github.com/gin-gonic/gin"
)

// Endpoint - Возвращает пагинированый список книг
// ListBook godoc
// @Summary      Возвращает пагинированый список книг
// @Tags         Books
// @Produce      json
// @Param  		 page   	query	int	 	false	"Номер страницы"  minimum(1)
// @Param  		 limit  	query	int	 	false	"Кол-во записей на странице" minimum(1)	maximum(20)
// @Param  		 book_id  	query	int		false	"Идентификатор книги"
// @Param  		 name  		query	string	false	"Поиск по названию книги"
// @Param  		 author_id  query	int		false	"Идентификатор автора"
// @Param  		 category  	query	string	false	"Категория"
// @Success      200  {object}  listBook.Response
// @Router       /book/list [get]
func Endpoint(ginContext *gin.Context) {
	context := localContext.LocalContext{Context: ginContext}
	dto := listBookRequest.GetRequest(context)

	books := bookRepository.ListBooks(context, dto)
	data := listBook.ListBookResponse{CurrentPage: dto.Page, Limit: dto.Limit, Books: books}

	result := listBook.Response{Data: data, Success: true}
	context.StatusOK(gin.H{"data": result.Data, "success": result.Success})
}
