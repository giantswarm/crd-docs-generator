---
title: G8sControlPlane CRD Schema Reference
linktitle: G8sControlPlane
technical_name: g8scontrolplanes.infrastructure.giantswarm.io
description: >
  Custom Resource/Custom Resource Definition schema reference page
  for the G8sControlPlane resource (g8scontrolplanes.infrastructure.giantswarm.io),
  as part of the Giant Swarm Control Plane Kubernetes API documentation.
date: 2020-03-03
weight: 100
---

# G8sControlPlane

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">g8scontrolplanes.infrastructure.giantswarm.io</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">infrastructure.giantswarm.io</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">g8scontrolplane</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">g8scontrolplanes</dd>
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

<div class="property depth-1" id=".spec.replicas">
<div class="property-header">
<h3 class="property-path">.spec.replicas</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">int</span>


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

<div class="property depth-1" id=".spec.infrastructureRef">
<div class="property-header">
<h3 class="property-path">.spec.infrastructureRef</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.infrastructureRef.kind">
<div class="property-header">
<h3 class="property-path">.spec.infrastructureRef.kind</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.infrastructureRef.name">
<div class="property-header">
<h3 class="property-path">.spec.infrastructureRef.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.infrastructureRef.namespace">
<div class="property-header">
<h3 class="property-path">.spec.infrastructureRef.namespace</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.infrastructureRef.apiVersion">
<div class="property-header">
<h3 class="property-path">.spec.infrastructureRef.apiVersion</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-1" id=".spec.replicas">
<div class="property-header">
<h3 class="property-path">.spec.replicas</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">int</span>


</div>

</div>
</div>


<h2 id="example-v1alpha2">Example CR</h2>

<p>TODO: Show example CR</p>

</div>
