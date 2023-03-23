package main

import "fmt"

// создание структуры с полями
type Men struct {
	age  int
	name string
}

func main() {
	// инициализация структуры
	p := Men{
		age:  21,
		name: "Peter",
	}
	fmt.Println(p)
	fmt.Println("---------------------------")
	// другой способ инициализации
	p = Men{40, "John"}
	fmt.Println(p)
	fmt.Println("---------------------------")
	//без значений ,некоторые типы будут получать ZeroValue
	pZeroValue := Men{}
	fmt.Println(pZeroValue)
}
