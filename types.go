package main

import (
	"log"
	"net/http"
)

type UError interface {
	error
	Send(http.ResponseWriter)
}

type uerror struct {
	HTTPCode int      `json:"-"`
	Message  string   `json:"message"`
	Errors   []string `json:"errors"`
}

// # Genereate new uerror instance
//
// Parameters `code` for http status code, `message` for error message
// and `e` for multiple error message (optional)
func New(code int, message string, e ...string) UError {
	if e != nil {
		log.Println(e)
	}
	return &uerror{
		HTTPCode: code,
		Message:  message,
		Errors:   e,
	}
}
