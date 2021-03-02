package tpl

func MakefileTemplate() []byte {
	return []byte(`Output := main

.PHONY: build clean

build:
	GO111MODULE=on go build -mod vendor -o $(Output) cmd/main.go

clean:
	@rm -f $(Output)
`)
}
