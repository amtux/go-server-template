all: deps clean build run-dev

deps:
	go get -v github.com/julienschmidt/httprouter
	go get -v github.com/sirupsen/logrus

clean:
	rm -rf bin/

build: deps
	go build -o bin/server

run-dev:
	./bin/server

.PHONY: \
	deps \
	clean \
	build \
	run-dev
    