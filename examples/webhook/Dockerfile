FROM golang:1.12-alpine as builder
WORKDIR /go/src/github.com/moov-io/ofac
RUN apk add -U make git
RUN adduser -D -g '' --shell /bin/false moov
COPY . .
ENV GO111MODULE=on
RUN go mod download
RUN make build-webhook-example
USER moov

FROM scratch
COPY --from=builder /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/ca-certificates.crt
COPY --from=builder /go/src/github.com/moov-io/ofac/bin/webhook-example /bin/webhook-example
COPY --from=builder /etc/passwd /etc/passwd
USER moov
EXPOSE 10101
ENTRYPOINT ["/bin/webhook-example"]
