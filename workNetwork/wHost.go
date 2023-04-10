package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	lis1, err := net.Listen("tcp4", "localhost:9090")
	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("server is running")

	con1, err := lis1.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {

		line1, err := bufio.NewReader(con1).ReadString('\n')
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("line:", string(line1))

		upperLine1 := strings.ToUpper(string(line1))

		if _, err := con1.Write([]byte(upperLine1)); err != nil {
			log.Fatalln(err)
		}
	}

}
