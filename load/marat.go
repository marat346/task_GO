package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("remember.txt")
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println("fall", err)
		return
	}
	defer file.Close()

	writer.WriteString("marat go to sushu")

	//ОБЯЗАТЕЛЬНО ОЧИЩАЕМ БУФЕР!!!
	writer.Flush()

}
