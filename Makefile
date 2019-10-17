APP?=os-cli
USERSPACE?=serboox
RELEASE?=1.0.0
PROJECT?=github.com/${USERSPACE}/${APP}
BUILD_TIME?=$(shell date -u '+%Y-%m-%d_%H:%M:%S')
GOOS?=linux
REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
	COMMIT := git-$(shell git rev-parse --short HEAD)
endif

ifndef TEST_DIR
	TEST_DIR :=...
endif

# Default target executed when no arguments are given to make.
default_target: run-r

clean:
	rm -f ${APP}

build: clean golangci-lint
	@echo "+ $@"
	GOFLAGS=-mod=vendor CGO_ENABLED=1 go build -v -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/src/version.Release=${RELEASE} -X ${PROJECT}/src/version.Commit=${COMMIT} -X ${PROJECT}/src/version.Repository=${REPO_INFO} -X ${PROJECT}/src/version.BuildTime=${BUILD_TIME}"

golangci-lint-base:
	@echo "+ $@"
	GO111MODULE=on golangci-lint run -c .golangci-simple.yml ./...

golangci-lint:
	@echo "+ $@"
	GO111MODULE=on golangci-lint run ./...

docker-build:
	@echo "+ $@"
	docker build -t ${APP}:latest . --no-cache --force-rm

test:
	@echo "+ $@"
	@go test $(shell go list ${PROJECT}/... | grep -v vendor)

test-v:
	@echo "+ $@"
	@go test -v $(shell go list ${PROJECT}/... | grep -v -P '(vendor|version|tools)')

cover:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c
