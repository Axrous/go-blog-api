package repository

import (
	"context"
	"database/sql"
	"errors"
	"go-blog-api/helper"
	"go-blog-api/model/domain"
)

type CommentRepositoryImpl struct {
}

// Delete implements CommentRepository.
func (repository *CommentRepositoryImpl) Delete(ctx context.Context, tx *sql.Tx, commentId int) {
	SQL := "delete from comments where id = ?"

	_, err := tx.ExecContext(ctx, SQL, commentId)
	helper.PanicIfError(err)
}

// FindById implements CommentRepository.
func (repository *CommentRepositoryImpl) FindById(ctx context.Context, tx *sql.Tx, commentId int) (domain.Comment, error) {
	SQL := "select comments.id, comments.content, comments.post_id, comments.user_id, users.name from comments join users on comments.user_id = users.id where comments.id = ?"

	rows, err := tx.QueryContext(ctx, SQL, commentId)
	helper.PanicIfError(err)
	defer rows.Close()

	comment := domain.Comment{}
	if rows.Next() {
		rows.Scan(&comment.Id, &comment.Content, &comment.PostId, &comment.UserId, &comment.UserName)
		return comment, nil
	} 

	return comment, errors.New("comment not found")

}

// FindByPostId implements CommentRepository.
func (repository *CommentRepositoryImpl) FindByPostId(ctx context.Context, tx *sql.Tx, postId int) ([]domain.Comment, error) {
	SQL := "select comments.id, comments.content, comments.post_id, comments.user_id, users.name from comments join users on comments.user_id = users.id where comments.post_id = ?"

	rows, err := tx.QueryContext(ctx, SQL, postId)
	helper.PanicIfError(err)
	defer rows.Close()

	var comments []domain.Comment

		for rows.Next() {
			comment := domain.Comment{}
			rows.Scan(&comment.Id, &comment.Content, &comment.PostId, &comment.UserId, &comment.UserName)
			comments = append(comments, comment)
		}

		if len(comments) == 0 {
			return comments, errors.New("comments with post id not found")
		}

		return comments, nil
}

// Save implements CommentRepository.
func (repository *CommentRepositoryImpl) Save(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment {
	SQL := "insert into comments (content, post_id, user_id) values(?,?,?)"

	result, err := tx.ExecContext(ctx, SQL, comment.Content, comment.PostId, comment.UserId)
	helper.PanicIfError(err)

	id, err := result.LastInsertId()
	helper.PanicIfError(err)
	comment.Id = int(id)
	return comment
}

// Update implements CommentRepository.
func (repository *CommentRepositoryImpl) Update(ctx context.Context, tx *sql.Tx, comment domain.Comment) domain.Comment {
	SQL := "update comments set content = ? where id = ?"

	_, err := tx.ExecContext(ctx, SQL, comment.Content, comment.Id)
	helper.PanicIfError(err)

	return comment
}

func NewCommentRepository() CommentRepository {
	return &CommentRepositoryImpl{}
}
