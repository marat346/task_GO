package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
)

func main() {

	lis, err := net.Listen("tcp4", "localhost:8081")

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("server is running")
	con, err := lis.Accept()

	if err != nil {
		log.Fatalln(err)
	}

	for {
		reder := bufio.NewReader(con)
		line, _, err := reder.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}

		fmt.Println("line recieved:", line)
		upperLine := strings.ToUpper(string(line))
		if _, err := con.Write([]byte(upperLine)); err != nil {
			log.Fatalln(err)
		}
	}
}
