package main

import "fmt"

func main() {

	m := func(a int) (int, bool) { return a, true }

	if v, ok := m(5); ok {
		fmt.Println("true")
		fmt.Println(v)
	}

}
