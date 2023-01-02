package data

import (
	"time"
)

type Book struct {
	ID            int64
	CreatedAt     time.Time
	Title         string
	Subtitle      string
	PublishedDate time.Time //year, month, day := fromDate.Date()
	Pages         int
	Genres        []string
	Rating        int
	Version       int32
}

// TODO Create a custom type for PublishedDate like Runtime in Let's Go Further
