package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(rw http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(rw, "Hello this is the base route")
	})

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error in starting the server")
	}
}
