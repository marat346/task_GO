package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int)

	go func() {
		for i := 0; i < 4; i++ {
			intChan <- i
		}
	}()

	// Читаем из нашего канала через беск.цикл
	for {
		val := <-intChan
		fmt.Println(val)
	}

}

func putBook1(rchan chan string) {

	fmt.Println("складываю книги")
}

func deliverBook1(rchan chan string) {

	fmt.Println("доставляю книги")
}

func burnBook1(rchan chan string) {
	fmt.Println("сжигаю книги")
}
