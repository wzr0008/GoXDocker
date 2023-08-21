package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Hello World")
	http.HandleFunc("/", heartResponse)
	http.HandleFunc("/health", healthResponse)
	err := http.ListenAndServe(":9092", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func healthResponse(writer http.ResponseWriter, request *http.Request) {
	writer.WriteHeader(http.StatusOK)
	fmt.Fprint(writer, "HELLO wORLD")
	return
}

func heartResponse(writer http.ResponseWriter, request *http.Request) {
	_, err := fmt.Fprintf(writer, "<3")
	if err != nil {
		log.Fatal(err)
	}
}
