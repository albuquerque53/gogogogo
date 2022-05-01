package main

import (
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", func(rw http.ResponseWriter, r *http.Request) {
		rw.Write([]byte("Hello, World!"))
	})

	log.Fatal(http.ListenAndServe(":2001", nil))
}
