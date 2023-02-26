package main

import "fmt"

// Это значит интерфейс IBaker может принимать значения типа DogName
//
//	↓▌↓ ↓ ↓ ↓
type IBarker interface {
	Bark() string
}

//	↑ ↑ ↑ ↑ ↑ ↑ ↑
//
// на типе DogName определен метод Bark
type DogName string

func (dn DogName) Bark() string {
	return string(dn) + ": aw aw aw aw"
}

// Создаем функцию PrintBark,и передаем в эту функцию Интерфейс IBarker
func PrintBark(i IBarker) {
	fmt.Println(i.Bark(), "это передача ИНтерфейса в функцию")
}

func main() {

	dn := DogName("sharik")
	fmt.Println(dn.Bark())

	// как работать с интерфейсом ,создать переменную типа Ibaker
	// передаем переменную dn ,или сразу создать
	var Idn IBarker = dn
	// у переменной Idn вызвать метод из интерфейса
	fmt.Println(Idn.Bark(), "это метод как работать с Интерфейсом")
	// Вызываем функцию PrintBark
	PrintBark(dn)
}
