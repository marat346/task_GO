package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	f, err := os.Open("fod.txt")
	if err != nil {
		fmt.Println("fall", err)
		return
	}
	defer f.Close()

	buf := make([]byte, 256)
	if _, err := io.ReadFull(f, buf); err != nil {
		fmt.Println("fall", err)
		return
	}
	defer f.Close()

	fmt.Println(string((f)))
}
