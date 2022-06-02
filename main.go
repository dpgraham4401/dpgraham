package main

import (
	"log"
	"net/http"
)

func main() {

	// assets
	fs := http.FileServer(http.Dir("assets"))
	mux := http.NewServeMux()
	mux.Handle("/assets/", http.StripPrefix("/assets/", fs))

	// Routing
	mux.HandleFunc("/", homeHandler)
	mux.HandleFunc("/blog/", blogHandler)

	log.Fatal(http.ListenAndServe(":8080", mux))
}
