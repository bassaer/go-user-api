.PHONY: build test clean

build:
	go build ./cmd/userapi/

test:
	go test -v -cover -race ./...

lint:
	go vet ./...
	golint -set_exit_status ./...

clean:
	rm -f userapi
	sudo rm -rf db/data
