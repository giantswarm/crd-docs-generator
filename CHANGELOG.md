# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.11.2] - 2024-12-17

## [0.11.1] - 2024-02-21

### Changed

- Use base images from `gsoci.azurecr.io`.
- Numerous dependency updates.

## [0.11.0] - 2022-10-28

### Changed

- **Potentially breaking:** Annotation doc comments have a new format, because go-fmt v1.19 breaks our old formatting. For an example look at the file /pkg/annotations/testdata/aws.go.

## [0.10.0] - 2022-03-22

- **Potentially breaking:** Your templates might have to be adapted like done in [this commit](https://github.com/giantswarm/crd-docs-generator/pull/98/files?file-filters%5B%5D=.template&show-viewed-files=true).
- The order of versions in a CRD output page is now guaranteed.
- The output path is now configurable via the config file directive `output_path`.
- If the output folder does not exist, it will be created.

## [0.9.0] - 2022-01-26

### Changed

Breaking: source paths are now configured in the config file.

Three new configuration keys have been introduced to configure paths per source_repository:

- `crd_paths`: paths to search for CRD YAML files.
- `cr_paths`: paths to search for example CR YAML files.
- `annotations_paths`: paths to search for Go files defining annotations.

All expect array values. Paths are relative to the source repo root. See the `config.example.yaml` file to learn how to use the keys.

## [0.8.0] - 2021-12-09

- Breaking: rename `.APIVersion` template field to `.CRDVersion`.
- Support multiple source repositories
- Update jwt-go dependency
- Refactoring to enable better testing
- Add test for output package
- Improve logging

## [0.7.1] - 2021-07-21

- Pass through deprecation metadata from apiextensions to output (https://github.com/giantswarm/crd-docs-generator/pull/46)

## [0.7.0] - 2021-07-19

- Use metadata from the apiextensions repository (https://github.com/giantswarm/crd-docs-generator/pull/44)

## [0.6.1] - 2021-05-14

- Add support for another apiextensions repo path `/helm/**/upstream.yaml`.
- Parse multiple CRDs from a single YAML file.

## [0.6.0] - 2021-05-14

- Change path where to look for CRD YAML in giantswarm/apiextensions from `/config/crd/v1` to `/config/crd`.

## [0.5.0] - 2021-02-18

- Breaking: Remove configuration option `commit_reference`, add command line flag `--commit-reference` for the same purpose instead.

## [0.4.0] - 2021-02-12

- Add flag `--template` to specify the template path
- Render annotations with documentation in CRDs.

## [0.3.0] - 2021-02-02

- Don't try to add syntax highlighting. Use tripple backtick instead.
- Update dependencies.

## [0.2.3] - 2021-01-14

- Change name "Management Cluster API" to "Management API".

## [0.2.2] - 2021-01-08

- Add more terminology changes and add aliases for redirects after URL changes.

## [0.2.1] - 2021-01-08

- Change name "Control Plane Kubernetes API" to "Management Cluster API".

## [0.2.0] - 2020-12-03

- Remove date field from front matter of generated pages, as it's no longer needed.

## [0.1.2] - 2020-10-05

- Remove whitespace around 'Required'.

## [0.1.1] - 2020-06-29

- Add a link target to every attribute name headline.

## [0.1.0] - 2020-05-06

- Add blacklisting feature to skip certain CRDs that should not get documented
- Move example CR above property details
- Fix a headline tag
- Adapt CRD input path to match latest changes in the apiextensions repo
- Refactor: move functions into services
- Use config file for settings instead of flags
- Switch CI from architect to architect-orb

[Unreleased]: https://github.com/giantswarm/crd-docs-generator/compare/v0.11.2...HEAD
[0.11.2]: https://github.com/giantswarm/crd-docs-generator/compare/v0.11.1...v0.11.2
[0.11.1]: https://github.com/giantswarm/crd-docs-generator/compare/v0.11.0...v0.11.1
[0.11.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.10.0...v0.11.0
[0.10.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.9.0...v0.10.0
[0.9.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.8.0...v0.9.0
[0.8.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.7.1...v0.8.0
[0.7.1]: https://github.com/giantswarm/crd-docs-generator/compare/v0.7.0...v0.7.1
[0.7.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.6.1...v0.7.0
[0.6.1]: https://github.com/giantswarm/crd-docs-generator/compare/v0.6.0...v0.6.1
[0.6.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.5.0...v0.6.0
[0.5.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.4.0...v0.5.0
[0.4.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.2.3...v0.3.0
[0.2.3]: https://github.com/giantswarm/crd-docs-generator/compare/v0.2.2...v0.2.3
[0.2.2]: https://github.com/giantswarm/crd-docs-generator/compare/v0.2.1...v0.2.2
[0.2.1]: https://github.com/giantswarm/crd-docs-generator/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.1.2...v0.2.0
[0.1.2]: https://github.com/giantswarm/crd-docs-generator/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/giantswarm/crd-docs-generator/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/giantswarm/crd-docs-generator/releases/tag/v0.1.0
