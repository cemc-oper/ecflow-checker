.PHONY: ecflow_checker

VERSION := $(shell cat VERSION)
BUILD_TIME := $(shell date --utc --rfc-3339 ns 2> /dev/null | sed -e 's/ /T/')
GIT_COMMIT := $(shell git rev-parse --short HEAD 2> /dev/null || true)
PROJECT_URL := "github.com/perillaroc/ecflow-checker"

ecflow_checker:
	go build \
		-ldflags "-X \"${PROJECT_URL}/cli/cmd.Version=${VERSION}\" \
		-X \"${PROJECT_URL}/cli/cmd.BuildTime=${BUILD_TIME}\" \
		-X \"${PROJECT_URL}/cli/cmd.GitCommit=${GIT_COMMIT}\" " \
		-o bin/ecflow_checker \
		cli/main.go