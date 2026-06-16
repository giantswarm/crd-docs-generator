FROM gsoci.azurecr.io/giantswarm/alpine:3.24.1

ARG TARGETARCH

RUN apk add --no-cache ca-certificates git

COPY ./crd-docs-generator-linux-${TARGETARCH} /opt/crd-docs-generator/crd-docs-generator

WORKDIR /opt/crd-docs-generator

ENTRYPOINT ["/opt/crd-docs-generator/crd-docs-generator"]
