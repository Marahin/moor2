#!make
include .env
export

install-dependencies:
		go mod download
		go get github.com/wadey/gocovmerge

install-now:
	npm i -g now@latest

local-run: install-now
		now dev

create-coverage-directory:
		mkdir -p .coverage

test-html-report:
		go tool cover -html=.coverage/coverage.out

test: create-coverage-directory
		go test -v ./... --tags=unit -coverprofile=.coverage/coverage.out -covermode=atomic
