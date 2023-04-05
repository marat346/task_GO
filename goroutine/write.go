package main

import (
	"fmt"
	"sync"
)

func putBook1(gg *sync.WaitGroup) {
	defer gg.Done()
	fmt.Println("складываю книги")
}

func deliverBook1(gg *sync.WaitGroup) {
	defer gg.Done()
	fmt.Println("перевожу книги")
}

func burnBook1() {

	fmt.Println("сжигаю книги")
}

func main() {

	var gg sync.WaitGroup
	gg.Add(2)

	go putBook1(&gg)
	go deliverBook1(&gg)

	gg.Wait()
	burnBook1()

}
