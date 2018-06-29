EAMEserser     := memcache
VERSION  := HEAD
#VERSION  := $(shell git describe --tags --abbrev=0
REVISION := $(shell git rev-parse --short HEAD)
BRANCH   := $(shell git rev-parse --abbrev-ref HEAD)
SRC      := $(shell find . -type f -name '*.go' -not -path "./vendor/*")
TARGET   := ./memcache/

LDFLAGS  := -ldflags="-s -w -X \"main.Name=$(NAME)\" -X \"main.Version=$(VERSION)\" -X \"main.Revision=$(REVISION)\" -extldflags \"-static\""

.PHONY: check clean fmt linux release simplify test

bin/$(NAME): $(SRC)
	CGO_ENABLED=0 go build -a -tags netgo -installsuffix netgo $(LDFLAGS) -o bin/$(NAME) $(TARGET)

check:
	@for d in $$(go list ./... | grep -v /vendor/); do golint $${d}; done
	@go tool vet $(SRC)

clean:
	@go clean -testcache
	rm -rf bin/*
	rm -rf dist/*

fmt:
	@gofmt -l -w $(SRC)

linux:
	for os in linux; do \
		for arch in amd64; do \
		  GOOS=$$os GOARCH=$$arch \
			CGO_ENABLED=0  go build -a -tags "netgo" -installsuffix netgo $(LDFLAGS) -o dist/$$os-$$arch/$(NAME) $(TARGET); \
		done; \
  done;

release:
	@goreleaser --rm-dist

simplify:
	@gofmt -s -l -w $(SRC)

test:
	@go test -cover ./...
