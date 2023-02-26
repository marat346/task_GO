package main

import "fmt"

func main() {
	// передаем различные типы
	PrintType(1)
	PrintType("При условии неизвестно какой тип")

	// var empty interface{} = Person{}
	// empty.???   И как видно на данном интерфейсе нету не каких методов
}

// ЭТО БУДЕТ ПОЛЕЗНО КОГДА МЫ НЕ ЗНАЕМ КАКОЙ ТИП ПРИЙДЕТ В ФУНКЦИЮ
func PrintType(i interface{}) {
	switch i.(type) {
	case string:
		fmt.Println("string")
	case int:
		fmt.Println("int")
	}
}
