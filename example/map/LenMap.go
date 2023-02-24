package main

import "fmt"

func main() {
	mLen := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}

	fmt.Println(len(mLen))
}
