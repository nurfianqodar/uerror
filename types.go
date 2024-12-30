package uerror

import (
	"log"
	"net/http"

	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
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

func FromValidator(ve validator.ValidationErrors, trans ut.Translator) UError {
	errs := make([]string, len(ve))
	for _, e := range ve {
		errs = append(errs, e.Translate(trans))
	}
	return &uerror{
		HTTPCode: http.StatusBadRequest,
		Message:  "validation error",
		Errors:   errs,
	}
}
