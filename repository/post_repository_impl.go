package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-blog-api/helper"
	"go-blog-api/model/domain"
)

type PostRepositoryImpl struct {
}

// Delete implements PostRepository.
func (*PostRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, postId int) {
	SQL := "delete from posts where id = ?"

	_, err := tx.ExecContext(ctx, SQL, postId)
	helper.PanicIfError(err)
}

// FindAll implements PostRepository.
func (*PostRepositoryImpl) FindAll(ctx context.Context, tx *sql.Tx) []domain.Post {
	SQL := "select id, title, content, created_at, author_id from posts"

	rows, err := tx.QueryContext(ctx, SQL)
	helper.PanicIfError(err)
	defer rows.Close()

	var posts []domain.Post
	for rows.Next() {
		post := domain.Post{}
		rows.Scan(&post.Id, &post.Title, &post.Content, &post.AuthorId, &post.CreatedAt)
		posts = append(posts, post)
	}

	return posts
}

// FindById implements PostRepository.
func (*PostRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error) {
	SQL := "select id, title, content, author_id, created_at from users where id = ?"

	rows, err := tx.QueryContext(ctx, SQL, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	post := domain.Post{}
	if rows.Next() {
		err := rows.Scan(&post.Id, &post.Title, &post.Content, &post.AuthorId, &post.CreatedAt)
		helper.PanicIfError(err)
		return post, nil
	}

	return post, errors.New("post is not found")
}

// Save implements PostRepository.
func (*PostRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "insert into posts(title, content, author_id, created_at) values(?, ?, ?, ?)"

	result, err := tx.ExecContext(ctx, SQL, post.Title, post.Content, post.AuthorId, post.CreatedAt)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)

	post.Id = int(id)
	return post
}

// Update implements PostRepository.
func (*PostRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post {
	SQL := "update posts set title = ?, content = ? where id = ?"

	_, err := tx.ExecContext(ctx, SQL, post.Title, post.Content, post.Id)
	helper.PanicIfError(err)

	return post
}

func NewPostRepository(PostRepositoryImpl) PostRepository {
	return &PostRepositoryImpl{}
}
