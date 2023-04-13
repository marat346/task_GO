package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

// Имя структуры должно быть экспортируемо (с бол.буквы) и поля
type User struct {
	Name string `json:"name"`
	Age  int    `json:"age`
}

// Создаем метод toString

func (u *User) toString() string {
	return fmt.Sprintf("name is %s and age is %d \n", u.Name, u.Age)
}

type serviceNew struct {
	store map[string]*User
}

func main() {

	mux := http.NewServeMux()
	srv := serviceNew{make(map[string]*User)}

	mux.HandleFunc("/create", srv.Create)
	mux.HandleFunc("/get", srv.GetAll)

	fmt.Println("server is running")
	http.ListenAndServe("localhost:8080", mux)
}

// Метод Запрос имеет POST
func (s *serviceNew) Create(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		content, err := ioutil.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// ВАЖНО ЗАКРЫВАТЬ ТЕЛО
		defer r.Body.Close()

		var u User

		if err := json.Unmarshal(content, &u); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte(err.Error()))
			return
		}

		// Убрали отсюда строку
		//  splittedContent := strings.Split(string(content), " ")

		//Изменил строку
		// s.store[splittedContent[0]] = string(content)

		s.store[u.Name] = &u

		w.WriteHeader(http.StatusCreated)
		w.Write([]byte("User was created" + u.Name))
	}

	w.WriteHeader(http.StatusBadRequest)
}

// Метод GET который возращает всех пользователей

func (s *serviceNew) GetAll(w http.ResponseWriter, r *http.Request) {
	if r.Method == "GET" {
		response := " "
		for _, user := range s.store {
			// Вызываем toString()
			response += user.toString()
		}
		w.WriteHeader(http.StatusOK)
		w.Write([]byte(response))
		return
	}
	w.WriteHeader(http.StatusBadRequest)
}
