package main

import (
	"log"
	"net/http"
	"os"
)

func main() {
	http.HandleFunc("/gate", GateEndpoint)
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), nil)
	if err != nil {
		log.Fatalln("Cannot listen:", err)
	}
}
