package bookRepository

import (
	"database/sql"
	"fmt"
	"github.com/GromoVlad/go_microsrv_books/internal/database/DB"
	"github.com/GromoVlad/go_microsrv_books/internal/model/books"
	"github.com/GromoVlad/go_microsrv_books/internal/request/createBook"
	"github.com/GromoVlad/go_microsrv_books/internal/request/listBookRequest"
	"github.com/GromoVlad/go_microsrv_books/internal/request/updateBook"
	"github.com/GromoVlad/go_microsrv_books/support/logger"
	"strconv"
	"time"
)

func ListBooks(dto listBookRequest.DTO) ([]books.Book, error) {
	var books []books.Book
	var queryArgs []any
	var countArgs int
	query := "SELECT * FROM books.books WHERE 1=1"

	if dto.BookId != 0 {
		countArgs++
		query += " AND book_id = $" + strconv.Itoa(countArgs) + " "
		queryArgs = append(queryArgs, dto.BookId)
	}
	if dto.AuthorId != 0 {
		countArgs++
		query += " AND author_id = $" + strconv.Itoa(countArgs) + " "
		queryArgs = append(queryArgs, dto.AuthorId)
	}
	if dto.Name != "" {
		countArgs++
		query += " AND name like $" + strconv.Itoa(countArgs) + " "
		queryArgs = append(queryArgs, "%"+dto.Name+"%")
	}
	if dto.Category != "" {
		countArgs++
		query += " AND category = $" + strconv.Itoa(countArgs)
		queryArgs = append(queryArgs, dto.Category)
	}

	limit := countArgs + 1
	offset := limit + 1
	query += " LIMIT $" + strconv.Itoa(limit) + " OFFSET $" + strconv.Itoa(offset) + " ;"
	queryArgs = append(queryArgs, dto.Limit, dto.Offset)

	connect := DB.Connect()
	defer connect.Close()
	err := connect.Select(&books, query, queryArgs...)

	if err != nil {
		logger.ErrorLog("InternalServerError", err.Error())
		return books, err
	}

	return books, nil
}

func FindOrFailBook(bookId int) (books.Book, error) {
	connect := DB.Connect()
	defer connect.Close()

	var book books.Book
	connect.Get(&book, "SELECT * FROM books.books WHERE book_id = $1", bookId)

	if book.BookId == 0 {
		err := fmt.Errorf(fmt.Sprintf("Книга с идентификатором %d не зарегистрирована в системе", bookId))
		logger.ErrorLog("NotFound", err.Error())
		return book, err
	}

	return book, nil
}

func CreateBook(dto createBook.DTO) (bool, error) {
	var book books.Book
	connect := DB.Connect()
	defer connect.Close()

	_ = connect.Get(
		&book,
		"SELECT book_id FROM books.books WHERE name = $1 AND author_id = $2",
		dto.Name,
		dto.AuthorId,
	)
	if book.BookId != 0 {
		err := fmt.Errorf("Книга с названием " + dto.Name + " уже зарегистрирована в системе")
		logger.ErrorLog("NotFound", err.Error())
		return false, err
	}

	transaction := connect.MustBegin()
	_, err := transaction.NamedExec(
		"INSERT INTO books.books (name, author_id, category, description, created_at, updated_at) "+
			"VALUES (:name, :author_id, :category, :description, :created_at, :updated_at)",
		&books.Book{
			Name:        dto.Name,
			AuthorId:    dto.AuthorId,
			Category:    dto.Category,
			Description: dto.Description,
			CreatedAt:   sql.NullTime{Time: time.Now(), Valid: true},
			UpdatedAt:   sql.NullTime{},
		},
	)
	if err != nil {
		logger.ErrorLog("StatusConflictError", err.Error())
		return false, err
	}

	err = transaction.Commit()
	if err != nil {
		logger.ErrorLog("InternalServerError", err.Error())
		return false, err
	}

	return true, nil
}

func UpdateBook(dto updateBook.DTO, bookId int) (books.Book, error) {
	book, err := FindOrFailBook(bookId)
	if err != nil {
		return book, err
	}

	mappingBook(&book, dto)
	connect := DB.Connect()
	defer connect.Close()

	transaction := connect.MustBegin()
	_, err = transaction.NamedExec(
		"UPDATE books.books SET updated_at = :updated_at, name = :name, category = :category, "+
			"author_id = :author_id, description = :description WHERE book_id = :book_id",
		&book,
	)
	if err != nil {
		logger.ErrorLog("StatusConflictError", err.Error())
		return book, err
	}

	err = transaction.Commit()
	if err != nil {
		logger.ErrorLog("InternalServerError", err.Error())
		return book, err
	}

	return book, nil
}

func DeleteBook(bookId int) (bool, error) {
	_, err := FindOrFailBook(bookId)
	if err != nil {
		return false, err
	}

	connect := DB.Connect()
	defer connect.Close()

	transaction := connect.MustBegin()
	_, err = transaction.NamedExec("DELETE FROM books.books WHERE book_id = :book_id", &books.Book{BookId: bookId})
	if err != nil {
		logger.ErrorLog("StatusConflictError", err.Error())
		return false, err
	}

	err = transaction.Commit()
	if err != nil {
		logger.ErrorLog("InternalServerError", err.Error())
		return false, err
	}

	return true, nil
}

func mappingBook(book *books.Book, dto updateBook.DTO) {
	book.UpdatedAt = sql.NullTime{Time: time.Now(), Valid: true}
	book.Description = dto.Description
	if dto.Name != "" {
		book.Name = dto.Name
	}
	if dto.Category != "" {
		book.Category = dto.Category
	}
	if dto.AuthorId != 0 {
		book.AuthorId = dto.AuthorId
	}
}
