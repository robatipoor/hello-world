package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

var address string

func init() {
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}
	address = "0.0.0.0:" + port
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", handler)
	log.Println("Start Server ", address)
	log.Fatalln(http.ListenAndServe(address, mux))
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World !")
}
