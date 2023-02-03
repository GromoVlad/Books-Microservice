-- +goose Up
-- +goose StatementBegin
CREATE SCHEMA books;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP SCHEMA books;
-- +goose StatementEnd
