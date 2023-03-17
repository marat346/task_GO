package main

import (
	"fmt"
	"log"
	"os"
)

func main() {
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(path string) *os.File {
	log.Println("Creating file")
	f, err := os.Create(path)
	if err != nil {
		log.Fatalf("err while creating file:%v", err)
	}
	return f
}

func writeFile(f *os.File) {
	log.Println("Writing file")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	log.Println("Closing file")
	err := f.Close()
	if err != nil {
		fmt.Fprintf(os.Stderr, "error:%v\n", err)
		os.Exit(1)
	}
}
