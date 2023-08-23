package repository

import (
	"context"
	"database/sql"
	"go-blog-api/model/domain"
)

type CommentRepository interface {
	Save(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment
	FindByPostId(ctx context.Context, tx *sql.Tx, postId int) ([]domain.Comment, error) 
	FindById(ctx context.Context, tx *sql.Tx, commentId int) (domain.Comment, error)
	Update(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment
	Delete(ctx context.Context, tx *sql.Tx, commentId int)
}