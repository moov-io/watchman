FROM golang:1.23-bookworm as backend
ARG VERSION
WORKDIR /src
COPY . /src
RUN go mod download
RUN apt-get update && apt-get install -y curl autoconf automake libtool pkg-config git

# Clone and install libpostal
RUN git clone https://github.com/openvenues/libpostal.git /src/libpostal
WORKDIR /src/libpostal
RUN ./bootstrap.sh && \
    ./configure && \
    make -j$(shell nproc) && \
    make install && \
    ldconfig

# Download libpostal data files
RUN libpostal_data download all /usr/local/share/libpostal

# Build the application
WORKDIR /src
RUN go build -ldflags "-X github.com/moov-io/watchman.Version=${VERSION}" -o ./bin/server /src/cmd/server

FROM node:22-bookworm as frontend
ARG VERSION
COPY webui/ /watchman/
WORKDIR /watchman/
RUN npm install --legacy-peer-deps
RUN npm run build

FROM debian:bookworm
LABEL maintainer="Moov <oss@moov.io>"

# Install required runtime dependencies
RUN apt-get update && \
    apt-get install -y \
    libssl3 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Create necessary directories
RUN mkdir -p /usr/local/share/libpostal

# Copy libpostal shared libraries and configuration
COPY --from=backend /usr/local/lib/libpostal.so* /usr/local/lib/
COPY --from=backend /usr/local/lib/pkgconfig/libpostal.pc /usr/local/lib/pkgconfig/
COPY --from=backend /usr/local/share/libpostal/ /usr/local/share/libpostal/

# Update shared library cache
RUN ldconfig

# Copy application files
COPY --from=backend /src/bin/server /bin/server
COPY --from=frontend /watchman/build/ /watchman/

ENV WEB_ROOT=/watchman/
ENV LD_LIBRARY_PATH=/usr/local/lib
ENV LIBPOSTAL_DATA_DIR=/usr/local/share/libpostal

EXPOSE 8084
EXPOSE 9094
ENTRYPOINT ["/bin/server"]
