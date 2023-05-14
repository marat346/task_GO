package main

import (
	"log"
	"os"
)

func main() {
	file, err := os.Create("marat.txt")
	if err != nil {
		log.Fatal(err)
	}

	file.WriteString("privet marat")
}
