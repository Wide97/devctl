package main

import (
	"devctl/internal/mockserver"
	"log"
	"net/http"
)

func main() {
	mux := mockserver.NewMux()

	addr := ":8080"
	log.Printf("mock server listening on %s", addr)
	if err := http.ListenAndServe(addr, mux); err != nil {
		log.Fatal(err)
	}
}
