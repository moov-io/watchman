# Backend build stage
FROM golang:1.23-bookworm AS backend
ARG VERSION
WORKDIR /src

# Install system dependencies first
RUN apt-get update && apt-get install -y \
    curl \
    autoconf \
    automake \
    libtool \
    pkg-config \
    git

# Clone and build libpostal (rarely changes)
RUN git clone https://github.com/openvenues/libpostal.git /src/libpostal
WORKDIR /src/libpostal
RUN ./bootstrap.sh && \
    ./configure && \
    make -j$(shell nproc) && \
    make install && \
    ldconfig

# Download libpostal data (rarely changes)
RUN libpostal_data download all /usr/local/share/libpostal

# Copy go.mod and go.sum first to cache dependencies
COPY go.mod go.sum /src/
RUN go mod download

# Now copy the rest of the source code (frequently changes)
COPY . /src/
WORKDIR /src
RUN VERSION=${VERSION} GOTAGS="-tags libpostal" make build-server

# Final stage
FROM debian:bookworm
LABEL maintainer="Moov <oss@moov.io>"

# Install runtime dependencies
RUN apt-get update && \
    apt-get install -y \
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

# Set environment variables
ENV LD_LIBRARY_PATH=/usr/local/lib
ENV LIBPOSTAL_DATA_DIR=/usr/local/share/libpostal

EXPOSE 8084
EXPOSE 9094
ENTRYPOINT ["/bin/server"]
