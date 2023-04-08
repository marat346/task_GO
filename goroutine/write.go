package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg1 sync.WaitGroup
	wg1.Add(2)

	go one(&wg1)
	go one1(&wg1)

	wg1.Wait()
	one2()

}

func one(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("сладываю тележку")
}

func one1(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("перевожу книги")
}

func one2() {
	fmt.Println("сжигаю книги")
}
