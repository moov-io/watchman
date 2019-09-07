FROM golang:1.13-buster as builder
WORKDIR /go/src/github.com/moov-io/ofac
RUN apt-get update && apt-get install make gcc g++
COPY . .
ENV GO111MODULE=on
RUN go mod download
RUN make build-server

FROM debian:10
RUN apt-get update && apt-get install -y ca-certificates

COPY --from=builder /go/src/github.com/moov-io/ofac/bin/server /bin/server
# USER moov # TODO(adam): non-root users

EXPOSE 8080
EXPOSE 9090
ENTRYPOINT ["/bin/server"]
