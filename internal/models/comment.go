package models

import (
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"time"
)

type Comments struct {
	CommentId uuid.UUID `json:"comment_id"`
	AuthorId  uuid.UUID `json:"author_id"`
	NewsId    uuid.UUID `json:"news_id"`
	Message   string    `json:"message"`
	CreatedAt time.Time `json:"created_at"`
	UpdateAt  time.Time `json:"update_at"`
}

type CommentList struct {
	TotalCount int         `json:"total_count"`
	TotalPages int         `json:"total_pages"`
	Page       int         `json:"page"`
	Size       int         `json:"size"`
	Comments   []*Comments `json:"comments"`
}
