[![CircleCI](https://circleci.com/gh/giantswarm/crd-docs-generator/tree/master.svg?style=shield&circle-token=c0f46d2b8c1482706d8d41b098d488efdf637a1f)](https://circleci.com/gh/giantswarm/crd-docs-generator/tree/master)
[![Docker Repository on Quay](https://quay.io/repository/giantswarm/crd-docs-generator/status "Docker Repository on Quay")](https://quay.io/repository/giantswarm/crd-docs-generator)

# crd-docs-generator

Generates schema reference documentation for Kubernetes Custom Resource Definitions (CRDs).

This tool is built to generate our Control Plane Kubernetes API schema reference in https://docs.giantswarm.io/.

The generated output consists of Markdown files packed with HTML. By itself, this does not provide a fully readable and user-friendly set of documentation pages. Instead it relies on the HUGO website context, as the [giantswarm/docs](https://github.com/giantswarm/docs) repository, to provide an index page and useful styling.

## Usage

The generator can be executed in Docker using a command like this:

```nohighlight
docker run \
    -v $PWD/path/to/output-folder:/opt/crd-docs-generator/output \
    quay.io/giantswarm/crd-docs-generator
```

The volume mapping defines where the generated output will land.

## TODO

- Read CRD YAML directly from [apiextensions](https://github.com/giantswarm/apiextensions/tree/master/docs/crd), where available
- Read example CR directly from [apiextensions](https://github.com/giantswarm/apiextensions/tree/master/docs/cr), where available
- Have a main description for each CRD's purpose (to be fixed in apiextensions source)
- Show CR example (per version)
- Date in front matter should ideally reflect last modification, not docs generation
- Parse template only once instead of for every CRD
