package main

import "fmt"

// создание структуры с полями
type OneMen struct {
	age  int
	name string
}

func main() {
	student := &OneMen{24, "John"}
	fmt.Println(student)
	// Обращение к полям
	fmt.Println(student.name)
	fmt.Println(student.age)
	fmt.Println("-----------------------")
	// Изменеие полей структуры
	student.name = "Firdaus"
	fmt.Println(student.name)

}
