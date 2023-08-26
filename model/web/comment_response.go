package web

type CommentResponse struct {
	Id      int
	Content string
	PostId  int
	User    interface{}
}