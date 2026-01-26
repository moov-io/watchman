# Backend build stage
FROM ghcr.io/knights-analytics/hugot AS backend
ARG VERSION
WORKDIR /src

# Install system dependencies first
RUN yum makecache && yum install -y autoconf automake libtool pkgconfig git

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

ENV PKG_CONFIG_PATH=/src/libpostal/
RUN VERSION=${VERSION} GOTAGS="-tags libpostal,embeddings,ORT" make build-server postal-server

# Pull Models
FROM moov/watchman-onnx-export AS models

# Final stage
FROM debian:trixie
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

# Copy model files
COPY --from=models /model/* ./models/multilingual-minilm/

# Set environment variables
ENV LD_LIBRARY_PATH=/usr/local/lib
ENV LIBPOSTAL_DATA_DIR=/usr/local/share/libpostal

EXPOSE 8084
EXPOSE 9094

ENTRYPOINT ["/usr/bin/tini", "--"]
CMD ["/bin/server"]
