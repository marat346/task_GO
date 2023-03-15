package main

import "fmt"

func main() {

	m := map[int]string{
		1: "fist",
		2: "second",
		3: "third",
	}

	v, ok := m[3]
	fmt.Println(v, "это зероноль")
	fmt.Println(ok, "это true and false")
}
