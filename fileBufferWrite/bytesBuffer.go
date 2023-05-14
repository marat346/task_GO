package main

import (
	"bytes"
	"fmt"
	"os"
)

func main() {

	fmt.Println("Введите логин:")
	var login string
	fmt.Scan(&login)

	fmt.Println("Введите пароль:")
	var password string
	fmt.Scan(&password)

	fmt.Println("Введите возраст:")
	var age string
	fmt.Scan(&age)

	file, err := os.Create("password.txt")
	if err != nil {
		fmt.Println("fall", err)
		return
	}
	defer file.Close()

	var b bytes.Buffer
	b.WriteString(fmt.Sprintf("Ваш логин:%s\n", login))
	b.WriteString(fmt.Sprintf("Ваш пароль:%s\n", password))
	b.WriteString(fmt.Sprintf("Ваш возраст:%s\n", age))

	_, err = file.Write(b.Bytes())
	if err != nil {
		fmt.Println("fall", err)
		return
	}

}
