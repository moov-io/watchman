PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+(-[a-zA-Z0-9]*)?)' version.go)

.PHONY: build build-server build-examples docker release check

build: check build-server build-batchsearch build-sanctiontest build-examples
	cd webui/ && npm install && npm run build && cd ../

build-server:
	CGO_ENABLED=1 go build -o ./bin/server github.com/moov-io/sanctionsearch/cmd/server

build-batchsearch:
	CGO_ENABLED=0 go build -o ./bin/batchsearch github.com/moov-io/sanctionsearch/cmd/batchsearch

build-sanctiontest:
	CGO_ENABLED=0 go build -o ./bin/sanctiontest github.com/moov-io/sanctionsearch/cmd/sanctiontest

build-examples: build-webhook-example

build-webhook-example:
	CGO_ENABLED=0 go build -o ./bin/webhook-example github.com/moov-io/sanctionsearch/examples/webhook

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
	go build github.com/moov-io/sanctionsearch/client
	go test ./client

.PHONY: clean
clean:
	@rm -rf client/
	@rm -rf bin/
	@rm -f openapi-generator-cli-*.jar

dist: clean client build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/sanctionsearch-windows-amd64.exe github.com/moov-io/sanctionsearch/cmd/server
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -o bin/sanctionsearch-$(PLATFORM)-amd64 github.com/moov-io/sanctionsearch/cmd/server
endif

docker:
# main server Docker image
	docker build --pull -t moov/sanctionsearch:$(VERSION) -f Dockerfile .
	docker tag moov/sanctionsearch:$(VERSION) moov/sanctionsearch:latest
# sanctiontest image
	docker build --pull -t moov/sanctiontest:$(VERSION) -f ./cmd/sanctiontest/Dockerfile .
	docker tag moov/sanctiontest:$(VERSION) moov/sanctiontest:latest
# webhook example
	docker build --pull -t moov/sanctionsearch-webhook-example:$(VERSION) -f ./examples/webhook/Dockerfile .
	docker tag moov/sanctionsearch-webhook-example:$(VERSION) moov/sanctionsearch-webhook-example:latest

release: docker AUTHORS
	go vet ./...
	go test -coverprofile=cover-$(VERSION).out ./...
	git tag -f $(VERSION)

release-push:
	docker push moov/sanctionsearch:$(VERSION)
	docker push moov/sanctionsearch:latest
	docker push moov/batchsearch:$(VERSION)
	docker push moov/sanctionsearch-webhook-example:$(VERSION)

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
	./bin/batchsearch -local -threshold 0.95

# From https://github.com/genuinetools/img
.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@
