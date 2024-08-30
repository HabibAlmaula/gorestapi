package user

import (
	"context"
	"database/sql"
	"learning/restapi/model/domain"
)

type UserRepository interface {
	Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User
	Delete(ctx context.Context, tx *sql.Tx, id int)
	GetById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error)
	GetByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.User
}
