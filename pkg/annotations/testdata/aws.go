// This is an example Go file
// containing special doc comments we need to
// associate annotations with CRDs.
package aws

// Comment above is possible
// CRD_DOCS_GENERATOR:
//
//	support:
//	  - crd: awsclusters.infrastructure.giantswarm.io
//	    apiversion: v1alpha2
//	    release: Since 14.0.0
//	documentation:
//	  This annotation allows configuration of the MINIMUM_IP_TARGET parameter for AWS CNI.
const AWSCNIMinimumIPTarget = "alpha.cni.aws.giantswarm.io/minimum-ip-target"
