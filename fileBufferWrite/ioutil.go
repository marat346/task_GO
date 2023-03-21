package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

func main() {
	var b bytes.Buffer

	fmt.Println("Введите логин:")
	var login string
	fmt.Scan(&login)

	fmt.Println("Введите пароль:")
	var password string
	fmt.Scan(&password)

	fmt.Println("Введите возраст:")
	var age string
	fmt.Scan(&age)

	b.WriteString(fmt.Sprintf("Ваш логин:%s\n", login))
	b.WriteString(fmt.Sprintf("Ваш пароль:%s\n", password))
	b.WriteString(fmt.Sprintf("Ваш возраст:%s\n", age))

	fileName := "hello2.txt"

	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		panic(err)
	}

	file, err := os.Open(fileName)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	resulBytes, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	fmt.Println("Сохраненый лог:")
	fmt.Println(string(resulBytes))
}
