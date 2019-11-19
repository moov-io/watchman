FROM golang:1.13-buster as backend
WORKDIR /go/src/github.com/moov-io/sanctionsearch
RUN apt-get update && apt-get install make gcc g++
COPY . .
ENV GO111MODULE=on
RUN go mod download
RUN make build-server

FROM node:12-buster as frontend
COPY webui/ /sanctionsearch/
WORKDIR /sanctionsearch/
RUN npm install
RUN npm run build

FROM debian:10
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=backend /go/src/github.com/moov-io/sanctionsearch/bin/server /bin/server

COPY --from=frontend /sanctionsearch/build/ /sanctionsearch/
ENV WEB_ROOT=/sanctionsearch/

# USER moov # TODO(adam): non-root users

EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/server"]
