package data

import (
	"time"
)

type Book struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"created_at"`
	Title         string    `json:"title"`
	Subtitle      string    `json:"year"`
	PublishedDate time.Time `json:"published_date"` //year, month, day := fromDate.Date()
	Pages         int       `json:"pages"`
	Genres        []string  `json:"genres"`
	Rating        int       `json:"rating"`
	Version       int32     `json:"version"`
}

// TODO Create a custom type for PublishedDate like Runtime in Let's Go Further
