package main

import (
	"bufio"
	"log"
	"net"
	"strings"
)

func main() {

	lis, err := net.Listen("tcp4", "localhost:8080")

	if err != nil {
		log.Fatalln(err)
	}

	con, err := lis.Accept()

	if err != nil {
		log.Fatalln(err)
	}

	fmt.Println("server is running")

	for {
		reder := bufio.NewReader(con)
		line, _, err := reder.ReadLine()
		if err != nil {
			log.Fatalln(err)
		}
		upperLine := strings.ToUpper(string(line))
		if _, err := con.Write([]byte (upperLine));err != nil {
			log.Fatalln(err)
		}
	}
