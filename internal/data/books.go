package data

import (
	"time"
)

type Book struct {
	ID            int64     `json:"id"`
	CreatedAt     time.Time `json:"-"`
	Title         string    `json:"title"`
	Subtitle      string    `json:"year"`
	PublishedDate time.Time `json:"published_date,omitempty"` //year, month, day := fromDate.Date()
	Pages         int       `json:"pages,omitempty"`
	Genres        []string  `json:"genres,omitempty"`
	Rating        int       `json:"rating,omitempty"`
	Version       int32     `json:"version"`
}

// TODO Create a custom type for PublishedDate like Runtime in Let's Go Further
