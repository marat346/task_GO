package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {

	d1, err := net.Dial("tcp4", "localhost:9090")
	if err != nil {
		log.Fatalln(err)
	}

	text, _, err := bufio.NewReader(os.Stdout).ReadLine()

	if err != nil {
		log.Fatalln(err)
	}

	if _, err := d1.Write([]byte(text)); err != nil {
		log.Fatalln(err)
	}

	upperText := []byte{}

	if _, err := d1.Read(upperText); err != nil {
		log.Fatalln(err)
	}

}