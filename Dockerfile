FROM golang:1.16-buster as backend
WORKDIR /go/src/github.com/moov-io/watchman
RUN apt-get update && apt-get install make gcc g++
COPY . .
RUN go mod download
RUN make build-server

FROM node:12-buster as frontend
COPY webui/ /watchman/
WORKDIR /watchman/
RUN npm install
RUN npm run build

FROM debian:10
LABEL maintainer="Moov <support@moov.io>"

RUN apt-get update && apt-get install -y ca-certificates
COPY --from=backend /go/src/github.com/moov-io/watchman/bin/server /bin/server

COPY --from=frontend /watchman/build/ /watchman/
ENV WEB_ROOT=/watchman/

# USER moov # TODO(adam): non-root users

EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/server"]
