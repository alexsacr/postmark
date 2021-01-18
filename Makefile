.PHONY: default test integration setup

default: test

test:
	go test -v -cover ./...
	go vet ./...
	go list ./... | grep -v /vendor/ | xargs -n 1 golint
	errcheck ./...

integration:
	INTEGRATION=true go test -v

setup:
	go get github.com/golang/lint/golint
	go get github.com/kisielk/errcheck
