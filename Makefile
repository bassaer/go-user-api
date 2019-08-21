.PHONY: build test clean

build:
	go build ./cmd/userapi/

test:
	go test -v -cover -race ./...

clean:
	rm -f userapi
