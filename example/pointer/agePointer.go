package main

import "fmt"

func main() {
	var a int
	var b int
	var c int
	fmt.Println(&a)
	fmt.Println(&b)
	fmt.Println(&c)
	f := &a
	fmt.Println(f, "это F")
}
