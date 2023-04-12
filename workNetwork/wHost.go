package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {
	ears, err := net.Listen("tcp4", "localhost:9191")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("server is running")

	mik, err := ears.Accept()

	if err != nil {
		log.Fatalln(err)
	}

	for {
		line, err := bufio.NewReader(mik).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("line:", string(line))

		upperLine := strings.ToUpper(line)

		if _, err := mik.Write([]byte(upperLine)); err != nil {
			log.Fatalln(err)
		}
	}
}
