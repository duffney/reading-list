package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func home(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	// TODO call model to get data
	bookModel := &Book{}
	books, err := bookModel.GetAll()
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/home.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal Server Error", 500)
		return
	}

	// TODO convert data to template data
	data := &templateData{
		Books: books,
	}

	err = ts.ExecuteTemplate(w, "base", data)
	if err != nil {
		log.Print(err.Error())
		http.Error(w, "Internal server error", 500)
		return
	}
}

func showBook(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(r.URL.Query().Get("id"))
	if err != nil || id < 1 {
		http.NotFound(w, r)
		return
	}

	bookModel := &Book{}
	book, err := bookModel.Get(id)
	if err != nil {
		log.Println(err)
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/book.html",
	}

	funcs := template.FuncMap{"join": strings.Join}

	ts, err := template.New("showBook").Funcs(funcs).ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}

	err = ts.ExecuteTemplate(w, "base", book)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func createBook(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case "GET":
		createBookForm(w, r)
	case "POST":
		createBookPost(w, r)
	default:
		http.Error(w, http.StatusText(405), http.StatusMethodNotAllowed)
		return
	}
}

func createBookForm(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.NotFound(w, r)
		return
	}

	files := []string{
		"./ui/html/base.html",
		"./ui/html/partials/nav.html",
		"./ui/html/pages/create.html",
	}

	ts, err := template.ParseFiles(files...)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
	err = ts.ExecuteTemplate(w, "base", nil)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", 500)
		return
	}
}

func createBookPost(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.NotFound(w, r)
		return
	}

	err := r.ParseForm()
	if err != nil {
		log.Println(err)
		http.Error(w, "Bad request", http.StatusBadRequest)
		return
	}

	title := r.PostForm.Get("title")
	published, _ := strconv.Atoi(r.PostForm.Get("published"))
	pages, _ := strconv.Atoi(r.PostForm.Get("pages"))
	genres := r.Form["genres[]"]

	book := struct {
		Title     string   `json:"title"`
		Pages     int      `json:"pages"`
		Published int      `json:"published"`
		Genres    []string `json:"genres"`
	}{
		Title:     title,
		Pages:     pages,
		Published: published,
		Genres:    genres,
	}
	data, _ := json.Marshal(book)
	//TODO add API endpoint to appConfig
	req, _ := http.NewRequest("POST", "http://localhost:4000/v1/books", bytes.NewBuffer(data))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println(resp.Status)

	// http.Redirect(w, r, fmt.Sprintf("/book?id=%d", book.ID), http.StatusSeeOther)
	http.Redirect(w, r, "/", http.StatusSeeOther)
}
