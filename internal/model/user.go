package model

import (
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"time"
)

type User struct {
	UserId    uuid.UUID `json:"user_id"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Email     string    `json:"email"`
	Password  string    `json:"password,omitempty"`
	About     string    `json:"about,omitempty"`
	City      string    `json:"city,omitempty"`
	Country   string    `json:"country,omitempty"`
	Birthday  time.Time `json:"birthday"`
}

type UserList struct {
	TotalCount int     `json:"total_count"`
	TotalPages int     `json:"total_pages"`
	Page       int     `json:"page"`
	Size       int     `json:"size"`
	Users      []*User `json:"users"`
}
