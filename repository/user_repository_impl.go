package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-blog-api/helper"
	"go-blog-api/model/domain"
)

type UserRepositoryImpl struct {
}

// FindForAuth implements UserRepository.
func (repository *UserRepositoryImpl) FindForAuth(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "select username, password from users where username = ?"

	rows := tx.QueryRowContext(ctx, SQL, user.Username)

	user = domain.User{}
	rows.Scan(&user.Username, &user.Password)
	return user
}

// FindAll implements UserRepository.
func (repository *UserRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.User {
	SQL := "SELECT id, username, name from users"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()
	
	var users []domain.User
	for rows.Next() {
		user := domain.User{}
		err := rows.Scan(&user.Id, &user.Username, &user.Name)
		helper.PanicIfError(err)
		users = append(users, user)
	}

	return users
}

// FindById implements UserRepository.
func (repository *UserRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, userId int) (domain.User, error) {
	SQL := "select id, username, name from users wehere id = ?"

	rows, err := tx.QueryContext(ctx, SQL, userId)
	helper.PanicIfError(err)
	defer rows.Close()

	user := domain.User{}
	if rows.Next() {
		err := rows.Scan(&user.Id, &user.Username, &user.Name)
		helper.PanicIfError(err)
		return user, nil
	}
	return user, errors.New("user not found")
}

// Save implements UserRepository.
func (repository *UserRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	SQL := "insert into users(username, name, password) values(?, ?, ?)"

	result, err := tx.ExecContext(ctx, SQL, user.Username, user.Name, user.Password)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	user.Id = int(id)
	return user

}

// Update implements UserRepository.
func (repository *UserRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, user domain.User) domain.User {
	panic("unimplemented")
}

func NewUserRepository() UserRepository {
	return &UserRepositoryImpl{}
}
