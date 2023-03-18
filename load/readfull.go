package main

import (
	"fmt"
	"os"
)

func main() {
	f, err := os.Open("poem.txt")
	if err != nil {
		fmt.Println("не смогли открыть файл", err)
		return
	}
	defer f.Close()

	//buf := make([]byte, 56)
	//if _, err := io.ReadFull(f, buf); err != nil {
	//	fmt.Println("Не смогли прочитать достаточное кол-во байтов", err)
	//	return
	//}
	//fmt.Println(string(buf))

	//buf := make([]byte, 26)
	//_, err = f.Read(buf)
	//if err != nil {
	//	fmt.Println("Не смогли прочитать достаточное кол-во байтов", err)
	//	return
	//}
	//fmt.Println(string(buf))

	s, err := f.Stat()
	if err != nil {
		fmt.Println("fall")
		return
	}
	fmt.Printf("Size: %d\n", s.Size())
}
