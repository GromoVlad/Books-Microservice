package findBook

import "github.com/GromoVlad/go_microsrv_books/internal/model/books"

type Response struct {
	Data    books.Book `json:"data"`
	Success bool       `json:"success"`
}
