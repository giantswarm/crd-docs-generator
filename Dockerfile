FROM alpine:3.11

RUN apk add --no-cache ca-certificates

RUN mkdir -p /opt/crd-docs-generator
ADD ./crd-docs-generator /opt/crd-docs-generator/crd-docs-generator

WORKDIR /opt/crd-docs-generator

ENTRYPOINT ["/opt/crd-docs-generator/crd-docs-generator"]
