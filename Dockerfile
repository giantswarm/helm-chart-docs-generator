FROM gsoci.azurecr.io/giantswarm/alpine:3.19.1

RUN apk add --no-cache ca-certificates git

COPY . /opt/helm-chart-docs-generator

WORKDIR /opt/helm-chart-docs-generator

ENTRYPOINT ["/opt/helm-chart-docs-generator/helm-chart-docs-generator"]
