APPLICATION = nm-registry
ORGANIZATION = skunkworxs

TTY = $(shell if [ -t 0 ]; then echo "-ti"; fi)

DOCKER_DIR = /go/src/github.com/${ORGANIZATION}/${APPLICATION}
DOCKER_IMAGE = ${ORGANIZATION}/golang:1.7-wheezy
DOCKER_CMD = docker run --rm -v ${PWD}:${DOCKER_DIR} ${TTY} -w ${DOCKER_DIR} ${DOCKER_IMAGE}

MAIN_PATH = cmd/${APPLICATION}

# variable definitions
SHELL = /bin/bash
COMMITHASH = $(shell git describe --tags --always --dirty)
BUILDDATE = $(shell date -u)
BUILDER = $(shell echo "`git config user.name` <`git config user.email`>")
GOVERSION = $(shell go version)
OSARCH = $(shell uname -sm)
RELEASEVERSION = 0.1



#Flags to pass to main.go
LDFLAGS = -ldflags "-X 'main.binaryName=${APPLICATION}' \
		    -X 'main.buildDate=${BUILDDATE}' \
		    -X 'main.buildUser=${BUILDER}' \
		    -X 'main.commitHash=${COMMITHASH}' \
		    -X 'main.goVersion=${GOVERSION}' \
		    -X 'main.osArch=${OSARCH}' \
		    -X 'main.releaseVersion=${RELEASEVERSION}' "

#Some tests need to run for 5+ seconds, which trips Ginkgo Slow Test warning
SLOWTEST = 10

.PHONY: shell deps deps-local build build-local lint lint-local test test-local release

default: deps test build

shell:
	@${DOCKER_CMD} /bin/bash

consul-shell:
	@docker run --rm -ti --net nmregistry_default -v ${PWD}:${DOCKER_DIR} -w ${DOCKER_DIR} ${DOCKER_IMAGE} /bin/bash

clean:
	@${DOCKER_CMD} make clean-local

clean-local:
	@rm -rf bin vendor

deps:
	@${DOCKER_CMD} make deps-local

deps-local:
	@if ! [ -f glide.lock ]; then glide init --non-interactive; fi
	@glide install

build:
	@${DOCKER_CMD} make build-local

build-local: lint-local
	@go build -o bin/${APPLICATION} ${LDFLAGS} ${MAIN_PATH}/*.go
	@go build -o cmd/ssdpspoofer/bin/ssdpspoofer cmd/ssdpspoofer/*.go

lint:
	@${DOCKER_CMD} make lint-local

lint-local:
	@gometalinter --vendor --fast --disable=dupl --disable=gotype --skip=grpc ./...

test:
	@${DOCKER_CMD} make test-local

test-local: lint-local
	@ginkgo -race -trace -cover -randomizeAllSpecs --slowSpecThreshold=${SLOWTEST}

release: deps build
	@docker build -t ${ORGANIZATION}/${APPLICATION} .
	@docker build -t ${ORGANIZATION}/ssdpspoofer cmd/ssdpspoofer/

run: release
	@docker-compose up --force-recreate
