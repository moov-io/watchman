FROM golang:1.13-buster as backend
WORKDIR /go/src/github.com/moov-io/ofac
RUN apt-get update && apt-get install make gcc g++
COPY . .
ENV GO111MODULE=on
RUN go mod download
RUN make build-server

FROM node:12-buster as frontend
COPY webui/ /ofac/
WORKDIR /ofac/
RUN npm install
RUN npm run build

FROM debian:10
RUN apt-get update && apt-get install -y ca-certificates
COPY --from=backend /go/src/github.com/moov-io/ofac/bin/server /bin/server

COPY --from=frontend /ofac/build/ /ofac/
ENV WEB_ROOT=/ofac/

# USER moov # TODO(adam): non-root users

EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/server"]
