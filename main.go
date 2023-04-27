package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello world this is testing")
}

func main() {
	http.HandleFunc("/", mainHandler)

	err := http.ListenAndServe(":8080", nil)

	if err != nil {
		log.Fatal("Error in starting the server", err)
	}
}
