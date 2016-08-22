ORGANIZATION = RackHD
PROJECT = neighborhood-manager
REGISTRY = registry
PROXY = rackhd


TTY = $(shell if [ -t 0 ]; then echo "-ti"; fi)

DOCKER_DIR = /go/src/github.com/${ORGANIZATION}/${PROJECT}
DOCKER_IMAGE = rackhd/golang:1.7.0-wheezy
DOCKER_CMD = docker run --rm -v ${PWD}:${DOCKER_DIR} ${TTY} -w ${DOCKER_DIR} ${DOCKER_IMAGE}


# variable definitions
COMMITHASH = $(shell git describe --tags --always --dirty)
BUILDDATE = $(shell date -u)
BUILDER = $(shell echo "`git config user.name` <`git config user.email`>")
GOVERSION = $(shell go version)
OSARCH = $(shell uname -sm)
RELEASEVERSION = 0.1



#Flags to pass to main.go
REGFLAGS = -ldflags "-X 'main.binaryName=${REGISTRY}' \
		    -X 'main.buildDate=${BUILDDATE}' \
		    -X 'main.buildUser=${BUILDER}' \
		    -X 'main.commitHash=${COMMITHASH}' \
		    -X 'main.goVersion=${GOVERSION}' \
		    -X 'main.osArch=${OSARCH}' \
		    -X 'main.releaseVersion=${RELEASEVERSION}' "

PROXYFLAGS = -ldflags "-X 'main.binaryName=${PROXY}' \
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

coveralls:
	@go get github.com/mattn/goveralls
	@go get github.com/modocache/gover
	@go get golang.org/x/tools/cmd/cover
	@gover
	@goveralls -coverprofile=gover.coverprofile -service=travis-ci

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
	@glide install --strip-vcs --strip-vendor

build:
	@build-reg
	@build-proxy

build-proxy:
	@${DOCKER_CMD} make build-proxy-local

build-proxy-local:
	@go build -o bin/${PROXY} ${PROXYFLAGS} rackhd/cmd/rackhd/*.go
	@go build -o bin/endpoint rackhd/cmd/utils/*.go

build-reg:
	@${DOCKER_CMD} make build-reg-local

build-reg-local: lint-local
	@go build -o bin/${REGISTRY} ${LDFLAGS} registry/cmd/registry/*.go
	@go build -o registry/cmd/ssdpspoofer/bin/ssdpspoofer registry/cmd/ssdpspoofer/*.go

lint:
	@${DOCKER_CMD} make lint-local

lint-local:
	@gometalinter --vendor --fast --disable=dupl --disable=gotype --skip=grpc ./...

test:
	@${DOCKER_CMD} make test-local

test-local: lint-local
	@ginkgo -r -race -trace -cover -randomizeAllSpecs --slowSpecThreshold=${SLOWTEST}

release: deps build
	@docker build -t rackhd/${REGISTRY} registry
	@docker build -t rackhd/ssdpspoofer registry/cmd/ssdpspoofer/
	@docker build -t rackhd/${PROXY} rackhd
	@docker build -t rackhd/endpoint rackhd/cmd/utils/


run: release
	@docker-compose up --force-recreate
