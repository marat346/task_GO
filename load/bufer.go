package main

import (
	"bufio"
	"fmt"
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
	writer.WriteString("Say hi")
	writer.Write([]byte{65, 66, 67})
	writer.WriteString("\n")
	writer.Flush()

}
