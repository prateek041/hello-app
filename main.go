package main

import (
	"fmt"
	"log"
	"net/http"
)

func mainHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "Hello world this is testing")
}

func secondaryHandler(rw http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(rw, "This is the secondary route")
}

func main() {
	http.HandleFunc("/", mainHandler)
	http.HandleFunc("/second", secondaryHandler)

	err := http.ListenAndServe(":8081", nil)

	if err != nil {
		log.Fatal("Error in starting the server", err)
	}
}
