package deleteBook

import (
	"context"
	protobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/deleteBook/gRPC"
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
)

func (s *deleteBookGRPC) DeleteBook(ctx context.Context, request *protobuf.Request) (*protobuf.Response, error) {
	bookId := int(request.BookId)
	deleted, err := bookRepository.DeleteBook(bookId)
	var message string
	if err != nil {
		message = err.Error()
	}

	bookResponse := &protobuf.Response{Success: deleted, ErrorMessage: message}
	return bookResponse, nil
}

type deleteBookGRPC struct {
	protobuf.UnimplementedDeleteBookServer
	savedFeatures []*protobuf.Response // read-only after initialized
}

func NewServer() *deleteBookGRPC {
	return &deleteBookGRPC{}
}
