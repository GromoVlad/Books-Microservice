package updateBook

import (
	"github.com/GromoVlad/go_microsrv_books/support/localContext"
	"strconv"
)

type DTO struct {
	Name        string `form:"name,omitempty"        json:"name,omitempty"          binding:"omitempty"`
	Category    string `form:"category,omitempty"    json:"category,omitempty"      binding:"omitempty"`
	AuthorId    int    `form:"author_id,omitempty"   json:"author_id,omitempty"     binding:"omitempty,number"`
	Description string `form:"description,omitempty" json:"description,omitempty"   binding:"omitempty"`
}

func GetRequest(context localContext.LocalContext) (DTO, int) {
	var dto DTO

	err := context.Context.ShouldBindJSON(&dto)
	context.BadRequestError(err)

	bookId, err := strconv.Atoi(context.Context.Param("bookId"))
	context.BadRequestError(err)

	return dto, bookId
}
