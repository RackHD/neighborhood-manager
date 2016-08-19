ORGANIZATION = RackHD
PROJECT = neighborhood-manager

TTY = $(shell if [ -t 0 ]; then echo "-ti"; fi)

PROJECT_DIR = /go/src/github.com/${ORGANIZATION}/${PROJECT}
DOCKER_DIR = ${PROJECT_DIR}
DOCKER_IMAGE = rackhd/golang:1.7.0-wheezy
DOCKER_CMD = docker run --rm -v ${PWD}:${PROJECT_DIR} ${TTY} -w ${DOCKER_DIR} ${DOCKER_IMAGE}

noop:
	@echo Neighborhood Manager

coveralls:
	@go get github.com/mattn/goveralls
	@go get github.com/modocache/gover
	@go get golang.org/x/tools/cmd/cover
	@gover
	@goveralls -coverprofile=gover.coverprofile -service=travis-ci
