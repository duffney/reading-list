package main

import (
	"fmt"
	"net/http"
	"strconv"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintln(w, "status: available")
	fmt.Fprintf(w, "environment: %s\n", app.config.env)
	fmt.Fprintf(w, "version: %s\n", version)
}

func (app *application) crBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Add a new book to the reading list")
	}

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "return reading list books")
	}

}

func (app *application) rudBooksHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPut && r.Method != http.MethodDelete {
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}

	id := r.URL.Path[len("v1/books//"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		http.Error(w, http.StatusText(400), http.StatusBadRequest)
		return
	}

	if r.Method == http.MethodGet {
		fmt.Fprintf(w, "Get record for bookID %d\n", idInt)
	}

	if r.Method == http.MethodPut {
		fmt.Fprintf(w, "Update record for bookID %d\n", idInt)
	}

	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, "Delete record for bookID %d\n", idInt)
	}
}
