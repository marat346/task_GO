package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(2)
	go putBook1(&wg)
	go deliverBook1(&wg)

	wg.Wait()
	burnBook1()
}

func putBook1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("складываю книги")
}

func deliverBook1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("доставляю книги")
}

func burnBook1() {
	fmt.Println("сжигаю книгу")
}
