package main

import "fmt"

func main() {
	oneChan := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			oneChan <- i
		}
		close(oneChan)
	}()

	for {
		v, ok := <-oneChan
		if !ok {
			break
		}
		fmt.Println(v)
	}

}
