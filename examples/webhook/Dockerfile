FROM golang:1.11-alpine as builder
WORKDIR /go/src/github.com/moov-io/ofac
RUN apk add -U make
RUN adduser -D -g '' --shell /bin/false moov
COPY . .
RUN make build-examples
USER moov

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/ofac/bin/example /bin/example
COPY --from=builder /etc/passwd /etc/passwd
USER moov
EXPOSE 10101
ENTRYPOINT ["/bin/example"]
