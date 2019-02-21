FROM golang:1.11-stretch as builder
WORKDIR /go/src/github.com/moov-io/ofac
RUN apt-get update && apt-get install make gcc g++
COPY . .
RUN make build-server

FROM debian:9
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /go/src/github.com/moov-io/ofac/bin/server /bin/server
# USER moov # TODO(adam): non-root users

EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/server"]
