package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	lis1, err := net.Listen("tcp4", "localhost:8181")
	if err != nil {
		log.Fatalln(err)
	}
	fmt.Println("server is running")

	con, err := lis1.Accept()
	if err != nil {
		log.Fatalln(err)
	}

	for {
		line1, err := bufio.NewReader(con).ReadString('\n')

		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("line:", string(line1))

		upperLine1 := strings.ToUpper(line1)

		if _, err := con.Write([]byte(upperLine1)); err != nil {
			if err != nil {
				log.Fatalln(err)
			}
		}

	}

}
