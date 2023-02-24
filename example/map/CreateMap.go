package main

import "fmt"

func main() {
	m := make(map[int]string)
	fmt.Println(m)
	// добавление в map
	m[1] = "first"
	m[2] = "second"
	fmt.Println(m)

	fmt.Println(m[2])
}
