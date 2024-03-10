package main

import (
	"github.com/gorilla/handlers"
	"log"
	"net/http"
	"os"
)

func main() {
	http.Handle("/static/", http.FileServer(http.Dir("./static")))
	http.HandleFunc("/gate", GateEndpoint)
	err := http.ListenAndServe(os.Getenv("LISTEN_ADDRESS"), handlers.LoggingHandler(os.Stdout, http.DefaultServeMux))
	if err != nil {
		log.Fatalln("Cannot listen:", err)
	}
}
