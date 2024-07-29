FROM gsoci.azurecr.io/giantswarm/alpine:3.20.2

RUN apk add --no-cache ca-certificates git && apk add --no-cache libc6-compat

# Schemadocs version
ARG SCHEMADOCS_VERSION=0.0.5

# Download and install schemadocs
RUN wget https://github.com/giantswarm/schemadocs/releases/download/v${SCHEMADOCS_VERSION}/schemadocs-v${SCHEMADOCS_VERSION}-linux-amd64.tar.gz && \
    tar -C /tmp -xzf schemadocs-v${SCHEMADOCS_VERSION}-linux-amd64.tar.gz && \
    rm schemadocs-v${SCHEMADOCS_VERSION}-linux-amd64.tar.gz

# Move schemadocs binary to /usr/local/bin
RUN mv /tmp/schemadocs-v${SCHEMADOCS_VERSION}-linux-amd64/schemadocs /usr/local/bin/schemadocs

COPY . /opt/helm-chart-docs-generator

WORKDIR /opt/helm-chart-docs-generator

ENTRYPOINT ["/opt/helm-chart-docs-generator/helm-chart-docs-generator"]
