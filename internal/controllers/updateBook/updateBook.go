package updateBook

import (
	"context"
	protobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/updateBook/gRPC"
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
	updateBookRequest "github.com/GromoVlad/go_microsrv_books/internal/request/updateBook"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func (s *updateBookGRPC) UpdateBook(ctx context.Context, request *protobuf.Request) (*protobuf.Response, error) {
	dto := updateBookRequest.GetRequest(request)
	book, err := bookRepository.UpdateBook(dto, int(request.BookId))

	response := &protobuf.Response{
		Success:     true,
		Name:        book.Name,
		BookId:      int32(book.BookId),
		AuthorId:    int32(book.AuthorId),
		Category:    book.Category,
		Description: nil,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	if err != nil {
		response.ErrorMessage = err.Error()
		return response, nil
	}

	if book.Description.Valid == true {
		response.Description = &wrappers.StringValue{Value: book.Description.String}
	}
	if book.CreatedAt.Valid == true {
		response.CreatedAt = &wrappers.StringValue{Value: book.CreatedAt.Time.String()}
	}
	if book.UpdatedAt.Valid == true {
		response.UpdatedAt = &wrappers.StringValue{Value: book.UpdatedAt.Time.String()}
	}

	return response, nil
}

type updateBookGRPC struct {
	protobuf.UnimplementedUpdateBookServer
	savedFeatures []*protobuf.Response // read-only after initialized
}

func NewServer() *updateBookGRPC {
	return &updateBookGRPC{}
}
