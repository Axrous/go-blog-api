package domain

type Comment struct {
	Id      int
	Content string
	PostId  int
	UserId  int
}