package main

import (
	"log"
	"net/http"
)

func Handler(w http.ResponseWriter, r *http.Request) {
	log.Print("Hello world received a request.")
	
	// fmt.Fprintf(w, "Hello %s!\n", target)
}
