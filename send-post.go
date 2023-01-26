package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
)

type Book struct {
	Title     string   `json:"title"`
	Pages     int      `json:"pages"`
	Published int      `json:"published"`
	Genres    []string `json:"genres"`
}

func main() {
	book := Book{
		Title:     "The Great Gatsby",
		Pages:     180,
		Published: 1925,
		Genres:    []string{"Fiction", "Drama"},
	}

	jsonData, _ := json.Marshal(book)
	fmt.Println(string(jsonData))

	req, _ := http.NewRequest("POST", "http://localhost:4000/v1/books", bytes.NewBuffer(jsonData))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	fmt.Println(resp.Status)
}
