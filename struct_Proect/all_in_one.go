package main

import "fmt"

func main() {
	// создаем пользователя,прикрепляем к нему друзей,сохранить в хранилище
	u := newUser("Firdaus")
	rentik := newUser("Rentik")
	u.makeFriend(rentik)
	// сохраняем пользователя Рентик в storage
	us := newUserStorage()
	us.put(u)

	fmt.Println(us.get(u.Name))
}

type User struct {
	Name    string
	Posts   []*Post
	Friends []*User
}

func newUser(name string) *User {
	return &User{
		Name:    name,
		Posts:   make([]*Post, 0),
		Friends: make([]*User, 0),
	}
}

func (u *User) makeFriend(userFriend *User) {
	u.Friends = append(u.Friends, userFriend)
}

func (u *User) addPost(p *Post) {
	u.Posts = append(u.Posts, p)
}

type Post struct {
	Author  string
	Content string
}

func newPost(author, content string) *Post {
	return &Post{
		Author:  author,
		Content: content,
	}
}

// Необходимо написать наше хранилище с помощью map - ов

// map для юзеров
type UserStorage map[string]*User

func newUserStorage() UserStorage {
	return make(map[string]*User)
}

// создадим метод put ,который положить нашего пользователя в storage
func (us UserStorage) put(u *User) {
	us[u.Name] = u
}

// метод get ,который вернет его из storage
// Ошибка в конце фукции ,это общий паттент во всех функциях в GO (error)!!!!!!!!!!!!!!!!!!
func (us UserStorage) get(userName string) (*User, error) {
	u, ok := us[userName]
	if !ok {
		return nil, fmt.Errorf("no such user")
	}
	return u, nil
}

// map для постов
type PostStorage map[string]*Post

func newPostStorage() PostStorage {
	return make(map[string]*Post)
}

// создадим метод put ,который положить нашего пользователя в storage
func (ps PostStorage) put(p *Post) {
	ps[p.Author] = p
}

// метод get ,который вернет его из storage
// Ошибка в конце фукции ,это общий паттент во всех функциях в GO (error)!!!!!!!!!!!!!!!!!!
func (ps PostStorage) get(authorName string) (*Post, error) {
	u, ok := ps[authorName]
	if !ok {
		return nil, fmt.Errorf("no such user")
	}
	return u, nil
}
