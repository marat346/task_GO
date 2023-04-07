package main

import "fmt"

func main() {

	fc1 := one()
	hc1 := one1(fc1)
	sc1 := one2(hc1)

	fmt.Println(<-sc1)

}

func one() chan string {
	firstChan1 := make(chan string)
	go func() {
		firstChan1 <- "складываюю книги"
	}()

	return firstChan1
}

func one1(firstChan1 chan string) chan string {
	secondChan1 := make(chan string)
	fmt.Println(<-firstChan1)
	go func() {
		secondChan1 <- "доставляю книги"
	}()

	return firstChan1
}

func one2(secondChan1 chan string) chan string {
	thirdChan1 := make(chan string)
	fmt.Println(<-secondChan1)
	go func() {
		thirdChan1 <- "сжигаю книги"
	}()

	return thirdChan1
}
