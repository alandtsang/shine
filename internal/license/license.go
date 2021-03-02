package license

import (
	"fmt"
	"os"
	"text/template"
	"time"
)

const (
	Apache2License = "apache2"
)

type License struct {
	Name   string
	Header string
	Text   string
}

func New(name string) *License {
	if name != "apache2" {
		return nil
	}

	return initApache2()
}

func initApache2() *License {
	return &License{
		Name:   "apache2",
		Header: apache2Header,
		Text:   apache2Text,
	}
}

func CreateLicenseFile(path, author string) error {
	data := map[string]interface{}{
		"copyright": CreateCopyright(author),
	}

	licenseFile, err := os.Create(fmt.Sprintf("%s/LICENSE", path))
	if err != nil {
		return err
	}
	defer licenseFile.Close()

	licenseTemplate := template.Must(template.New("license").Parse(apache2Text))
	return licenseTemplate.Execute(licenseFile, data)
}

func CreateCopyright(author string) string {
	year := time.Now().Format("2006")
	return "Copyright © " + year + " " + author
}
