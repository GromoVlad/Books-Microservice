package createBook

import (
	"context"
	protobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/createBook/gRPC"
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
	createBookRequest "github.com/GromoVlad/go_microsrv_books/internal/request/createBook"
)

func (s *createBookGRPC) CreateBook(ctx context.Context, request *protobuf.Request) (*protobuf.Response, error) {
	dto := createBookRequest.GetRequest(request)
	created, err := bookRepository.CreateBook(dto)
	var message string
	if err != nil {
		message = err.Error()
	}

	bookResponse := &protobuf.Response{Success: created, Message: message}
	return bookResponse, nil
}

type createBookGRPC struct {
	protobuf.UnimplementedCreateBookServer
	savedFeatures []*protobuf.Response // read-only after initialized
}

func NewServer() *createBookGRPC {
	return &createBookGRPC{}
}
