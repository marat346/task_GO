package main

import "fmt"

// Простое понимание ключевого слова defer
// Выполняется действие defer ,после выполнения всех действий в фуекции
func iAmUsingDefer() {
	defer fmt.Println("exiting function _______ВТОРОЙ ВЫХОД")
	fmt.Println("before defer______ПЕРВВЫЙ ВЫХОД")
}

func main() {
	iAmUsingDefer()
	fmt.Println("DONE______ТРЕТИЙ ВЫХОД")

}
