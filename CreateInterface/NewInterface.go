package main

import "fmt"

type IPerson interface {
	Printer()
}

type Person struct {
	Name string
	Age  int
}

func (p Person) Printer() {
	fmt.Println(p)
}

func main() {
	p := Person{"Makar", 30}

	var iPrinter IPerson = p
	iPrinter.Printer()

}
