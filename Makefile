Output := shine

.PHONY: build clean

build:
	GO111MODULE=on go build -mod vendor -o $(Output) cmd/main.go

clean:
	@rm -f $(Output)