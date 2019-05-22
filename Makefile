.PHONY: ecflow_checker

VERSION := v0.0.1
BUILD_TIME := $(shell date --utc --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2> /dev/null || true)

ecflow_checker:
	go build \
		-ldflags "-X \"github.com/perillaroc/ecflow-client-go/ecflow_checker/cmd.Version=${VERSION}\" \
		-X \"github.com/perillaroc/ecflow-client-go/ecflow_checker/cmd.BuildTime=${BUILD_TIME}\" \
		-X \"github.com/perillaroc/ecflow-client-go/ecflow_checker/cmd.GitCommit=${GIT_COMMIT}\" " \
		-o bin/ecflow_checker \
		ecflow_checker/main.go