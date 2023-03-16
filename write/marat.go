package main

import "fmt"

func main() {

	m := func(a int) (int, bool) { return a, true }

	if _, ok := m(5); ok {
		fmt.Println("true")
	}

}
