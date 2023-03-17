package main

import (
	"fmt"
	"os"
)

func main() {
	text := "Hello Marat"
	file, err := os.Create("hello.txt")
	if err != nil {
		fmt.Println("Не смогли создать файл", err)
		return
	}
	//важно закрывать файл!!!
	defer file.Close()

	file.WriteString(text)
	fmt.Println(file.Name())
}
