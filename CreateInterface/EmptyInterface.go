package main

import "fmt"

type IEmployee interface {
	Printer()
}

type Employee struct {
	Name string
	Age  int
}

func (p Employee) Printer() {
	fmt.Println(p)
}

// func main() {
// var empty interface{} = Employee {}
// empty.???   И как видно на данном интерфейсе нету не каких методов
// }

func main() {
	// передаем различные типы
	PrintType(1)
	PrintType("При условии неизвестно какой тип")
}

// ЭТО БУДЕТ ПОЛЕЗНО КОГДА МЫ НЕ ЗНАЕМ КАКОЙ ТИП ПРИЙДЕТ В ФУНКЦИЮ
func PrintType(i interface{}) {

	// конструкция typeswitching
	switch i.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	}
}
