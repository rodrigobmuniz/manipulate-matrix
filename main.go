package main

import (
	"fmt"
	"log"
	matrix "manipulate-matrix/app/delivery/http"
	"net/http"
)

func main() {
	http.HandleFunc("/echo", matrix.Echo)
	http.HandleFunc("/invert", matrix.Invert)
	http.HandleFunc("/flatten", matrix.Flatten)
	http.HandleFunc("/sun", matrix.Sum)
	http.HandleFunc("/multiply", matrix.Multiply)
	startServe()
}

func startServe() {
	fmt.Printf("Starting server at localhost:8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
