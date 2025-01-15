PLATFORM=$(shell uname -s | tr '[:upper:]' '[:lower:]')
PWD := $(shell pwd)

ifndef VERSION
	ifndef RELEASE
	# If we're not publishing a release, set the dev commit hash
		ifndef DEV_TAG_SHA
			COMMIT_HASH :=$(shell git rev-parse --short=7 HEAD)
		else
			COMMIT_HASH :=$(shell echo ${DEV_TAG_SHA} | cut -c 1-7)
		endif
		VERSION := dev-${COMMIT_HASH}
	else
		VERSION := $(shell git describe --tags --abbrev=0)
	endif
endif

# If arch is define create a new env variable with the path slash for docker.
ifdef ARCH
	ARCH_SUFFIX := .${ARCH}
	ARCH_PATH := ${ARCH}/
endif

.PHONY: all run build docker release check test

all: build

run:
	go run github.com/moov-io/watchman/cmd/server

# Detect OS
ifeq ($(OS),Windows_NT)
    detected_OS := Windows
else
    detected_OS := $(shell uname -s)
endif

# Detect architecture for macOS
ifeq ($(detected_OS),Darwin)
    ARCH := $(shell uname -m)
    ifeq ($(ARCH),arm64)
        CONFIGURE_FLAGS := --disable-sse2
    endif
else
    CONFIGURE_FLAGS :=
endif

# Detect if we need sudo
SUDO := $(shell if command -v sudo >/dev/null 2>&1 && sudo -n true >/dev/null 2>&1; then echo "sudo"; else echo ""; fi)

# Installation target
install:
ifeq ($(detected_OS),Windows)
	@$(MAKE) install-windows
else ifeq ($(detected_OS),Linux)
	@$(MAKE) install-linux
else ifeq ($(detected_OS),Darwin)
	@$(MAKE) install-macos
else
	@echo "Unsupported operating system: $(detected_OS)"
	@exit 1
endif

install-linux:
	@$(MAKE) install-libpostal

install-macos:
	brew install curl autoconf automake libtool pkg-config
ifeq ($(ARCH),arm64)
	@echo "ARM architecture detected (M1/M2). SSE2 will be disabled."
else
	@echo "Intel architecture detected. SSE2 optimizations will be enabled."
endif
	@$(MAKE) install-libpostal

install-windows:
	pacman -Syu
	pacman -S autoconf automake curl git make libtool gcc mingw-w64-x86_64-gcc
	@$(MAKE) install-libpostal

install-libpostal:
	@echo "Cloning libpostal repository..."
	git clone https://github.com/openvenues/libpostal || true
	cd libpostal && \
	./bootstrap.sh && \
	./configure $(CONFIGURE_FLAGS) && \
	make -j$(shell nproc || echo 4) && \
	if [ "$(detected_OS)" = "Windows" ]; then \
		make install; \
	else \
		$(SUDO) make install; \
	fi

.PHONY: install install-linux install-macos install-windows install-libpostal

models: models-setup us-csl-models

models-setup:
	go install github.com/gocomply/xsd2go/cli/gocomply_xsd2go@latest

us-csl-models:
	@mkdir -p ./pkg/csl_us/gen/
	wget -O ./pkg/csl_us/gen/ENHANCED_XML.xsd https://sanctionslistservice.ofac.treas.gov/api/PublicationPreview/exports/ENHANCED_XML.xsd
	gocomply_xsd2go convert \
		./pkg/csl_us/gen/ENHANCED_XML.xsd \
		github.com/moov-io/watchman/pkg/csl_us/gen ./pkg/csl_us/gen/

.PHONY: models models-setup us-csl-models

.PHONY: build build-server
build: build-server

build-server:
	go build ${GOTAGS} -ldflags "-X github.com/moov-io/watchman.Version=${VERSION}" -o ./bin/server github.com/moov-io/watchman/cmd/server

.PHONY: check
check:
ifeq ($(OS),Windows_NT)
	go test ./... -short
else
	@wget -O lint-project.sh https://raw.githubusercontent.com/moov-io/infra/master/go/lint-project.sh
	@chmod +x ./lint-project.sh
	COVER_THRESHOLD=disabled DISABLE_GITLEAKS=true GOLANGCI_LINTERS=gocheckcompilerdirectives,mirror,tenv STRICT_GOLANGCI_LINTERS=no ./lint-project.sh
endif

.PHONY: clean
clean:
ifeq ($(OS),Windows_NT)
	@echo "Skipping cleanup on Windows, currently unsupported."
else
	@rm -rf ./bin/ cover.out coverage.txt lint-project.sh misspell* staticcheck* openapi-generator-cli-*.jar
endif

dist: clean build
ifeq ($(OS),Windows_NT)
	GOOS=windows go build -o bin/watchman.exe github.com/moov-io/watchman/cmd/server
else
	GOOS=${PLATFORM} go build -o bin/watchman-${PLATFORM}-amd64 github.com/moov-io/watchman/cmd/server
endif

docker: clean docker-hub docker-openshift docker-static

docker-hub:
	docker build --pull --build-arg VERSION=${VERSION} -t moov/watchman:${VERSION} -f Dockerfile .
	docker tag moov/watchman:${VERSION} moov/watchman:latest

docker-openshift:
	docker build --pull --build-arg VERSION=${VERSION} -t quay.io/moov/watchman:${VERSION} -f Dockerfile-openshift --build-arg VERSION=${VERSION} .
	docker tag quay.io/moov/watchman:${VERSION} quay.io/moov/watchman:latest

docker-static:
	docker build --pull -t moov/watchman:static -f Dockerfile-static .

release-push:
	docker push moov/watchman:${VERSION}
	docker push moov/watchman:latest
	docker push moov/watchman:static

quay-push:
	docker push quay.io/moov/watchman:${VERSION}
	docker push quay.io/moov/watchman:latest
