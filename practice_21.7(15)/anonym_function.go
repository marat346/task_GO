package main

import (
	"fmt"
)

func main() {
	anonym(7, 4, func(x, y int) int { return x + y })
	anonym(7, 4, func(x, y int) int { return x * y })
	anonym(7, 4, func(x, y int) int { return x / y })
}
func anonym(x int, y int, A func(x, y int) int) {
	defer fmt.Println(A(x, y))
}