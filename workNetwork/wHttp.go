package main

import (
	"fmt"
	"net/http"
)

func main() {

	mux := http.NewServeMux()

	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("марат молодец не останавливайся!!!!!!!"))
	})

	fmt.Println("server is running")

	http.ListenAndServe("localhost:5656", mux)

}
