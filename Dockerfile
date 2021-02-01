FROM alpine:3.13.1

RUN apk add --no-cache ca-certificates git

ADD . /opt/crd-docs-generator

WORKDIR /opt/crd-docs-generator

ENTRYPOINT ["/opt/crd-docs-generator/crd-docs-generator"]
