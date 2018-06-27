#!/bin/make

GO_TAG:=$(shell /bin/sh -c 'eval `go tool dist env`; echo "$${GOOS}_$${GOARCH}"')
MACHINE:=$(shell uname -s)
GIT_TAG:=$(shell git rev-parse --short HEAD)
GOPATH:=$(shell go env GOPATH)
SOURCES:=$(shell find .. -name '*.go')

.PHONY: all deps update fmt test check doc gen

all: dualcontour-go

dualcontour-go: $(SOURCES)
	$(GOPATH)/bin/goimports -w -l ..
	go build -v -gcflags "-N -l"

clean:
	go clean

deps:
	go get -v .

update:
	go get -u .

fmt:
	go fmt ./...
	$(GOPATH)/bin/goimports -w -l .

test:
	go test ./...

gen:
	@if [ ! -f $(GOPATH)/bin/stringer ]; then go get golang.org/x/tools/cmd/stringer; fi
	go generate ./...

check:
	@if [ ! -f $(GOPATH)/bin/gometalinter ]; then go get github.com/alecthomas/gometalinter; fi
	$(GOPATH)/bin/gometalinter ./...

doc:
	@if [ ! -f $(GOPATH)/bin/godoc ]; then go get golang.org/x/tools/cmd/godoc; fi
	$(GOPATH)/bin/godoc -v -http=:6060 -index -play
