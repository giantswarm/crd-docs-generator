# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

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

[Unreleased]: https://github.com/giantswarm/crd-docs-generator/compare/v0.4.0...HEAD
[0.4.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.3.0...v0.4.0
[0.3.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.2.3...v0.3.0
[0.2.3]: https://github.com/giantswarm/crd-docs-generator/compare/v0.2.2...v0.2.3
[0.2.2]: https://github.com/giantswarm/crd-docs-generator/compare/v0.2.1...v0.2.2
[0.2.1]: https://github.com/giantswarm/crd-docs-generator/compare/v0.2.0...v0.2.1
[0.2.0]: https://github.com/giantswarm/crd-docs-generator/compare/v0.1.2...v0.2.0
[0.1.2]: https://github.com/giantswarm/crd-docs-generator/compare/v0.1.1...v0.1.2
[0.1.1]: https://github.com/giantswarm/crd-docs-generator/compare/v0.1.0...v0.1.1
[0.1.0]: https://github.com/giantswarm/crd-docs-generator/releases/tag/v0.1.0
