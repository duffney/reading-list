package main

import (
	"database/sql"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/Duffney/reading-list/internal/data"
	_ "github.com/lib/pq"
)

const version = "1.0.0"

type config struct {
	port int
	env  string
	db   struct {
		dsn string
	}
}

type application struct {
	config config //type embedding
	logger *log.Logger
	models data.Models
}

func main() {
	var cfg config

	flag.IntVar(&cfg.port, "port", 4000, "API server port")
	flag.StringVar(&cfg.env, "env", "dev", "Environemnt (dev|stage|prod")
	flag.StringVar(&cfg.db.dsn, "db-dsn", os.Getenv("READINGLIST_DB_DSN"), "PostgreSQL DSN")
	flag.Parse()

	logger := log.New(os.Stdout, "", log.Ldate|log.Ltime)

	db, err := sql.Open("postgres", cfg.db.dsn)
	if err != nil {
		logger.Fatal(err)
	}

	defer db.Close()

	err = db.Ping()
	if err != nil {
		logger.Fatal(err)
	}

	logger.Printf("database connection pool established")

	app := &application{
		config: cfg,
		logger: logger,
		models: data.NewModels(db),
	}
	addr := fmt.Sprintf(":%d", cfg.port)

	http.HandleFunc("/", app.notFoundResponse) // route for custom 404 reponse
	http.HandleFunc("/v1/healthcheck", app.healthcheckHandler)
	http.HandleFunc("/v1/books", app.bookHandler)
	http.HandleFunc("/v1/books/", app.multiplexer)

	logger.Printf("starting %s server on %s", cfg.env, addr)
	err = http.ListenAndServe(addr, nil)
	if err != nil {
		logger.Fatal(err)
	}
}
