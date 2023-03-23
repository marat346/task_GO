package main

import "fmt"

type Person struct {
	Age  int
	Name string
}

// как объявляется метод структуры
// (p Person) - это ресивер ,он указывается с первой буквы структуры(Person -> p)
func (p Person) Print() {
	fmt.Println(p)
}

// Метод Ресивет по указателю
func (p *Person) SetNameByPointer(n string) {
	p.Name = n
}

// Метод Ресивет по копии
func (p Person) SetNameByValue(n string) {
	p.Name = n
}

func main() {
	p := Person{30, "Corado"}
	// она определенна только на структуре p.Print()
	p.Print() // напрямую уже не можем Print()
	fmt.Println("----------------------")

	p.SetNameByValue("Yna") // через копию не меняется
	fmt.Println(p, "не меняется")
	p.SetNameByPointer("Ycha") // через указатель меняется
	fmt.Println(p, "меняется")
}
