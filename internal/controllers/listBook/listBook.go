package listBook

import (
	"context"
	protobuf "github.com/GromoVlad/go_microsrv_books/internal/controllers/listBook/gRPC"
	"github.com/GromoVlad/go_microsrv_books/internal/model/books"
	"github.com/GromoVlad/go_microsrv_books/internal/repository/bookRepository"
	"github.com/GromoVlad/go_microsrv_books/internal/request/listBookRequest"
	"github.com/golang/protobuf/ptypes/wrappers"
)

func (s *listBookGRPC) ListBook(ctx context.Context, request *protobuf.Request) (*protobuf.Response, error) {
	dto := listBookRequest.GetRequest(request)

	response := &protobuf.Response{
		CurrentPage:  int32(dto.Page),
		Limit:        int32(dto.Limit),
		Books:        nil,
		ErrorMessage: "",
	}

	books, err := bookRepository.ListBooks(dto)

	if err != nil {
		response.ErrorMessage = err.Error()
		return response, err
	}

	protobufBooks := fillBooks(books)
	response.Books = protobufBooks

	return response, nil
}

func fillBooks(books []books.Book) []*protobuf.Book {
	var protobufBooks []*protobuf.Book

	for _, book := range books {
		protobufBook := protobuf.Book{
			Name:        book.Name,
			BookId:      int32(book.BookId),
			AuthorId:    int32(book.AuthorId),
			Category:    book.Category,
			Description: nil,
			CreatedAt:   nil,
			UpdatedAt:   nil,
		}
		if book.Description.Valid == true {
			protobufBook.Description = &wrappers.StringValue{Value: book.Description.String}
		}
		if book.CreatedAt.Valid == true {
			protobufBook.CreatedAt = &wrappers.StringValue{Value: book.CreatedAt.Time.String()}
		}
		if book.UpdatedAt.Valid == true {
			protobufBook.UpdatedAt = &wrappers.StringValue{Value: book.UpdatedAt.Time.String()}
		}
		protobufBooks = append(protobufBooks, &protobufBook)
	}

	return protobufBooks
}

type listBookGRPC struct {
	protobuf.UnimplementedListBookServer
	savedFeatures []*protobuf.Response // read-only after initialized
}

func NewServer() *listBookGRPC {
	return &listBookGRPC{}
}
