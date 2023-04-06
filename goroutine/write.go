package main

import "fmt"

func main() {

	chanNew := make(chan string)

	go func() {
		chanNew <- "mARAT BEST"
	}()

	a := <-chanNew
	fmt.Println(a)
}
