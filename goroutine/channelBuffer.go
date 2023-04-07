package main

import "fmt"

func main() {

	// Буфер вмешает три знач.,если +1 буфер заблокируется
	bufChan := make(chan int, 3)
	bufChan <- 1
	bufChan <- 1
	bufChan <- 1
	// bufChan <- 1

	close(bufChan) // закроем канал чтоб не вышла ошибка

	v, ok := <-bufChan
	fmt.Println(v, ok)

	// Полезные функции Len() Cap()
	fmt.Println(len(bufChan), cap(bufChan))

}
