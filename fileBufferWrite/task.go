package main

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func createFileCloseFile() {
	text := "Rune literal is expressed as one or more characters in single quotes, excluding unquoted single quotes and newlines.String literal is a concatenation of characters, a character sequence.That show st the darkness thou canst not dispel."

	file, err := os.Create("marat.txt")
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()
	file.WriteString(text)
}

func readFull() {
	f, err := os.Open("marat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 128)

	if _, err := io.ReadFull(f, buf); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", buf)
}

func read() {
	f, err := os.Open("marat.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	buf := make([]byte, 128)
	_, err = f.Read(buf)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(string(buf))
}

func writeBufioNewWriteFlush() {

	text := "marat good"

	file, err := os.Create("bufio.txt")
	if err != nil {
		log.Fatal(err)
	}

	write := bufio.NewWriter(file)

	write.WriteString(text)
	write.Flush()

}

func main() {
	// createFileCloseFile()
	// readFull()
	// read()
	writeBufioNewWriteFlush()

}
