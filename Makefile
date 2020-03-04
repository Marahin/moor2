#!make
include .env
export

install-dependencies:
		go mod download
		npm i -g now@latest
		go get github.com/wadey/gocovmerge

local-run:
		now dev

create-coverage-directory:
		mkdir -p .coverage

test-html-report:
		go tool cover -html=.coverage/coverage.out

test: create-coverage-directory
		go test -v ./... --tags=unit -coverprofile=.coverage/coverage.out -covermode=atomic
