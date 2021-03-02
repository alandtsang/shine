package tpl

func ValidatorTemplate() []byte {
	return []byte(`// Package validator handles the verification of echo http request parameters.
/*
{{ .Copyright }}
{{ if .License.Header }}{{ .License.Header }}{{ end }}
*/
package validator

import (
	"fmt"

	"github.com/go-playground/locales/en"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	en_translations "github.com/go-playground/validator/v10/translations/en"
)

var trans ut.Translator

// Validator is structure to validate http request parameters.
type Validator struct {
	validator *validator.Validate
}

// New return validator instance.
func New() *Validator {
	translator := en.New()
	uni := ut.New(translator, translator)
	trans, _ = uni.GetTranslator("en")

	validate := validator.New()
	_ = en_translations.RegisterDefaultTranslations(validate, trans)

	return &Validator{validator: validate}
}

// Validate is function to validate http request parameters.
func (v *Validator) Validate(i interface{}) error {
	return v.validator.Struct(i)
}

// Translate converts the error information of the validator to the default error information.
func Translate(err error) error {
	for _, e := range err.(validator.ValidationErrors) {
		fmt.Printf("%v\n", err)
		return fmt.Errorf(e.Translate(trans))
	}
	return nil
}
`)
}
