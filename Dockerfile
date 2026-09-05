FROM gsoci.azurecr.io/giantswarm/alpine:3.24.1

RUN apk add --no-cache ca-certificates git && apk add --no-cache libc6-compat

# buildx sets TARGETARCH per platform (amd64, arm64). Both schemadocs and
# architect/go-build publish their linux builds under that arch name.
ARG TARGETARCH

# Schemadocs version
ARG SCHEMADOCS_VERSION=0.4.0

# Download and install schemadocs for the image's own architecture
RUN wget -q https://github.com/giantswarm/schemadocs/releases/download/v${SCHEMADOCS_VERSION}/schemadocs-v${SCHEMADOCS_VERSION}-linux-${TARGETARCH}.tar.gz && \
    tar -C /tmp -xzf schemadocs-v${SCHEMADOCS_VERSION}-linux-${TARGETARCH}.tar.gz && \
    rm schemadocs-v${SCHEMADOCS_VERSION}-linux-${TARGETARCH}.tar.gz && \
    mv /tmp/schemadocs-v${SCHEMADOCS_VERSION}-linux-${TARGETARCH}/schemadocs /usr/local/bin/schemadocs && \
    rm -rf /tmp/schemadocs-v${SCHEMADOCS_VERSION}-linux-${TARGETARCH}

# architect/go-build emits one static binary per target platform
# (helm-chart-docs-generator-linux-amd64, -linux-arm64) plus an unsuffixed copy
# of the linux/amd64 build. Copy only the one matching TARGETARCH; the rest of
# the checkout is not needed at runtime. For a local build, produce it first:
#   CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o helm-chart-docs-generator-linux-amd64 .
COPY ./helm-chart-docs-generator-linux-${TARGETARCH} /opt/helm-chart-docs-generator/helm-chart-docs-generator

WORKDIR /opt/helm-chart-docs-generator

ENTRYPOINT ["/opt/helm-chart-docs-generator/helm-chart-docs-generator"]
