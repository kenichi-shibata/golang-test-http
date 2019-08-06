GIT_HASH=`git rev-parse --short HEAD`

.PHONY: run
run:
	go run main.go

.PHONY: build
build:
	docker build -t quay.io/kenichi_shibata/golang-http-test:${GIT_HASH} .

.PHONY: test
test:
	go test -v 