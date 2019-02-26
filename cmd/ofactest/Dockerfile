FROM golang:1.12-alpine as builder
RUN apk add -U make git
RUN adduser -D -g '' --shell /bin/false moov

# Build ofactest binary (using client generated with)
WORKDIR /go/src/github.com/moov-io/ofac
COPY . .
RUN make build-ofactest
USER moov

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/ofac/bin/ofactest /bin/ofactest
COPY --from=builder /etc/passwd /etc/passwd
USER moov
ENTRYPOINT ["/bin/ofactest"]
