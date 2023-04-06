package main

import (
	"fmt"
)

func main() {
	intChan := make(chan int)
	// ОБЩИЙ ПАТТЕРН ЗАКРЫВАТЬ КАНАЛ В ЗАПИСЫВАЮЩЕЙ ГОРУТИНЕ
	go func() {
		for i := 0; i < 4; i++ {
			intChan <- i
		}
		close(intChan)
	}()

	// Читаем из нашего канала через беск.цикл
	// В ЧИТАЮЩЕЙ ГОРУТИНЕ КАНАЛ НЕ ЗАКРЫВАЕМ!!!
	// ПЕРВЫЙ СПОСОБ
	for {
		val, ok := <-intChan
		if !ok {
			break
		}
		fmt.Println(val, "первый способ")
	}

	//ВТОРОЙ СПОСОБ

	for val1 := range intChan {
		fmt.Println(val1, "второй способ")
	}

}
