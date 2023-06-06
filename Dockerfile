FROM golang:1.19-bullseye


RUN set -eux; \
    apt-get update; \
    apt-get install -y --no-install-recommends \
        ca-certificates  curl \
        sudo bash \
    ;\
    curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $(go env GOPATH)/bin v1.52.2; \
    apt-get clean; \
    rm -rf /var/lib/apt/lists/*;

WORKDIR /src
CMD ["/bin/bash"]