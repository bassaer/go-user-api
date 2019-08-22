.PHONY: build test clean

build:
	go build ./app/cmd/userapi/

test:
	go test -v -cover -race ./app/...

lint:
	go vet ./app/...
	golint -set_exit_status ./app/...

clean:
	rm -f userapi
	sudo rm -rf db/data
	sudo rm -rf log
