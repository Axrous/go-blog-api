package web

type CommentCreateRequest struct {
	Content string
	PostId  int
	UserId  int
}