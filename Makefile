.PHONY: gotool build clean

BINARY="main"

all: gotool	build

gotool:
	@go fmt ./
	@go vet ./

build:
	@go build -o $(BINARY)

run:
	@go run main.go

clean: 
	@if [ -f $(BINARY) ] ; then rm $(BINARY) ; fi
