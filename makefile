PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
VERSION := $(shell grep -Eo '(v[0-9]+[\.][0-9]+[\.][0-9]+(-[a-zA-Z0-9]*)?)' version.go)

.PHONY: run build build-server docker release check test

run:
	CGO_ENABLED=1 go run github.com/moov-io/watchman/cmd/server

build: build-server build-batchsearch build-watchmantest
ifeq ($(OS),Windows_NT)
	@echo "Skipping webui build on Windows."
else
	cd webui/ && npm install --legacy-peer-deps && npm run build && cd ../
endif

build-server:
	CGO_ENABLED=1 go build -o ./bin/server github.com/moov-io/watchman/cmd/server

build-batchsearch:
	CGO_ENABLED=0 go build -o ./bin/batchsearch github.com/moov-io/watchman/cmd/batchsearch

build-watchmantest:
	CGO_ENABLED=0 go build -o ./bin/watchmantest github.com/moov-io/watchman/cmd/watchmantest

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	@echo "Skipping checks on Windows, currently unsupported."
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	STRICT_GOLANGCI_LINTERS=no GOLANGCI_LINTERS=gocheckcompilerdirectives,mirror,tenv ./lint-project.sh
endif

.PHONY: admin
admin:
	@rm -rf ./admin
	docker run --rm \
		-u $(USERID):$(GROUPID) \
		-v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 batch -- /local/.openapi-generator/admin-generator-config.yml
	rm -f ./admin/go.mod ./admin/go.sum ./admin/.travis.yml
	gofmt -w ./admin/
	go build github.com/moov-io/watchman/admin

.PHONY: client
client:
	@rm -rf ./client
	docker run --rm \
		-u $(USERID):$(GROUPID) \
		-v ${PWD}:/local openapitools/openapi-generator-cli:v4.3.1 batch -- /local/.openapi-generator/client-generator-config.yml
	rm -f ./client/go.mod ./client/go.sum ./client/.travis.yml
	gofmt -w ./client/
	go build github.com/moov-io/watchman/client


.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
	@echo "Skipping cleanup on Windows, currently unsupported."
else
	@rm -rf ./bin/ cover.out coverage.txt lint-project.sh misspell* staticcheck* openapi-generator-cli-*.jar
endif

dist: clean build
ifeq ($(OS),Windows_NT)
	CGO_ENABLED=1 GOOS=windows go build -o bin/watchman.exe github.com/moov-io/watchman/cmd/server
else
	CGO_ENABLED=1 GOOS=$(PLATFORM) go build -o bin/watchman-$(PLATFORM)-amd64 github.com/moov-io/watchman/cmd/server
endif

docker: clean docker-hub docker-openshift docker-static docker-watchmantest

docker-hub:
	docker build --pull -t moov/watchman:$(VERSION) -f Dockerfile .
	docker tag moov/watchman:$(VERSION) moov/watchman:latest

docker-openshift:
	docker build --pull -t quay.io/moov/watchman:$(VERSION) -f Dockerfile-openshift --build-arg VERSION=$(VERSION) .
	docker tag quay.io/moov/watchman:$(VERSION) quay.io/moov/watchman:latest

docker-static:
	docker build --pull -t moov/watchman:static -f Dockerfile-static .

docker-watchmantest:
	docker build --pull -t moov/watchmantest:$(VERSION) -f ./cmd/watchmantest/Dockerfile .
	docker tag moov/watchmantest:$(VERSION) moov/watchmantest:latest

release: docker AUTHORS
	go vet ./...
	go test -coverprofile=cover-$(VERSION).out ./...
	git tag -f $(VERSION)

release-push:
	docker push moov/watchman:$(VERSION)
	docker push moov/watchman:latest
	docker push moov/watchman:static
	docker push moov/watchmantest:$(VERSION)

quay-push:
	docker push quay.io/moov/watchman:$(VERSION)
	docker push quay.io/moov/watchman:latest

.PHONY: cover-test cover-web
cover-test:
	go test -coverprofile=cover.out ./...
cover-web:
	go tool cover -html=cover.out

clean-integration:
	docker-compose kill && docker-compose rm -v -f

# TODO: this test is working but due to a default timeout on the admin server we get an empty reply
# for now this shouldn't hold up out CI pipeline
test-integration: clean-integration
	docker compose up -d
	sleep 30
	time curl -v --max-time 30 http://localhost:9094/data/refresh # hangs until download and parsing completes
	./bin/batchsearch -local -threshold 0.95

# From https://github.com/genuinetools/img
.PHONY: AUTHORS
AUTHORS:
	@$(file >$@,# This file lists all individuals having contributed content to the repository.)
	@$(file >>$@,# For how it is generated, see `make AUTHORS`.)
	@echo "$(shell git log --format='\n%aN <%aE>' | LC_ALL=C.UTF-8 sort -uf)" >> $@
