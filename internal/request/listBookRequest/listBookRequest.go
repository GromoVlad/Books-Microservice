package listBookRequest

import (
	protobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/listBook/gRPC"
)

type DTO struct {
	Page     int    `form:"page,omitempty"       json:"page,omitempty"       binding:"omitempty,number"`
	Limit    int    `form:"limit,omitempty"      json:"limit,omitempty"      binding:"omitempty,number"`
	BookId   int    `form:"book_id,omitempty"    json:"book_id,omitempty"    binding:"omitempty,number"`
	Name     string `form:"name,omitempty"       json:"name,omitempty"       binding:"omitempty"`
	AuthorId int    `form:"author_id,omitempty"  json:"author_id,omitempty"  binding:"omitempty,number"`
	Category string `form:"category,omitempty"   json:"category,omitempty"   binding:"omitempty"`
	Offset   int
}

func GetRequest(request *protobuf.Request) DTO {
	var dto DTO

	dto.Page = int(request.Page.Value)
	dto.Limit = int(request.Limit.Value)
	dto.Offset = (dto.Page - 1) * dto.Limit
	if request.BookId != nil {
		dto.BookId = int(request.BookId.Value)
	}
	if request.Name != nil {
		dto.Name = request.Name.Value
	}
	if request.AuthorId != nil {
		dto.AuthorId = int(request.AuthorId.Value)
	}
	if request.Category != nil {
		dto.Category = request.Category.Value
	}

	return dto
}
