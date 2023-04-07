package main

import (
	"fmt"
)

func main() {
	oneChan := make(chan string)

	go func() {
		a := <-oneChan
		fmt.Println(a)

	}()

	oneChan <- "marat molodec"

}
