package data

import (
	"strconv"
	"time"
)

type Published string

const (
	layoutISO = "2006-01"
	layoutUS  = "January, 2006"
)

func (p Published) MarshalJSON() ([]byte, error) {

	date := string(p)

	t, _ := time.Parse(layoutISO, date)

	quotedJSON := strconv.Quote(t.Format(layoutUS))

	return []byte(quotedJSON), nil
}

/*
Take a string, 3/15/2022
Return a date, March 15, 2022
TODO add date validator to published input, needs to be year-month

Resources:
- https://yourbasic.org/golang/format-parse-string-time-date-example/
*/
