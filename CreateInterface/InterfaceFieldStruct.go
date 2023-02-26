package main

import "fmt"

// Так же интерфейс может являться полем структуры

type IWorker interface {
	work() string
}

func (m Men) work() string {
	fmt.Println("метод")
}

type Men struct {
	Name string
	Age  int
	// и объявим его полем структуры
	w IWorker
}

func main() {

}
