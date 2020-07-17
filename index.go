package main

import (
	"log"
	"net/http"
	"path"
)

func serveIndex(w http.ResponseWriter, r *http.Request) {
	log.Printf("%s %s\n", r.Method, r.URL)
	http.ServeFile(w, r, path.Join("static", r.URL.String()))
}
