package main

import "fmt"

// defer Необходим для обработки ошибок внутр.функции
// и выхода и закрытия чего-нибудь (файлов)
func iAmUsingDefer2() {
	defer func() {
		fmt.Println(1)
		fmt.Println(2)
	}()
	fmt.Println("before defer")
}

func main() {
	iAmUsingDefer2()
	fmt.Println("DONE")

}
