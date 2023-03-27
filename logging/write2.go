package main

import (
	"fmt"
	"time"
)

func runAndWait() int {

	time.Sleep(time.Second * 1)
	return 10
}

func main() {
	for i := 0; i < 5; i++ {
		a := runAndWait()
		fmt.Println(time.Now(), "runAndWait finished...", "result", a)
	}

	fmt.Println("done")
}
