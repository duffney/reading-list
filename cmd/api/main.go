package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environemnt (dev|stage|prod")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	addr := fmt.Sprintf(":%d", cfg.port)

	http.HandleFunc("/v1/healthcheck", healthcheckHandler)

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err := http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Fatal(err)
	}
}
func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "status: available")
	fmt.Fprintln(w, "environment: development")
	fmt.Fprintf(w, "version: %s\n", version)
}
