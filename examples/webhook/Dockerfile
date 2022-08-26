FROM golang:1.19-alpine as builder
WORKDIR /go/src/github.com/moov-io/watchman
RUN apk add -U make git
RUN adduser -D -g '' --shell /bin/false moov
COPY . .
RUN go mod download
RUN make build-webhook-example
USER moov

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/watchman/bin/webhook-example /bin/webhook-example
COPY --from=builder /etc/passwd /etc/passwd
USER moov
EXPOSE 10101
ENTRYPOINT ["/bin/webhook-example"]
