package tpl

func GoModTemplate() []byte {
	return []byte(`module {{ .PkgName }}

go 1.14

require (
	github.com/go-playground/locales v0.13.0
	github.com/go-playground/universal-translator v0.17.0
	github.com/go-playground/validator/v10 v10.4.1
	github.com/labstack/echo/v4 v4.1.17
)
`)
}

func GoSumTemplate() []byte {
	return []byte(`github.com/go-playground/locales v0.13.0 h1:HyWk6mgj5qFqCT5fjGBuRArbVDfE4hi8+e8ceBS/t7Q=
github.com/go-playground/locales v0.13.0/go.mod h1:taPMhCMXrRLJO55olJkUXHZBHCxTMfnGwq/HNwmWNS8=
github.com/go-playground/universal-translator v0.17.0 h1:icxd5fm+REJzpZx7ZfpaD876Lmtgy7VtROAbHHXk8no=
github.com/go-playground/universal-translator v0.17.0/go.mod h1:UkSxE5sNxxRwHyU+Scu5vgOQjsIJAF8j9muTVoKLVtA=
github.com/go-playground/validator/v10 v10.4.1 h1:pH2c5ADXtd66mxoE0Zm9SUhxE20r7aM3F26W0hOn+GE=
github.com/go-playground/validator/v10 v10.4.1/go.mod h1:nlOn6nFhuKACm19sB/8EGNn9GlaMV7XkbRSipzJ0Ii4=
`)
}
