package updateBook

import (
	"database/sql"
	protobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/updateBook/gRPC"
)

type DTO struct {
	Name        string
	Category    string
	AuthorId    int
	Description sql.NullString
}

func GetRequest(request *protobuf.Request) DTO {
	var dto DTO
	dto.Name = request.Name
	dto.Category = request.Category
	dto.AuthorId = int(request.AuthorId)
	dto.Description = sql.NullString{String: "", Valid: false}

	if request.Description != nil {
		dto.Description = sql.NullString{String: request.Description.Value, Valid: true}
	}

	return dto
}
