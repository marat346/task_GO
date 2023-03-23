package post

type Post struct {
	Author  string
	Content string
}

func NewPost(author, content string) *Post {
	return &Post{
		Author:  author,
		Content: content,
	}
}
