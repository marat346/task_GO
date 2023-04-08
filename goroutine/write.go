package main

import "fmt"

func main() {
	oneChan := make(chan string)

	go func() {

		oneChan <- "marat molodec"
	}()

	a := <-oneChan

	fmt.Println(a)

}

func one() {

}

func one1() {

}

func one2() {

}
