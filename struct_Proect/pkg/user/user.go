package user

import "main/struct_Proect/pkg/post"

type User struct {
	Name    string
	Posts   []*post.Post
	Friends []*User
}

func NewUser(name string) *User {
	return &User{
		Name:    name,
		Posts:   make([]*post.Post, 0),
		Friends: make([]*User, 0),
	}
}

func (u *User) MakeFriend(userFriend *User) {
	u.Friends = append(u.Friends, userFriend)
}

func (u *User) AddPost(p *post.Post) {
	u.Posts = append(u.Posts, p)
}
