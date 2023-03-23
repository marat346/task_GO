package main

import "fmt"

// Так же интерфейс может являться полем структуры

type IWorker interface {
	Work() string
}

func (m Men) Work() {
	fmt.Println(m)
}

type Men struct {
	Name string
	Age  int
	// и объявим его полем структуры
	w IWorker
}

func main() {
	p := Men{"Makar", 40}

}
