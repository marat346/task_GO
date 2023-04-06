package main

import "fmt"

func main() {

	fc := putBook2()
	sc := deliverBook2(fc)
	tc := burnBook2(sc)

	fmt.Println(<-tc)

}

func putBook2() chan string {

	firstChan := make(chan string)
	go func() {
		firstChan <- "собираю книги в тележку"
	}()

	return firstChan
}

func deliverBook2(firstChan chan string) chan string {
	secondChan := make(chan string)
	fmt.Println(<-firstChan)
	go func() {
		secondChan <- "везу тележку"
	}()
	return secondChan
}

func burnBook2(secondChan chan string) chan string {
	thirdChan := make(chan string)
	fmt.Println(<-secondChan)
	go func() {
		thirdChan <- "сжигаю книги"
	}()
	return thirdChan
}
