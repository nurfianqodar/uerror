package main

import (
	"log"

	"github.com/go-playground/locales/en"
	"github.com/go-playground/locales/id"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	enTrans "github.com/go-playground/validator/v10/translations/en"
	idTrans "github.com/go-playground/validator/v10/translations/id"
)

var (
	enLoc = en.New()
	idLoc = id.New()
	uni   = ut.New(enLoc, idLoc)
)

func NewIDTrans(v *validator.Validate) ut.Translator {
	trans, found := uni.GetTranslator("id")
	if !found {
		log.Fatal("TRANSLATOR ERROR. id TRANSLATOR NOT FOUND")
	}
	if err := idTrans.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}
	return trans
}

func NewENTrans(v *validator.Validate) ut.Translator {
	trans, found := uni.GetTranslator("en")
	if !found {
		log.Fatal("TRANSLATOR ERROR. id TRANSLATOR NOT FOUND")
	}
	if err := enTrans.RegisterDefaultTranslations(v, trans); err != nil {
		log.Fatal(err)
	}
	return trans
}
