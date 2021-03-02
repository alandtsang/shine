package tpl

func MainTemplate() []byte {
	return []byte(`/*
{{ .Copyright }}
{{ if .License.Header }}{{ .License.Header }}{{ end }}
*/
package main

import (
	"fmt"
	"log"

	"{{ .PkgName }}/configs"
	"{{ .PkgName }}/internal/router"
)

func main() {
	addr := fmt.Sprintf(":%d", configs.ServerPort)

	r := router.New()
	log.Fatal(r.ListenAndServe(addr))
}
`)
}
