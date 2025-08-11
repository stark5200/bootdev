package main

import (
	//"fmt"
	"log"
	"net/http"
	//"path/filepath"
)

func main() {
	mux := http.NewServeMux()
	port := "8080"
	filepath := "."

	http_server := &http.Server{
		Addr:    ":" + port,
		Handler: mux,
	}

	mux.Handle("/", http.FileServer(http.Dir(filepath)))

	log.Printf("✅ created server and actively listening on port %s! and fileroot %s", port, filepath)
	log.Fatal(http_server.ListenAndServe())

}