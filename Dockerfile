FROM golang:1.21-bookworm as backend
WORKDIR /go/src/github.com/moov-io/watchman
RUN apt-get update && apt-get upgrade -y && apt-get install make gcc g++
COPY . .
RUN go mod download
RUN make build-server

FROM node:21-bookworm as frontend
COPY webui/ /watchman/
WORKDIR /watchman/
RUN npm install --legacy-peer-deps
RUN npm run build

FROM debian:stable-slim
LABEL maintainer="Moov <oss@moov.io>"

RUN apt-get update && apt-get upgrade -y && apt-get install -y ca-certificates
COPY --from=backend /go/src/github.com/moov-io/watchman/bin/server /bin/server

COPY --from=frontend /watchman/build/ /watchman/
ENV WEB_ROOT=/watchman/

# USER moov # TODO(adam): non-root users

EXPOSE 8084
EXPOSE 9094
ENTRYPOINT ["/bin/server"]
