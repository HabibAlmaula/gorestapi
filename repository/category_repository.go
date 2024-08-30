package repository

import (
	"context"
	"database/sql"
	"learning/restapi/model/domain"
)

type CategoryRepository interface {
	Create(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category
	Delete(ctx context.Context, tx *sql.Tx, id int)
	GetById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error)
	GetAll(ctx context.Context, tx *sql.Tx) []domain.Category
}
