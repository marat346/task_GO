package main

import (
	"fieldOfView/person"
	"fmt"
)

func main() {

	p := person.Person{
		Name: "Sobrano Tony",
		Age:  40,
		// Если написано поле с маленькой буквы sibiling :"aa",значит поле НЕЭКСПОРТИРУЕМА
		Sibiling: "Экпортируемое поле",
	}

	fmt.Println(p)
	fmt.Println("-------------------")
	// вызываем функцию PrintInfo
	fmt.Println(person.PrintInfo(p))
	fmt.Println("-------------------")
	// в main чего не знает о функции person.info
	fmt.Println("-------------------")
	// если константа с мал.буквы main о ней не знает
	fmt.Println(person.C)
	fmt.Println("-------------------")
	//способ сделать видимой константу
	fmt.Println(person.GetString())
}
