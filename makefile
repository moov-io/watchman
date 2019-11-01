PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+(-[a-zA-Z0-9]*)?)' version.go)

.PHONY: build build-server build-examples docker release check

build: check build-server build-ofaccheck build-ofactest build-examples
	cd webui/ && npm install && npm run build && cd ../

build-server:
	CGO_ENABLED=1 go build -o ./bin/server github.com/moov-io/ofac/cmd/server

build-ofaccheck:
	CGO_ENABLED=0 go build -o ./bin/ofaccheck github.com/moov-io/ofac/cmd/ofaccheck

build-ofactest:
	CGO_ENABLED=0 go build -o ./bin/ofactest github.com/moov-io/ofac/cmd/ofactest

build-examples: build-webhook-example

build-webhook-example:
	CGO_ENABLED=0 go build -o ./bin/webhook-example github.com/moov-io/ofac/examples/webhook

check:
	go fmt ./...
	@mkdir -p ./bin/

.PHONY: client
client:
# Versions from https://github.com/OpenAPITools/openapi-generator/releases
	@chmod +x ./openapi-generator
	@rm -rf ./client
	OPENAPI_GENERATOR_VERSION=4.2.0 ./openapi-generator generate -i openapi.yaml -g go -o ./client
	rm -f client/go.mod client/go.sum
	go fmt ./...
	go build github.com/moov-io/ofac/client
	go test ./client

.PHONY: clean
clean:
	@rm -rf client/
	@rm -rf bin/
	@rm -f openapi-generator-cli-*.jar

dist: clean client build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/ofac-windows-amd64.exe github.com/moov-io/ofac/cmd/server
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -o bin/ofac-$(PLATFORM)-amd64 github.com/moov-io/ofac/cmd/server
endif

docker:
# Main OFAC server Docker image
	docker build --pull -t moov/ofac:$(VERSION) -f Dockerfile .
	docker tag moov/ofac:$(VERSION) moov/ofac:latest
# ofactest image
	docker build --pull -t moov/ofactest:$(VERSION) -f ./cmd/ofactest/Dockerfile .
	docker tag moov/ofactest:$(VERSION) moov/ofactest:latest
# webhook example
	docker build --pull -t moov/ofac-webhook-example:$(VERSION) -f ./examples/webhook/Dockerfile .
	docker tag moov/ofac-webhook-example:$(VERSION) moov/ofac-webhook-example:latest

release: docker AUTHORS
	go vet ./...
	go test -coverprofile=cover-$(VERSION).out ./...
	git tag -f $(VERSION)

release-push:
	docker push moov/ofac:$(VERSION)
	docker push moov/ofac:latest
	docker push moov/ofactest:$(VERSION)
	docker push moov/ofac-webhook-example:$(VERSION)

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out

clean-integration:
	docker-compose kill
	docker-compose rm -v -f

test-integration: clean-integration
	docker-compose up -d
	sleep 10
	curl -v http://localhost:9094/data/refresh # hangs until download and parsing completes
	./bin/ofactest -local

# From https://github.com/genuinetools/img
.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@
