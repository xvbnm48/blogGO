package post

type PostRequest struct {
	Title   string `json:"title" required:"true"`
	Content string `json:"content" required:"true"`
}
