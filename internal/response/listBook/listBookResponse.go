package listBook

import (
	"github.com/GromoVlad/go_microsrv_books/internal/model/books"
)

type Response struct {
	Data    ListBookResponse `json:"data"`
	Success bool             `json:"success"`
}

type ListBookResponse struct {
	CurrentPage int          `json:"current_page"  example:"1"`
	Limit       int          `json:"limit"         example:"10"`
	Books       []books.Book `json:"books"`
}
