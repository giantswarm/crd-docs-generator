---
title: AWSControlPlane CRD Schema Reference
linktitle: AWSControlPlane
technical_name: awscontrolplanes.infrastructure.giantswarm.io
description: >
  Custom Resource/Custom Resource Definition schema reference page
  for the AWSControlPlane resource (awscontrolplanes.infrastructure.giantswarm.io),
  as part of the Giant Swarm Control Plane Kubernetes API documentation.
date: 2020-03-03
weight: 100
---

# AWSControlPlane

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">awscontrolplanes.infrastructure.giantswarm.io</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">infrastructure.giantswarm.io</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">awscontrolplane</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">awscontrolplanes</dd>
<dt class="scope">Scope:</dt>
<dd class="scope">Namespaced</dd>
<dt class="versions">Versions:</dt>
<dd class="versions"><a class="version" href="#v1alpha1" title="Show schema for version v1alpha1">v1alpha1</a><a class="version" href="#v1alpha2" title="Show schema for version v1alpha2">v1alpha2</a></dd>
</dl>



<div id="v1alpha1">
<h2>Schema for version v1alpha1</h2>


<div class="property depth-0" id=".spec">
<div class="property-header">
<h3 class="property-path">.spec</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-1" id=".spec.instanceType">
<div class="property-header">
<h3 class="property-path">.spec.instanceType</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>


<h2 id="example-v1alpha1">Example CR</h2>

<p>TODO: Show example CR</p>

</div>


<div id="v1alpha2">
<h2>Schema for version v1alpha2</h2>


<div class="property depth-0" id=".spec">
<div class="property-header">
<h3 class="property-path">.spec</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-1" id=".spec.instanceType">
<div class="property-header">
<h3 class="property-path">.spec.instanceType</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-1" id=".spec.availabilityZones">
<div class="property-header">
<h3 class="property-path">.spec.availabilityZones</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.availabilityZones[*]">
<div class="property-header">
<h3 class="property-path">.spec.availabilityZones[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>


<h2 id="example-v1alpha2">Example CR</h2>

<p>TODO: Show example CR</p>

</div>
