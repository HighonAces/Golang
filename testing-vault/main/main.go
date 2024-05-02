package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {
	handler := http.NewServeMux()
	handler.HandleFunc("/api/readSecret", readSecret)
	http.ListenAndServe("0.0.0.0:8080", handler)
	// Listen to all request to port 8080.
}

func readSecret(w http.ResponseWriter, r *http.Request) {
	log.Println("Reading secret")
	data, err := os.ReadFile("/var/my-secret")
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("Data: %s", data)
	_, err = io.WriteString(w, string(data))
	if err != nil {
		return
	}
}
