package main

import "fmt"

func main() {

	// Направленные каналы
	dataChan := getData()

	for val := range dataChan {
		fmt.Println(val)
	}
}

// Сделали канал только на запись
// Возращает канал только для чтения <-chan int
func getData() <-chan int {
	dataChan := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			dataChan <- i
		}
		close(dataChan)
	}()

	return dataChan
}
