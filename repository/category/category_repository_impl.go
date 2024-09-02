package category

import (
	"context"
	"database/sql"
	"errors"
	"learning/restapi/helper"
	"learning/restapi/model/domain"
)

type CategoryRepositoryImpl struct {
}

func NewCategoryRepository() CategoryRepository {
	return &CategoryRepositoryImpl{}
}

func (c *CategoryRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, category domain.Category, user domain.User) domain.Category {
	SQL := "INSERT INTO category (name, user_id) VALUES (?, ?)"
	result, err := tx.ExecContext(ctx, SQL, category.Name, user.Id)
	helper.PanicIfError(err)
	insertId, err := result.LastInsertId()
	helper.PanicIfError(err)
	category.Id = int(insertId)
	return category
}

func (c *CategoryRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, category domain.Category) domain.Category {
	SQL := "UPDATE category SET name = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, category.Name, category.Id)
	helper.PanicIfError(err)
	return category
}

func (c CategoryRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM category WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (c CategoryRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	defer rows.Close()
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}

func (c CategoryRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.Category {
	SQL := "SELECT id, name FROM category"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		helper.PanicIfError(err)
	}(rows)

	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}

func (c *CategoryRepositoryImpl) GetAllByUserId(ctx context.Context, tx *sql.Tx, userId string) []domain.Category {
	SQL := "SELECT id, name FROM category WHERE user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			helper.PanicIfError(err)
		}
	}(rows)
	var categories []domain.Category
	for rows.Next() {
		category := domain.Category{}
		err := rows.Scan(&category.Id, &category.Name, &category.UserId)
		helper.PanicIfError(err)
		categories = append(categories, category)
	}
	return categories
}

func (c *CategoryRepositoryImpl) GetByIdAndUserId(ctx context.Context, tx *sql.Tx, id int, userId string) (domain.Category, error) {
	SQL := "SELECT id, name FROM category WHERE id = ? AND user_id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id, userId)
	helper.PanicIfError(err)
	defer func(rows *sql.Rows) {
		err := rows.Close()
		if err != nil {
			helper.PanicIfError(err)
		}
	}(rows)
	category := domain.Category{}
	if rows.Next() {
		err := rows.Scan(&category.Id, &category.Name)
		helper.PanicIfError(err)
		return category, nil
	} else {
		return category, errors.New("category not found")
	}
}
