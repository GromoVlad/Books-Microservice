package bookRepository

import (
	"database/sql"
	"errors"
	"fmt"
	"github.com/GromoVlad/go_microsrv_books/internal/database/DB"
	"github.com/GromoVlad/go_microsrv_books/internal/model/books"
	"github.com/GromoVlad/go_microsrv_books/internal/request/createBook"
	"github.com/GromoVlad/go_microsrv_books/internal/request/listBookRequest"
	"github.com/GromoVlad/go_microsrv_books/internal/request/updateBook"
	"github.com/GromoVlad/go_microsrv_books/support/localContext"
	"strconv"
	"time"
)

func ListBooks(context localContext.LocalContext, dto listBookRequest.DTO) []books.Book {
	var books []books.Book
	var queryArgs []any
	var i int
	query := "SELECT * FROM books.books WHERE 1=1"

	if dto.BookId != 0 {
		i++
		query += " AND book_id = $" + strconv.Itoa(i) + " "
		queryArgs = append(queryArgs, dto.BookId)
	}
	if dto.AuthorId != 0 {
		i++
		query += " AND author_id = $" + strconv.Itoa(i) + " "
		queryArgs = append(queryArgs, dto.AuthorId)
	}
	if dto.Name != "" {
		i++
		query += " AND name like $" + strconv.Itoa(i) + " "
		queryArgs = append(queryArgs, "%"+dto.Name+"%")
	}
	if dto.Category != "" {
		i++
		query += " AND category = $" + strconv.Itoa(i)
		queryArgs = append(queryArgs, dto.Category)
	}

	limit := i + 1
	offset := limit + 1
	query += " LIMIT $" + strconv.Itoa(limit) + " OFFSET $" + strconv.Itoa(offset) + " ;"
	queryArgs = append(queryArgs, dto.Limit, dto.Offset)

	err := DB.Connect().Select(&books, query, queryArgs...)
	context.InternalServerError(err)

	return books
}

func FindOrFailBook(context localContext.LocalContext, bookId int) books.Book {
	var book books.Book
	DB.Connect().Get(&book, "SELECT * FROM books.books WHERE book_id = $1", bookId)

	if book.BookId == 0 {
		context.NotFoundError(
			errors.New(fmt.Sprintf("Книга с идентификатором %d не зарегистрирована в системе", bookId)),
		)
	}
	return book
}

func CreateBook(context localContext.LocalContext, dto createBook.DTO) {
	var book books.Book

	_ = DB.Connect().Get(
		&book,
		"SELECT book_id FROM books.books WHERE name = $1 AND author_id = $2",
		dto.Name,
		dto.AuthorId,
	)
	if book.BookId != 0 {
		context.AlreadyExistsError(errors.New("Книга с названием " + dto.Name + " уже зарегистрирована в системе"))
	}

	transaction := DB.Connect().MustBegin()
	_, err := transaction.NamedExec(
		"INSERT INTO books.books (name, author_id, category, description, created_at, updated_at) "+
			"VALUES (:name, :author_id, :category, :description, :created_at, :updated_at)",
		&books.Book{
			Name:        dto.Name,
			AuthorId:    dto.AuthorId,
			Category:    dto.Category,
			Description: sql.NullString{String: dto.Description, Valid: true},
			CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt:   sql.NullTime{},
		},
	)
	context.StatusConflictError(err)

	err = transaction.Commit()
	context.InternalServerError(err)
}

func UpdateBook(context localContext.LocalContext, dto updateBook.DTO, bookId int) books.Book {
	book := FindOrFailBook(context, bookId)
	mappingBook(&book, dto)

	transaction := DB.Connect().MustBegin()
	_, err := transaction.NamedExec(
		"UPDATE books.books SET updated_at = :updated_at, name = :name, category = :category, "+
			"author_id = :author_id, description = :description WHERE book_id = :book_id",
		&book,
	)
	context.StatusConflictError(err)

	err = transaction.Commit()
	context.InternalServerError(err)

	return book
}

func DeleteBook(context localContext.LocalContext, bookId int) {
	FindOrFailBook(context, bookId)

	transaction := DB.Connect().MustBegin()
	_, err := transaction.NamedExec("DELETE FROM books.books WHERE book_id = :book_id", &books.Book{BookId: bookId})
	context.StatusConflictError(err)

	err = transaction.Commit()
	context.InternalServerError(err)
}

func mappingBook(book *books.Book, dto updateBook.DTO) {
	book.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	if dto.Name != "" {
		book.Name = dto.Name
	}
	if dto.Category != "" {
		book.Category = dto.Category
	}
	if dto.AuthorId != 0 {
		book.AuthorId = dto.AuthorId
	}
	if dto.Description != "" {
		book.Description = sql.NullString{String: dto.Description, Valid: true}
	}
}
