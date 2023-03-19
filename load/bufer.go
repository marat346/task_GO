package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func main() {
	file, err := os.Create("buffer.txt")
	writer := bufio.NewWriter(file)
	if err != nil {
		fmt.Println("Ошибка при создание файла", err)
		return
	}
	defer file.Close()
	writer.WriteString("Say hi")
	writer.WriteString("\n")
	writer.WriteRune('R')
	writer.WriteString("\n")
	writer.WriteByte(67)
	writer.WriteString("\n")
	writer.Write([]byte{65, 66, 67})
	writer.WriteString("\n")
	writer.Flush()

	f, err := os.Open("buffer.txt")
	if err != nil {
		fmt.Println("не смогли открыть файл", err)
		return
	}

	buf := make([]byte, 10)
	if _, err = io.ReadFull(f, buf); err != nil {
		fmt.Println("Не смогли прочитать достаточное кол-во байтов", err)
		return
	}
	fmt.Println(string(buf))
}
