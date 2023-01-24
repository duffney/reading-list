package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
)

type Book struct {
	ID        int      `json:"id"`
	Title     string   `json:"title"`
	Published int      `json:"published"`
	Pages     string   `json:"pages"`
	Genres    []string `json:"genres"`
	Version   int      `json:"version"`
}

func main() {
	jsonStr := []byte(`{
        "book": {
                "id": 13,
                "title": "Make It Stick",
                "published": 2014,
                "pages": "336",
                "genres": [
                        "non-fiction",
                        "education"
                ],
                "version": 1
        }
}`)

	// Create a new decoder
	dec := json.NewDecoder(bytes.NewBuffer(jsonStr))

	// Decode the JSON data into an intermediate variable
	var bookData map[string]json.RawMessage
	for {
		if err := dec.Decode(&bookData); err == io.EOF {
			break
		} else if err != nil {
			fmt.Println(err)
		}
	}

	// Unmarshal the book data into the struct
	var book Book
	err := json.Unmarshal(bookData["book"], &book)
	if err != nil {
		fmt.Println(err)
	}

	// Print the struct
	fmt.Println(book)
}
