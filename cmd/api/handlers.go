package main

import (
	"encoding/json"
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

	data := map[string](string){
		"status":      "available",
		"environment": app.config.env,
		"version":     version,
	}

	js, err := json.Marshal(data)
	if err != nil {
		app.logger.Print(err)
		http.Error(w, "The server encournted a problem and could not process your request", http.StatusInternalServerError)
		return
	}

	js = append(js, '\n')

	w.Header().Set("Content-Type", "application/json")
	w.Write(js)
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

		js, err := json.Marshal(book)
		if err != nil {
			app.logger.Print(err)
			http.Error(w, "The server encournted a problem and could not process your request", http.StatusInternalServerError)
			return
		}

		js = append(js, '\n')

		w.Header().Set("Content-Type", "application/json")
		w.Write(js)
	}

	if r.Method == http.MethodPut {
		fmt.Fprintf(w, "Update record for bookID %d\n", idInt)
	}

	if r.Method == http.MethodDelete {
		fmt.Fprintf(w, "Delete record for bookID %d\n", idInt)
	}
}
