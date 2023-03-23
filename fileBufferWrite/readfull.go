package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("poem.txt")
	if err != nil {
		fmt.Println("не смогли открыть файл", err)
		return
	}
	defer f.Close()

	buf1 := make([]byte, 56)
	if _, err := io.ReadFull(f, buf1); err != nil {
		fmt.Println("Не смогли прочитать достаточное кол-во байтов", err)
		return
	}
	fmt.Println(string(buf1))

	buf2 := make([]byte, 26)
	_, err = f.Read(buf2)
	if err != nil {
		fmt.Println("Не смогли прочитать достаточное кол-во байтов", err)
		return
	}
	fmt.Println(string(buf2))

	s, err := f.Stat()
	if err != nil {
		fmt.Println("fall")
		return
	}
	fmt.Printf("Size: %d\n", s.Size())
}
