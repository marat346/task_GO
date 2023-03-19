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

	file, err := os.Create("ioutil.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	_, err = file.Write(b.Bytes())
	if err != nil {
		panic(err)
	}

	endLog, err = ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

}
