	GOBIN=$(shell pwd)/bin
	IMAGE=saas-interview-challenge1
    
    all: deps test build
    build: 
			GOBIN=$(GOBIN) go install ./...
    test:
			go test -v ./...
    clean: 
			go clean
			rm -r $(GOBIN)
	docker:
			docker build -t (IMAGE) .
    deps:
			go get github.com/gomodule/redigo/redis
			go get github.com/gorilla/mux