BUILDTAGS=

APP?=os-cli
USERSPACE?=serboox
RELEASE?=1.0.0
PROJECT?=github.com/${USERSPACE}/${APP}
GOOS?=linux
HTTP_PORT?=8080

REPO_INFO=$(shell git config --get remote.origin.url)

ifndef COMMIT
	COMMIT := git-$(shell git rev-parse --short HEAD)
endif

ifndef TEST_DIR
	TEST_DIR :=...
endif

clean:
	rm -f ${APP}

vendor: clean
	go get -u github.com/golang/dep/cmd/dep \
	&& dep ensure

build:
	GOOS=${GOOS} CGO_ENABLED=1 go build -a -v -installsuffix cgo \
		-ldflags "-s -w -X ${PROJECT}/version.Release=${RELEASE} -X ${PROJECT}/version.Commit=${COMMIT} -X ${PROJECT}/version.Repo=${REPO_INFO}" \
		-o ~/Documents/Linux_Shared/${APP}

fmt:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"gofmt -s -l {{.Dir}}"{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c

lint:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"golint {{.Dir}}/..."{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c

vet:
	@echo "+ $@"
	@go vet $(shell go list ${PROJECT}/... | grep -v vendor)

test:
	@echo "+ $@"
	@go test $(shell go list ${PROJECT}/... | grep -v vendor)

test-v:
	@echo "+ $@"
	@go test -v $(shell go list ${PROJECT}/... | grep -v -P '(vendor|version|tools)')

test-custom:
	@go test -v $(shell go list ${PROJECT}/... | grep ${TEST_DIR})		

test-race:
	@echo "+ $@"
	@go test -v -race $(shell go list ${PROJECT}/... | grep -v vendor)	

test-full: vendor fmt lint vet
	@echo "+ $@"
	@go test -v -race -tags "$(BUILDTAGS) cgo" $(shell go list ${PROJECT}/... | grep -v vendor)

cover:
	@echo "+ $@"
	@go list -f '{{if len .TestGoFiles}}"go test -coverprofile={{.Dir}}/.coverprofile {{.ImportPath}}"{{end}}' $(shell go list ${PROJECT}/... | grep -v vendor) | xargs -L 1 sh -c