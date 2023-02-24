package main

import "fmt"

func main() {
	a := 10
	writeStr(a)
	fmt.Println(&a)
}

func writeStr(b int) *int {
	b = b + 13
	fmt.Println(&b, "функция")
	return &b

}
