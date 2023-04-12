package main

import (
	"bufio"
	"log"
	"net"
	"os"
)

func main() {

	phone, err := net.Deal("tcp4", "localhost:9191")

	if err != nil {
		log.Fatalln(err)
	}

	for {

		text1, _, err := bufio.NewReader(os.Stdout).ReadLine()
		if err != nil {
			log.Fatalln(err)
		}

		if _, err := phone.Write([]byte(text1)); err != nil {
			log.Fatalln(err)
		}

		upperText1 := []byte{}

		if _, err := phone.Read(upperText1); err != nil {
			log.Fatalln(err)
		}
	}
}
