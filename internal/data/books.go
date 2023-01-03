package data

import (
	"time"
)

type Book struct {
	ID        int64     `json:"id"`
	CreatedAt time.Time `json:"-"`
	Title     string    `json:"title"`
	Subtitle  string    `json:"subtitle"`
	Published Published `json:"published,omitempty"`    //year, month, day := fromDate.Date()
	Pages     int       `json:"pages,omitempty,string"` //changed int to string in JSON
	Genres    []string  `json:"genres,omitempty"`
	Rating    int       `json:"rating,omitempty"`
	Version   int32     `json:"version"`
}

// TODO Rename as models, move to api?
