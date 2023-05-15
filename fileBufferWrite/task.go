package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
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

	defer file.Close()

	write := bufio.NewWriter(file)

	write.WriteString(text)
	write.Flush()

}

func bytesBuffer() {
	text := "new text"

	file, err := os.Create("buffer.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var b bytes.Buffer

	b.WriteString(text)

	_, err = file.Write(b.Bytes())
	if err != nil {
		log.Fatal(err)
	}
}

func ioutilReadAll() {

	text := "Rune literal is expressed as one or more characters in single quotes, excluding unquoted single quotes and newlines."

	// ioutil.WriteFile сам автоматически создает файл.
	fileName := "newLog.txt"

	var b bytes.Buffer
	b.WriteString(text)

	if err := ioutil.WriteFile(fileName, b.Bytes(), 0666); err != nil {
		log.Fatal(err)
	}

	f, err := os.Open(fileName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	resultBytes, err := ioutil.ReadAll(f)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Сохраненный лог:")
	fmt.Println(string(resultBytes))
}

func main() {
	// createFileCloseFile()
	// readFull()
	// read()
	// writeBufioNewWriteFlush()
	// bytesBuffer()
	ioutilReadAll()
}
