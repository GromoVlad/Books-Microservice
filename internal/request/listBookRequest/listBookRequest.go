package listBookRequest

import "github.com/GromoVlad/go_microsrv_books/support/localContext"

type DTO struct {
	Page     int    `form:"page,omitempty"       json:"page,omitempty"       binding:"omitempty,number"`
	Limit    int    `form:"limit,omitempty"      json:"limit,omitempty"      binding:"omitempty,number"`
	BookId   int    `form:"book_id,omitempty"    json:"book_id,omitempty"    binding:"omitempty,number"`
	Name     string `form:"name,omitempty"       json:"name,omitempty"       binding:"omitempty"`
	AuthorId int    `form:"author_id,omitempty"  json:"author_id,omitempty"  binding:"omitempty,number"`
	Category string `form:"category,omitempty"   json:"category,omitempty"   binding:"omitempty"`
	Offset   int
}

func GetRequest(context localContext.LocalContext) DTO {
	var dto DTO

	err := context.Context.ShouldBindQuery(&dto)
	context.BadRequestError(err)

	if dto.Page == 0 {
		dto.Page = 1
	}
	if dto.Limit == 0 {
		dto.Limit = 10
	}
	dto.Offset = (dto.Page - 1) * dto.Limit

	return dto
}
