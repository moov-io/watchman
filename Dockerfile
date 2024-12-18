FROM golang:1.23-bookworm as backend
ARG VERSION
WORKDIR /src
COPY . /src
RUN go mod download
RUN apt-get update && apt-get install -y curl autoconf automake libtool pkg-config
RUN make install
RUN go build -ldflags "-X github.com/moov-io/watchman.Version=${VERSION}" -o ./bin/server /src/cmd/server

FROM node:22-bookworm as frontend
ARG VERSION
COPY webui/ /watchman/
WORKDIR /watchman/
RUN npm install --legacy-peer-deps
RUN npm run build

FROM debian:bookworm
LABEL maintainer="Moov <oss@moov.io>"
COPY --from=backend /src/bin/server /bin/server
COPY --from=frontend /watchman/build/ /watchman/
ENV WEB_ROOT=/watchman/


EXPOSE 8084
EXPOSE 9094

ENTRYPOINT ["/bin/server"]
