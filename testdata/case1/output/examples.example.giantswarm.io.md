description: This is the description of examples.example.giantswarm.io v1alpha2 from file crd2.yaml.
group: example.giantswarm.io
name_plural: examples
name_singular: example
scope: Namespaced
source_repository_ref: main
source_repository: https://github.com/giantswarm/crd-docs-generator
title: Example
versions: [v1alpha2 v1alpha3]
weight: 100

Metadata:

topics: [topic1 topic2]
providers: [provider1 provider2]





Version v1alpha2



Properties



Depth: 0
Path: .apiVersion
string



<p>APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources</a></p>


Depth: 0
Path: .kind
string



<p>Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</a></p>


Depth: 0
Path: .metadata
object




Depth: 0
Path: .spec
object
Required


<p>Specifies the desired state of an example resource.</p>


Depth: 1
Path: .spec.myObject
object



<p>An example object that is nullable, without any required properties.</p>


Depth: 2
Path: .spec.myObject.mySubObject
object



<p>A nullable object as a property of &lsquo;myObject&rsquo; with some required props.</p>


Depth: 3
Path: .spec.myObject.mySubObject.name
string
Required


<p>Name is a name of something.</p>


Depth: 3
Path: .spec.myObject.mySubObject.namespace
string



<p>Namespace is the namespace of something.</p>


Depth: 2
Path: .spec.myObject.otherSubObject
object



<p>A non-nullable sub object.</p>


Depth: 3
Path: .spec.myObject.otherSubObject.first
string



<p>First string property.</p>


Depth: 3
Path: .spec.myObject.otherSubObject.second
string



<p>Second string property.</p>




Version v1alpha3



Properties



Depth: 0
Path: .apiVersion
string



<p>APIVersion defines the versioned schema of this representation of an object. Servers should convert recognized schemas to the latest internal value, and may reject unrecognized values. More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#resources</a></p>


Depth: 0
Path: .kind
string



<p>Kind is a string value representing the REST resource this object represents. Servers may infer this from the endpoint the client submits requests to. Cannot be updated. In CamelCase. More info: <a href="https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds">https://git.k8s.io/community/contributors/devel/sig-architecture/api-conventions.md#types-kinds</a></p>


Depth: 0
Path: .metadata
object




Depth: 0
Path: .spec
object
Required


<p>Specifies the desired state of an example resource.</p>


Depth: 1
Path: .spec.myObject
object



<p>An example object that is nullable, without any required properties.</p>


Depth: 2
Path: .spec.myObject.mySubObject
object



<p>A nullable object as a property of &lsquo;myObject&rsquo; with some required props.</p>


Depth: 3
Path: .spec.myObject.mySubObject.name
string
Required


<p>Name is a name of something.</p>


Depth: 3
Path: .spec.myObject.mySubObject.namespace
string



<p>Namespace is the namespace of something.</p>



