package models

import (
	"errors"
	"time"
)

var ErrRecord = errors.New("models: mo matching record found")

type Snippet struct {
	ID      int
	Title   string
	Content string
	Created time.Time
	Expires time.Time
}
