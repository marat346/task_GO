package main

import (
	"fmt"
)

func main() {
	tc := one()
	fc := one1(tc)
	ec := one2(fc)

	fmt.Println(<-ec)

}

func one() chan string {
	fistChan := make(chan string)

	go func() {
		fistChan <- "сладываю тележку"
		close(fistChan)
	}()
	fmt.Println(<-fistChan)
	return fistChan
}

func one1(fistChan chan string) chan string {
	secondChan := make(chan string)

	go func() {
		secondChan <- "перевожу книги"
		close(secondChan)
	}()
	fmt.Println(<-secondChan)
	return secondChan
}

func one2(secondChan chan string) chan string {
	thirdChan := make(chan string)

	go func() {
		thirdChan <- "сжигаю книги"
		close(thirdChan)
	}()
	fmt.Println(<-thirdChan)
	return thirdChan
}
