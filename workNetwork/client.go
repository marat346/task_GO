package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {
	d, err := net.Dial("tcp4", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}

	for {
		text, _, err := bufio.NewReader(os.Stdout).ReadLine()
		if err != nil {
			log.Fatalln(err)
		}

		if _, err = d.Write([]byte(text)); err != nil {
			log.Fatalln(err)
		}

		upperText := []byte{}

		if _, err := d.Read(upperText); err != nil {
			log.Fatalln(err)
		}

	}
}
