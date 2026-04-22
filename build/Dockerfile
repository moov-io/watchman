# Backend build stage
FROM golang:1.26-bookworm AS backend
ARG VERSION
WORKDIR /src

# Install system dependencies first
RUN apt-get update && apt-get install -y autoconf automake libtool pkg-config git

# Clone and build libpostal (rarely changes)
RUN git clone https://github.com/openvenues/libpostal.git /src/libpostal
WORKDIR /src/libpostal
RUN ./bootstrap.sh && \
    ./configure && \
    make -j4 && \
    make install && \
    ldconfig

# Download libpostal data (rarely changes)
# This data comes from Senzing's improvements over the openvenues offered data
# See: https://github.com/Senzing/libpostal-data#version-120
RUN mkdir -p /usr/local/share/libpostal && \
    wget -q https://public-read-libpostal-data.s3.amazonaws.com/v1.2.0/parser.tar.gz && \
    wget -q https://public-read-libpostal-data.s3.amazonaws.com/v1.1.0/libpostal_data.tar.gz && \
    wget -q https://public-read-libpostal-data.s3.amazonaws.com/v1.1.0/language_classifier.tar.gz && \
    tar -xzf parser.tar.gz -C /usr/local/share/libpostal && \
    tar -xzf libpostal_data.tar.gz -C /usr/local/share/libpostal && \
    tar -xzf language_classifier.tar.gz -C /usr/local/share/libpostal && \
    rm *.tar.gz

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum /src/
RUN go mod download

# Now copy the rest of the source code (frequently changes)
COPY . /src/
WORKDIR /src

ENV PKG_CONFIG_PATH=/src/libpostal/
RUN VERSION=${VERSION} GOTAGS="-tags libpostal" make build-server postal-server

# Final stage
FROM debian:trixie-20260421
LABEL maintainer="Moov <oss@moov.io>"

# Install runtime dependencies
RUN apt-get update && \
    apt-get install -y \
    tini \
    libssl3 \
    ca-certificates \
    && rm -rf /var/lib/apt/lists/*

# Create necessary directories and copy libpostal files
RUN mkdir -p /usr/local/share/libpostal
COPY --from=backend /usr/local/lib/libpostal.so* /usr/local/lib/
COPY --from=backend /usr/local/lib/pkgconfig/libpostal.pc /usr/local/lib/pkgconfig/
COPY --from=backend /usr/local/share/libpostal/ /usr/local/share/libpostal/
RUN ldconfig

# Copy application files
COPY --from=backend /src/bin/server /bin/server
COPY --from=backend /src/bin/postal-server /bin/postal-server
ENV POSTAL_SERVER_BIN_PATH=/bin/postal-server

# Set environment variables
ENV LD_LIBRARY_PATH=/usr/local/lib
ENV LIBPOSTAL_DATA_DIR=/usr/local/share/libpostal

EXPOSE 8084
EXPOSE 9094

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["/bin/server"]
