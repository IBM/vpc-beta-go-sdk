# Makefile

all: build lint tidy

ci: build lint tidy test-unit

build:
    go build ./vpcbetav1

test-unit:
    go test `go list ./... | grep vpcbetav1` -v -tags=unit

test-integration:
    go test `go list ./... | grep vpcbetav1` -v -tags=integration -skipForMockTesting -testCount

test-examples:
    go test `go list ./... | grep vpcbetav1` -v -tags=examples

lint:
    golangci-lint --timeout=2m run

tidy:
    go mod tidy