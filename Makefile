Output := shine

.PHONY: build clean

build:
	GO111MODULE=on go build -mod vendor -o $(Output) cmd/shine/shine.go

clean:
	@rm -rf $(Output)