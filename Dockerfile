FROM quay.io/giantswarm/alpine:3.14.2

RUN apk add --no-cache ca-certificates git

COPY . /opt/crd-docs-generator

WORKDIR /opt/crd-docs-generator

ENTRYPOINT ["/opt/crd-docs-generator/crd-docs-generator"]
