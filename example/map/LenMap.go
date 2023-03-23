package main

import "fmt"

func main() {
	// сколько ключей записано в карту
	mLen := map[int]string{
		1: "first",
		2: "second",
		3: "third",
	}

	fmt.Println(len(mLen))
	fmt.Println(mLen)
}
