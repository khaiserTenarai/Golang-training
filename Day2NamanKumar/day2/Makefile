# Day 2 build automation
# Run one task:   make run T=05-salary
# Run mini project: make mini
T ?= 15-employee-cli

.PHONY: all run mini build test fmt vet clean list

all: fmt vet test build          ## format, vet, test, build

run:                             ## run one task: make run T=09-sales
	go run ./cmd/$(T)

mini:                            ## run the Day 2 mini project
	go run ./cmd/15-employee-cli

build:                           ## build every task into ./bin
	@mkdir -p bin
	@for d in cmd/*/; do n=$$(basename $$d); go build -o bin/$$n ./$$d && echo "built bin/$$n"; done

test:                            ## unit tests with coverage
	go test -cover ./...

fmt:
	go fmt ./...

vet:
	go vet ./...

list:                            ## list all task programs
	@ls cmd

clean:
	rm -rf bin
