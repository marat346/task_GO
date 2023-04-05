package main

import (
	"fmt"
)

func main() {

	respChan := make(chan string)

	go putBook1()
	go deliverBook1()

	burnBook1()
}

func putBook1(rchan chan string) {

	fmt.Println("складываю книги")
}

func deliverBook1(rchan chan string) {

	fmt.Println("доставляю книги")
}

func burnBook1() {
	fmt.Println("сжигаю книги")
}
