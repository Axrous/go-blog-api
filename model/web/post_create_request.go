package web

type PostCreateRequest struct {
	Title     string
	Content   string
	AuthorId  int
	CreatedAt string
}