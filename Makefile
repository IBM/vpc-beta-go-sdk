# Makefile

GO=go
LINT=golangci-lint
FORMATTER=goimports

all: tidy build test-unit lint

ci: tidy build test-unit lint

build:
	${GO} build ./vpcbetav1

test-unit:
	${GO} test `go list ./... | grep vpcbetav1` -v -tags=unit

test-integration:
	${GO} test `go list ./... | grep vpcbetav1` -v -tags=integration -skipForMockTesting -testCount

test-examples:
	${GO} test `go list ./... | grep vpcbetav1` -v -tags=examples

lint:
	${LINT} run
	DIFF=$$(${FORMATTER} -d vpcbetav1); if [ -n "$$DIFF" ]; then printf "\n$$DIFF\n" && exit 1; fi

tidy:
	${GO} mod tidy
