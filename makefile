export GOPATH := $(shell pwd)

build:
	@echo "--> Building..."
	@mkdir -p bin/
	go build -v -o bin/vdf src/cmd/main.go
	@chmod 755 bin/*

clean:
	@echo "--> Cleaning..."
	@go clean
	@rm -f bin/*

test:
	@echo "--> Testing..."
	@$(MAKE) testsqrt

testsqrt:
	go test -v sqrt

check:
	go get -v github.com/golangci/golangci-lint/cmd/golangci-lint
	bin/golangci-lint run -D errcheck src/sqrt/...

.PHONY: build clean install fmt test coverage
