package repository

import (
	"context"
	"database/sql"
	"go-blog-api/model/domain"
)

type PostRepository interface {
	Save(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	FindAll(ctx context.Context, tx *sql.Tx) []domain.Post
	FindById(ctx context.Context, tx *sql.Tx, postId int) (domain.Post, error)
	Update(ctx context.Context, tx *sql.Tx, post domain.Post) domain.Post
	Delete(ctx context.Context, tx *sql.Tx, postId int)
}