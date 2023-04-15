package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

const addr string = "localhost: 8085"

func main() {

	fmt.Println("server is running")
	http.HandleFunc("/", handle)
	log.Fatalln(http.ListenAndServe(addr, nil))
}

func handle(w http.ResponseWriter, r *http.Request) {

	bodyBytes, err := ioutil.ReadAll(r.Body)
	if err != nil {
		log.Fatalln(err)
	}

	text := string(bodyBytes)
	response := "1 instance: " + text + "\n"

	if _, err := w.Write([]byte(response)); err != nil {
		log.Fatalln(err)
	}
}
