FROM golang:1.21-alpine as builder
RUN apk add -U make git
RUN adduser -D -g '' --shell /bin/false moov
WORKDIR /go/src/github.com/moov-io/watchman
COPY . .
RUN go mod download
RUN make build-watchmantest
USER moov

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/watchman/bin/watchmantest /bin/watchmantest
COPY --from=builder /etc/passwd /etc/passwd
USER moov
ENTRYPOINT ["/bin/watchmantest"]
