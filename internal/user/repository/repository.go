package repository

import (
	"context"
	uuid "github.com/jackc/pgx/pgtype/ext/gofrs-uuid"
	"github.com/jvhab/Go-clean/internal/model"
)

type Repository interface {
	Register(context.Context, *model.User) (*model.User, error)
	Update(context.Context, *model.User) (*model.User, error)
	Delete(context.Context, uuid.UUID) error
	GetById(context.Context, uuid.UUID) (*model.User, error)
}
