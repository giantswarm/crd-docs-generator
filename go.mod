module github.com/giantswarm/crd-docs-generator

go 1.16

require (
	github.com/Masterminds/sprig/v3 v3.2.2
	github.com/ghodss/yaml v1.0.0
	github.com/giantswarm/microerror v0.3.0
	github.com/google/go-cmp v0.5.6
	github.com/huandu/xstrings v1.3.2 // indirect
	github.com/russross/blackfriday/v2 v2.1.0
	github.com/spf13/cobra v1.2.1
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apiextensions-apiserver v0.20.10
	k8s.io/apimachinery v0.22.1
)

replace (
	github.com/coreos/etcd => github.com/coreos/etcd v3.3.25+incompatible
	github.com/dgrijalva/jwt-go v3.2.0+incompatible => github.com/golang-jwt/jwt/v4 v4.0.0
	github.com/gogo/protobuf => github.com/gogo/protobuf v1.3.2 // CVE-2021-3121
)
