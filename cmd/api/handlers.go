package main

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/Duffney/reading-list/internal/data"
)

func (app *application) healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

	data := envelope{
		"status": "available",
		"system_info": map[string]string{
			"environment": app.config.env,
			"version":     version,
		},
	}

	if err := app.writeJSON(w, http.StatusOK, data, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}
func (app *application) multiplexer(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		app.getBookHandler(w, r)
	case "PUT":
		app.editBookHandler(w, r)
	case "DELETE":
		app.deleteBookHandler(w, r)
	default:
		app.methodNotAllowedResponse(w, r)
	}
}

func (app *application) bookHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		app.methodNotAllowedResponse(w, r)
		return
	}

	if r.Method == http.MethodPost {
		fmt.Fprintln(w, "Add a new book to the reading list")
	}

	if r.Method == http.MethodGet {
		fmt.Fprintln(w, "return reading list books")
	}
}

func (app *application) getBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("v1/books//"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	book := data.Book{
		ID:        idInt,
		CreatedAt: time.Now(),
		Title:     "Reclaim",
		Subtitle:  "Win the War Against Distraction and Rebuild the Linear Mind",
		Published: "2023-11",
		Pages:     125,
		Genres:    []string{"Nonfiction", "Productivity", "Self Help"},
		Rating:    4,
		Version:   1,
	}

	if err := app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil); err != nil {
		app.serverErrorResponse(w, r, err)
	}
}

func (app *application) editBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("v1/books//"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "Update record for bookID %d\n", idInt)
}

func (app *application) deleteBookHandler(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Path[len("v1/books//"):]
	idInt, err := strconv.ParseInt(id, 10, 64)
	if err != nil {
		app.badRequestResponse(w, r, err)
		return
	}
	fmt.Fprintf(w, "Delete record for bookID %d\n", idInt)
}
