.PHONY: build clean test run install

build:
	CGO_ENABLED=0 GOOS=linux go build -ldflags="-s -w" -o smh .

clean:
	rm -f smh

test:
	go test ./...

run: build
	./smh --help

install: build
	sudo cp smh /usr/local/bin/