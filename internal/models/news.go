package models

import (
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"time"
)

type News struct {
	NewsId    uuid.UUID `json:"news_id"`
	AuthorId  uuid.UUID `json:"author_id"`
	Title     string    `json:"title"`
	Content   string    `json:"content"`
	Category  string    `json:"category"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type NewsList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	News       []*News `json:"news"`
}
