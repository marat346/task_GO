package main

import "fmt"

// создание структуры с полями
type Person struct {
	age  int
	name string
}

func main() {
	p := Person{40, "Marat"}
	fmt.Println(p)
	fmt.Println("----------------")

	// создание указателя на структуру &person
	pPointer := &p
	fmt.Println(pPointer)
	fmt.Println("----------------")

	// создание новой структуры,на прямую указатель на структуру
	student := &Person{24, "John"}
	fmt.Println(student)
	fmt.Println("----------------")
	//Необязательно через оперсанд "&" иниц.указатель
	//Можно иницианализировать указатель с помощью встроенный функции NEW
	pNew := new(Person)
	fmt.Println(pNew)
}
