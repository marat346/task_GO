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

	for {

		text1, _, err := bufio.NewReader(os.Stdout).ReadLine()

		if err != nil {
			log.Fatalln(err)
		}

		if _, err := d1.Write([]byte(text1)); err != nil {
			log.Fatalln(err)
		}

		upperText1 := []byte{}

		if _, err := d1.Read(upperText1); err != nil {
			log.Fatalln(err)
		}

	}
}
