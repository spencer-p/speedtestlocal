package main

import (
	"flag"
	"log"
	"net/http"
)

func main() {
	flag.Parse()

	log.Println("Booting up")
	http.HandleFunc("/", serveIndex)
	http.HandleFunc("/ws", serveWS)
	log.Println(http.ListenAndServe(":8080", nil))
}
