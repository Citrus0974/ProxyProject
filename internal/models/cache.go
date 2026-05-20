package models

import (
	"net/http"
	"time"
)

type Entry struct {
	Status int
	Header http.Header
	Body   []byte

	ExpireAt time.Time
}
