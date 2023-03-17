package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("log.txt")
	if err != nil {
		fmt.Println("Ошибка открытия файла", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 150)
	_, err = f.Read(buf)
	if err != nil {
		fmt.Println("Не смогли прочитать последовательность байтов из файла", err)
		return
	}
	fmt.Println(string(buf))
}
