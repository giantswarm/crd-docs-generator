module github.com/giantswarm/crd-docs-generator

go 1.15

require (
	github.com/Masterminds/goutils v1.1.0 // indirect
	github.com/Masterminds/semver v1.5.0 // indirect
	github.com/Masterminds/sprig v2.22.0+incompatible
	github.com/ghodss/yaml v1.0.0
	github.com/giantswarm/microerror v0.3.0
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/mitchellh/copystructure v1.0.0 // indirect
	github.com/russross/blackfriday/v2 v2.1.0
	github.com/spf13/cobra v1.1.3
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apiextensions-apiserver v0.20.2
)

replace github.com/coreos/etcd => github.com/coreos/etcd v3.3.25+incompatible
