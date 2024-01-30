[![CircleCI](https://circleci.com/gh/giantswarm/crd-docs-generator/tree/master.svg?style=svg&circle-token=2847f4b99edcb9776cbd8ee622b294eb96bfd55f)](https://circleci.com/gh/giantswarm/crd-docs-generator/tree/master)

# crd-docs-generator

Generates schema reference documentation for Kubernetes Custom Resource Definitions (CRDs).

This tool is built to generate our [Management API CRD schema reference](https://docs.giantswarm.io/ui-api/management-api/crd/).

The generated output consists of Markdown files packed with HTML. By itself, this does not provide a fully readable and user-friendly set of documentation pages. Instead it relies on the HUGO website context, as the [giantswarm/docs](https://github.com/giantswarm/docs) repository, to provide an index page and useful styling.

## Assumptions/Prerequisites

This tool relies on:

- CRDs being defined in public source repositories in YAML format.
- CRDs providing an OpenAPIv3 validation schema
  - either in the `.spec.validation` section of a CRD containg only one version
  - or in the `.spec.versions[*].schema` position of a CRD containing multiple versions
- OpenAPIv3 schemas containing `description` attributes for every property.
- The topmost `description` value explaining the CRD itself. (For a CRD containing multiple versions, the first `description` found is used as such.)
- CR examples to be found in the source repository/repositories as one example per YAML file.

## Usage

### Docker

The generator can be executed in Docker using a command like this:

```nohighlight
docker run \
    -v $PWD/path/to/output-folder:/opt/crd-docs-generator/output \
    -v $PWD:/opt/crd-docs-generator/config \
    gsoci.azurecr.io/giantswarm/crd-docs-generator:0.11.0 \
      --config /opt/crd-docs-generator/config/config.example.yaml
```

Here, the tag `0.11.0` is the version number of the crd-docs-generator release you're going to use. See our GitHub releases for available tags.

The volume mapping defines where the generated output will land.

### Development

With Go installed and this repository cloned, you can exetute the program like this:

```nohighlight
go run main.go --config config.example.yaml
```

See the `config.example.yaml` file for an idea how to configure your source repositories.

## TODO

- Parse template only once instead of for every CRD
