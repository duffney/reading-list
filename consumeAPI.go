package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Book struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     string   `json:"pages"`
	Genres    []string `json:"genres"`
	Version   int      `json:"version"`
}

type Envelope struct {
	Book Book `json:"book"`
}

func main() {

	resp, err := http.Get("http://localhost:4000/v1/books/13")
	if err != nil {
		fmt.Println(err)
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}

	// Unmarshal the JSON data into the envelope struct
	var envelope Envelope
	err = json.Unmarshal(body, &envelope)
	if err != nil {
		fmt.Println(err)
	}

	// Extract the book data from the envelope
	book := envelope.Book

	// Print the struct
	fmt.Println(book)
}
