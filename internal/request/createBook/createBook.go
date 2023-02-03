package createBook

import "github.com/GromoVlad/go_microsrv_books/support/localContext"

type DTO struct {
	Name        string `form:"name"                   json:"name"                    binding:"required"`
	Category    string `form:"category"               json:"category"                binding:"required"`
	AuthorId    int    `form:"author_id"              json:"author_id"               binding:"required,number"`
	Description string `form:"description,omitempty"  json:"description,omitempty"   binding:"omitempty"`
}

func GetRequest(context localContext.LocalContext) DTO {
	var dto DTO
	err := context.Context.ShouldBindJSON(&dto)
	context.BadRequestError(err)
	return dto
}
