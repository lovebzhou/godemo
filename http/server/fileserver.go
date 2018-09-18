package main

import (
	"log"
	"net/http"
)

func mainFileServer() {
	// Simple static webserver:
	log.Fatal(http.ListenAndServe(":8080", http.FileServer(http.Dir(".."))))
}
