package main

import (
	"encoding/json"
	"net/http"
)

func (e *uerror) Error() string {
	return e.Message
}

// Write uerror into HTTP response
func (e *uerror) Send(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.HTTPCode)
	json.NewEncoder(w).Encode(e)
}
