package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const proxyAddr string = "localhost:9000"

var (
	counter            int    = 0
	firstInstanceHost  string = "http://localhost:8085"
	secondInstanceHost string = "http://localhost:8086"
)

func main() {

	fmt.Println("server is running")
	http.HandleFunc("/", handProxe)
	log.Fatalln(http.ListenAndServe("localhost:9000", nil))
}

func handProxe(w http.ResponseWriter, r *http.Request) {

	textBytes, err := ioutil.ReadAll(r.Body)

	if err != nil {
		log.Fatalln(err)
	}

	text := string(textBytes)

	if counter == 0 {
		resp, err := http.Post(firstInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
		if err != nil {
			log.Fatalln(err)
		}
		counter++

		textBytes, err = ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Fatalln(err)
		}
		defer resp.Body.Close()

		fmt.Println(string(textBytes))

		return
	}

	resp, err := http.Post(secondInstanceHost, "text/plain", bytes.NewBuffer([]byte(text)))
	if err != nil {
		log.Fatalln(err)
	}

	textBytes, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		log.Fatalln(err)
	}
	defer resp.Body.Close()

	fmt.Println(string(textBytes))

	counter--

}
