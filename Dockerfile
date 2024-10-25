FROM golang:alpine as backend
ARG VERSION
WORKDIR /src
COPY . /src
RUN go mod download
RUN CGO_ENABLED=0 go build -ldflags "-X github.com/moov-io/watchman.Version=${VERSION}" -o ./bin/server /src/cmd/server

FROM node:21-alpine as frontend
ARG VERSION
COPY webui/ /watchman/
WORKDIR /watchman/
RUN npm install --legacy-peer-deps
RUN npm run build

FROM alpine:latest
LABEL maintainer="Moov <oss@moov.io>"
COPY --from=backend /src/bin/server /bin/server
COPY --from=frontend /watchman/build/ /watchman/
ENV WEB_ROOT=/watchman/


EXPOSE 8084
EXPOSE 9094

ENTRYPOINT ["/bin/server"]
