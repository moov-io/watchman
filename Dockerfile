FROM golang:alpine as backend
WORKDIR /src
COPY . /src
RUN go mod download
RUN CGO_ENABLED=0 go build -o ./bin/server /src/cmd/server

FROM node:21-alpine as frontend
COPY webui/ /watchman/
WORKDIR /watchman/
RUN npm install --legacy-peer-deps
RUN npm run build

FROM alpine:latest
LABEL maintainer="Moov <oss@moov.io>"
COPY --from=backend /src/bin/server /bin/server
COPY --from=frontend /watchman/build/ /watchman/
ENV WEB_ROOT=/watchman/

# USER moov # TODO(adam): non-root users

EXPOSE 8084
EXPOSE 9094
ENTRYPOINT ["/bin/server"]
