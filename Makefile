SHELL := /bin/bash

TARGET := $(shell echo $${PWD\#\#*/})
.DEFAULT_GOAL: $(TARGET)

VERSION_FILE = ./VERSION
VERSION = `cat $(VERSION_FILE)`
SUFFIX = "-dev"
CODEPATH = `go list -m`

SRC = $(shell find . -type f -name '*.go' -not -path "./vendor/*")

export GO111MODULE = on

.PHONY: all build clean install uninstall fmt simplify check run

all: check install

$(TARGET): $(SRC)
	packr build -ldflags "-X ${CODEPATH}/cmd.version=${VERSION}${SUFFIX}" -o ${TARGET}

build: $(TARGET)
	@true

clean:
	@rm -f $(TARGET)

install:
	@packr install

uninstall: clean
	@rm -f $$(which ${TARGET})

fmt:
	@gofmt -l -w $(SRC)

simplify:
	@gofmt -s -l -w $(SRC)

check:
	@go get -u golang.org/x/lint/golint
	@test -z $(shell gofmt -l ./ | tee /dev/stderr) || echo "[WARN] Fix formatting issues with 'make fmt'"
	@for d in $$(go list ./...); do golint $${d}; done
	@go vet

packr:
	@go get -u github.com/gobuffalo/packr/...

run: install
	@$(TARGET)
