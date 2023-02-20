package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type User struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	Friends []int  `json:"friends"`
}

type storage struct {
	users      map[int]*User
	lastUserId int
}

type CreateUserResponse struct {
	Id int `json:"id"`
}

var srv = newStorage()
var u User

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/create", srv.Create)
	http.ListenAndServe("localhost:8080", mux)

}

func (s *storage) Create(w http.ResponseWriter, r *http.Request) {
	content, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	defer r.Body.Close()

	if err := json.Unmarshal(content, &u); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte(err.Error()))
		return
	}
	userId := srv.AddNewUser(&u)
	fmt.Println(s.users[userId])
	response, _ := json.Marshal(CreateUserResponse{Id: userId})

	w.Header().Set("Content-Type", "application/json")
	w.Write(response)
	w.WriteHeader(http.StatusCreated)

}

func newStorage() storage {
	return storage{
		users:      make(map[int]*User),
		lastUserId: 0,
	}
}

func (s *storage) AddNewUser(u *User) int {
	s.users[s.lastUserId] = u
	s.lastUserId++
	return s.lastUserId - 1
}
