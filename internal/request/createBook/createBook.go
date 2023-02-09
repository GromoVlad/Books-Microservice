package createBook

import (
	"database/sql"
	protobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/createBook/gRPC"
)

type DTO struct {
	Name        string         `form:"name"                   json:"name"                    binding:"required"`
	Category    string         `form:"category"               json:"category"                binding:"required"`
	AuthorId    int            `form:"author_id"              json:"author_id"               binding:"required,number"`
	Description sql.NullString `form:"description,omitempty"  json:"description,omitempty"   binding:"omitempty"`
}

func GetRequest(request *protobuf.Request) DTO {
	var dto DTO
	var isDescriptionValid bool

	if request.Description != nil {
		isDescriptionValid = true
	}

	dto.Name = request.Name
	dto.Category = request.Category
	dto.AuthorId = int(request.AuthorId)
	dto.Description = sql.NullString{String: dto.Description.String, Valid: isDescriptionValid}

	return dto
}
