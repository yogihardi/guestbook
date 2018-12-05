GB_PKG_VERSION?=0.0.0
COMMIT=`git rev-parse --short HEAD`

build: build_binary

install-mac:
	@GOOS=darwin GOARCH=amd64 \
	go install -v --ldflags "-w \
	-X github.com/yogihardi/guestbook/version.Version=$(GB_PKG_VERSION) \
	-X github.com/yogihardi/guestbook/version.GitCommit=$(COMMIT)" .

install-linux:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 \
	go install -v --ldflags "-w \
	-X github.com/yogihardi/guestbook/version.Version=$(GB_PKG_VERSION) \
	-X github.com/yogihardi/guestbook/version.GitCommit=$(COMMIT)" .

build_binary:
	@CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o guestbook -a --ldflags "-w \
	-X github.com/yogihardi/guestbook/version.Version=$(GB_PKG_VERSION) \
	-X github.com/yogihardi/guestbook/version.GitCommit=$(COMMIT)" .

test:
	@go test -v $(shell go list ./... | grep -v /vendor/)

vet:
	@go vet -v $(shell go list ./... | grep -v /vendor/)

clean:
	@rm -rf build
	@rm -rf guestbook*

.PHONY: test vet build build_binary clean
