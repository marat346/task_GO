package user_repo

import (
	"fmt"
	"main/struct_Proect/pkg/user"
)

// map для юзеров
type UserStorage map[string]*user.User

func NewUserStorage() UserStorage {
	return make(map[string]*user.User)
}

// создадим метод put ,который положить нашего пользователя в storage
func (us UserStorage) Put(u *user.User) {
	us[u.Name] = u
}

// метод get ,который вернет его из storage
// Ошибка в конце фукции ,это общий паттент во всех функциях в GO (error)!!!!!!!!!!!!!!!!!!
func (us UserStorage) Get(userName string) (*user.User, error) {
	u, ok := us[userName]
	if !ok {
		return nil, fmt.Errorf("no such user")
	}
	return u, nil
}
