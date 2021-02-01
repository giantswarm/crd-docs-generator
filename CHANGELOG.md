# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.3.0] - 2021-02-01

- Change name "Management Cluster API" to "Management API".

## v0.2.2

- Add more terminology changes and add aliases for redirects after URL changes.

## v0.2.1

- Change name "Control Plane Kubernetes API" to "Management Cluster API".

## v0.2.0

- Remove date field from front matter of generated pages, as it's no longer needed.

## v0.1.1

- Add a link target to every attribute name headline.

## v0.1.0

- Add blacklisting feature to skip certain CRDs that should not get documented
- Move example CR above property details
- Fix a headline tag
- Adapt CRD input path to match latest changes in the apiextensions repo
- Refactor: move functions into services
- Use config file for settings instead of flags
- Switch CI from architect to architect-orb

[Unreleased]: https://github.com/giantswarm/crd-docs-generator/compare/v0.3.0...HEAD
[0.3.0]: https://github.com/giantswarm/crd-docs-generator/releases/tag/v0.3.0
