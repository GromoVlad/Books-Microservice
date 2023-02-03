-- +goose Up
-- +goose StatementBegin
CREATE TABLE books.books (
    book_id serial PRIMARY KEY,
    name VARCHAR(255) NOT NULL,
    author_id INTEGER NOT NULL,
    category VARCHAR(512) NOT NULL,
    description VARCHAR(4096) NULL,
    created_at TIMESTAMP NULL,
    updated_at TIMESTAMP NULL,
    CONSTRAINT name_author_id_unique
    UNIQUE (name, author_id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE books.books
-- +goose StatementEnd
