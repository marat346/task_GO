package main

import "fmt"

func main() {

	// map ( [ТИП КЛЮЧА] ТИП ЗНАЧЕНИЯ )
	m := make(map[int]string)
	fmt.Println(m)

	// добавление в map
	m[1] = "first"
	m[2] = "second"
	fmt.Println(m, "ключ : значение")

	fmt.Println(m[2], "обращение к ключу 2 ")

	// удаление из типа map
	delete(m, 2)
	fmt.Println(m, "удаление")
}
