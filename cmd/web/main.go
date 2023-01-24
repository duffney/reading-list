package main

import (
	"log"
	"net/http"
)

func main() {
	srv := &http.Server{
		Addr:    ":8080",
		Handler: routes(),
	}

	log.Print("Starting server on :8080")
	err := srv.ListenAndServe()
	log.Fatal(err)
}
