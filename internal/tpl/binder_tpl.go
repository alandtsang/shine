package tpl

func BinderTemplate() []byte {
	return []byte(`// Package binder handles the binding of the http request parameters of echo.
/*
{{ .Copyright }}
{{ if .License.Header }}{{ .License.Header }}{{ end }}
*/
package binder

import (
    "fmt"

	"github.com/labstack/echo/v4"

	"{{ .PkgName }}/internal/validator"
)

// BindAndValidate is function to bind and validate data.
func BindAndValidate(c echo.Context, i interface{}) error {
	var err error

	if err = c.Bind(i); err != nil {
		fmt.Printf("%v\n", err)
		return err
	}

	if err = c.Validate(i); err != nil {
		return validator.Translate(err)
	}
	return nil
}
`)
}
