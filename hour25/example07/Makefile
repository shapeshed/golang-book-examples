.DEFAULT_GOAL := test
GOFILES_NOVENDOR = $(shell find . -type f -name '*.go' -not -path "./vendor/*")
SHELL=/bin/bash

test:
		@go test ${GOFILES_NOVENDOR}


