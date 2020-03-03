# crd-docs-generator

Generates schema reference documentation for Kubernetes Custom Resource Definitions (CRDs)

This tool (work in progress) is built to generate our Control Plane Kubernetes API schema reference in https://docs.giantswarm.io/.

The generated output consists of Markdown files loaded with HTML. By itself, this does not provide a fully readable and user-friendly set of documentation pages. Instead it relies on the HUGO website context, as the [giantswarm/docs](https://github.com/giantswarm/docs) repository.

## Usage

Execute the generator simply by running the following command in the repository folder:

```nohighlight
go run main.go
```

The Docker version can be executed like this:

```nohighlight
docker run \
    -v $PWD/path/to/output-folder:/opt/crd-docs-generator/output \
    quay.io/giantswarm/crd-docs-generator
```

## TODO

- Have a main description for each CRD's purpose
- Show CR example (per version)
- Date in front matter should ideally reflect last modification, not docs generation
- Parse template only once instead of for every CRD
