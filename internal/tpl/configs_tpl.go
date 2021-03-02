package tpl

func ConfigsTemplate() []byte {
	return []byte(`// Package configs contains the configuration of some environment variables.
/*
{{ .Copyright }}
{{ if .License.Header }}{{ .License.Header }}{{ end }}
*/
package configs

import (
	"os"
	"strconv"
)

const (
	serverPortEnv = "SERVER_PORT"
)

var (
    // ServerPort is the web server listening port.
	ServerPort int
)

func init() {
	setupConfigs()
}

func setupConfigs() {
	ServerPort = loadIntConfig(serverPortEnv, 8080)
}

func loadStringConfig(envCode, defaultValue string) string {
	v := os.Getenv(envCode)
	if v == "" {
		return defaultValue
	}
	return v
}

func loadIntConfig(envCode string, defaultValue int) int {
	v := os.Getenv(envCode)
	if v == "" {
		return defaultValue
	}

	val, err := strconv.Atoi(v)
	if err != nil {
		return defaultValue
	}
	return val
}
`)
}
