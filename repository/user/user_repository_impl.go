package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"learning/restapi/helper"
	"learning/restapi/model/domain"
)

type UserRepositoryImpl struct {
}

func NewUserRepository() *UserRepositoryImpl {
	return &UserRepositoryImpl{}
}

func (u *UserRepositoryImpl) Create(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "INSERT INTO users (id, full_name, email, password) VALUES (?, ?, ?, ?)"
	id := uuid.New()
	fmt.Println("UUID: ", id)
	result, err := tx.ExecContext(ctx, SQL, id, user.FullName, user.Email, user.Password)
	//print error
	fmt.Println("Error: ", err)
	helper.PanicIfError(err)
	_, err = result.LastInsertId()
	helper.PanicIfError(err)
	user.Id = id.String()
	fmt.Println("User ID: ", user.Id)
	return user
}

func (u *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "UPDATE users SET full_name = ?, email = ? WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, user.FullName, user.Email, user.Id)
	helper.PanicIfError(err)
	return user
}

func (u *UserRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, id int) {
	SQL := "DELETE FROM users WHERE id = ?"
	_, err := tx.ExecContext(ctx, SQL, id)
	helper.PanicIfError(err)
}

func (u *UserRepositoryImpl) GetById(ctx context.Context, tx *sql.Tx, id int) (domain.User, error) {
	SQL := "SELECT id, full_name, email FROM users WHERE id = ?"
	rows, err := tx.QueryContext(ctx, SQL, id)
	helper.PanicIfError(err)
	user := domain.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.FullName, &user.Email)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (u *UserRepositoryImpl) GetByEmail(ctx context.Context, tx *sql.Tx, email string) (domain.User, error) {
	SQL := "SELECT id, full_name, email, password FROM users WHERE email = ?"
	rows, err := tx.QueryContext(ctx, SQL, email)
	helper.PanicIfError(err)
	user := domain.User{}
	if rows.Next() {
		err = rows.Scan(&user.Id, &user.FullName, &user.Email, &user.Password)
		helper.PanicIfError(err)
		return user, nil
	} else {
		return user, errors.New("user not found")
	}
}

func (u *UserRepositoryImpl) GetAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, full_name, email FROM users"
	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err = rows.Scan(&user.Id, &user.FullName, &user.Email)
		helper.PanicIfError(err)
		users = append(users, user)
	}
	return users
}
