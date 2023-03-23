package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Введите имя пользователя:")
	var username string
	fmt.Scan(&username)
	fmt.Println("Введите пароль:")
	var password string
	fmt.Scan(&password)
	fmt.Println("Введите возраст:")
	var age int
	fmt.Scan(&age)

	file, err := os.Create("creadentials.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()

	var b bytes.Buffer

	b.WriteString(fmt.Sprintf("Ваш логин:%s \n", username))
	b.WriteString(fmt.Sprintf("Ваш пароль:%s \n", password))
	b.WriteString(fmt.Sprintf("Ваш возраст:%d", age))

	_, err = file.Write(b.Bytes())
	if err != nil {
		fmt.Println(err)
		return
	}
}
