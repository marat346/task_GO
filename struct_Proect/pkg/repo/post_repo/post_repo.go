package post_repo

import (
	"fmt"
	"main/struct_Proect/pkg/post"
)

// map для постов
type PostStorage map[string]*post.Post

func NewPostStorage() PostStorage {
	return make(map[string]*post.Post)
}

// создадим метод put ,который положить нашего пользователя в storage
func (ps PostStorage) Put(p *post.Post) {
	ps[p.Author] = p
}

// метод get ,который вернет его из storage
// Ошибка в конце фукции ,это общий паттент во всех функциях в GO (error)!!!!!!!!!!!!!!!!!!
func (ps PostStorage) Get(authorName string) (*post.Post, error) {
	u, ok := ps[authorName]
	if !ok {
		return nil, fmt.Errorf("no such user")
	}
	return u, nil
}
