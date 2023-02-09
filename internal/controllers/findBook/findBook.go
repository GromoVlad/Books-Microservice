package findBook

import (
	"context"
	protobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/findBook/gRPC"
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func (s *findBookGRPC) FindBook(ctx context.Context, request *protobuf.Request) (*protobuf.Response, error) {
	book, err := bookRepository.FindOrFailBook(int(request.BookId))

	bookResponse := &protobuf.Response{
		Name:        book.Name,
		BookId:      int32(book.BookId),
		AuthorId:    int32(book.AuthorId),
		Category:    book.Category,
		Description: nil,
		CreatedAt:   nil,
		UpdatedAt:   nil,
	}

	if err != nil {
		bookResponse.ErrorMessage = err.Error()
		return bookResponse, nil
	}

	if book.Description.Valid == true {
		bookResponse.Description = &wrappers.StringValue{Value: book.Description.String}
	}
	if book.CreatedAt.Valid == true {
		bookResponse.CreatedAt = &wrappers.StringValue{Value: book.CreatedAt.Time.String()}
	}
	if book.UpdatedAt.Valid == true {
		bookResponse.UpdatedAt = &wrappers.StringValue{Value: book.UpdatedAt.Time.String()}
	}

	return bookResponse, nil
}

type findBookGRPC struct {
	protobuf.UnimplementedFindBookServer
	savedFeatures []*protobuf.Response // read-only after initialized
}

func NewServer() *findBookGRPC {
	return &findBookGRPC{}
}
