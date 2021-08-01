package main

import (
	"fmt"
	"log"
	"manipulate-matrix/delivery/http"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", delivery.Echo)
	startServe()
}

func startServe() {
	fmt.Printf("Starting server at localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
