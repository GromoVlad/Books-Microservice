package books

import (
	"database/sql"
	_ "github.com/lib/pq"
)

type Book struct {
	BookId      int            `db:"book_id"     json:"book_id"     example:"42"                  format:"int"`
	Name        string         `db:"name"        json:"name"        example:"Мрак твоих глаз"     format:"string"`
	AuthorId    int            `db:"author_id"   json:"author_id"   example:"42"                  format:"int"`
	Category    string         `db:"category"    json:"category"    example:"Некрореализм"        format:"string"`
	Description sql.NullString `db:"description" json:"description" example:"Описание"            format:"string" swaggertype:"string"`
	CreatedAt   sql.NullTime   `db:"created_at"  json:"created_at"  example:"2022-01-01 00:00:00" format:"string" swaggertype:"string"`
	UpdatedAt   sql.NullTime   `db:"updated_at"  json:"updated_at"  example:"2022-01-01 00:00:00" format:"string" swaggertype:"string"`
}
