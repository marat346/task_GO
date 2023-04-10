package main

import (
	"net/http"
)

func main() {

	mux := http.NewServeMux()
	mux.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Esenin , born, in ,ict ,home ,of ,his ,grandfather, who ,was ,an ,Old, Believer. He ,went ,to ,Moscow."))
	})

	http.ListenAndServe("localhost:6060", mux)
}
