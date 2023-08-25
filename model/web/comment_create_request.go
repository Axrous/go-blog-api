package web

type CommentCreateRequest struct {
	Content string `json:"content"`
	PostId  int    `json:"postid"`
	UserId  int
}