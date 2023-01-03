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

	err := app.writeJSON(w, http.StatusOK, data, nil)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encounterd a problem and could not process your request", http.StatusInternalServerError)
	}
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
		book := data.Book{
			ID:            idInt,
			CreatedAt:     time.Now(),
			Title:         "Reclaim",
			Subtitle:      "Win the War Against Distraction and Rebuild the Linear Mind",
			PublishedDate: time.Date(2023, time.November, 12, 25, 0, 0, 0, time.UTC),
			Pages:         125,
			Genres:        []string{"Nonfiction", "Productivity", "Self Help"},
			Rating:        4,
			Version:       1,
		}

		err := app.writeJSON(w, http.StatusOK, envelope{"book": book}, nil)
		if err != nil {
			app.logger.Print(err)
			http.Error(w, "The server encounterd a problem and could not process your request", http.StatusInternalServerError)
		}
	}

	if r.Method == http.MethodPut {
		fmt.Fprintf(w, "Update record for bookID %d\n", idInt)
	}

	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, "Delete record for bookID %d\n", idInt)
	}
}
