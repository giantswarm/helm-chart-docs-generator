FROM quay.io/giantswarm/alpine:3.19.0

RUN apk add --no-cache ca-certificates git

COPY . /opt/cluster-app-docs-generator

WORKDIR /opt/cluster-app-docs-generator

ENTRYPOINT ["/opt/cluster-app-docs-generator/cluster-app-docs-generator"]
