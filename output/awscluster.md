---
title: AWSCluster CRD Schema Reference
linktitle: AWSCluster
technical_name: awsclusters.infrastructure.giantswarm.io
description: >
  Custom Resource/Custom Resource Definition schema reference page
  for the AWSCluster resource (awsclusters.infrastructure.giantswarm.io),
  as part of the Giant Swarm Control Plane Kubernetes API documentation.
date: 2020-03-03
weight: 100
---

# AWSCluster

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">awsclusters.infrastructure.giantswarm.io</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">infrastructure.giantswarm.io</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">awscluster</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">awsclusters</dd>
<dt class="scope">Scope:</dt>
<dd class="scope">Namespaced</dd>
<dt class="versions">Versions:</dt>
<dd class="versions"><a class="version" href="#v1alpha2" title="Show schema for version v1alpha2">v1alpha2</a></dd>
</dl>



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

<div class="property depth-1" id=".spec.cluster">
<div class="property-header">
<h3 class="property-path">.spec.cluster</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.cluster.description">
<div class="property-header">
<h3 class="property-path">.spec.cluster.description</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.cluster.dns">
<div class="property-header">
<h3 class="property-path">.spec.cluster.dns</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.cluster.dns.domain">
<div class="property-header">
<h3 class="property-path">.spec.cluster.dns.domain</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.cluster.dns.provider">
<div class="property-header">
<h3 class="property-path">.spec.cluster.dns.provider</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-4" id=".spec.cluster.dns.provider.master">
<div class="property-header">
<h3 class="property-path">.spec.cluster.dns.provider.master</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-5" id=".spec.cluster.dns.provider.master.availabilityZone">
<div class="property-header">
<h3 class="property-path">.spec.cluster.dns.provider.master.availabilityZone</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-5" id=".spec.cluster.dns.provider.master.instanceType">
<div class="property-header">
<h3 class="property-path">.spec.cluster.dns.provider.master.instanceType</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-4" id=".spec.cluster.dns.provider.region">
<div class="property-header">
<h3 class="property-path">.spec.cluster.dns.provider.region</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-0" id=".status">
<div class="property-header">
<h3 class="property-path">.status</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

<div class="property-description">
<p><strong>Note:</strong> The <code>.status</code> section schema has been added for demonstration purposes only.</p>

</div>

</div>
</div>

<div class="property depth-1" id=".status.cluster">
<div class="property-header">
<h3 class="property-path">.status.cluster</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-2" id=".status.cluster.conditions">
<div class="property-header">
<h3 class="property-path">.status.cluster.conditions</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>


</div>

<div class="property-description">
<p>Holds information to explain the current condition(s) of the cluster.</p>

</div>

</div>
</div>

<div class="property depth-3" id=".status.cluster.conditions[*]">
<div class="property-header">
<h3 class="property-path">.status.cluster.conditions[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

<div class="property-description">
<p>An individual condition.</p>

</div>

</div>
</div>

<div class="property depth-4" id=".status.cluster.conditions[*].lastTransitionTime">
<div class="property-header">
<h3 class="property-path">.status.cluster.conditions[*].lastTransitionTime</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

<div class="property-description">
<p>Time the cluster assumed the given condition.</p>

</div>

</div>
</div>

<div class="property depth-4" id=".status.cluster.conditions[*].type">
<div class="property-header">
<h3 class="property-path">.status.cluster.conditions[*].type</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

<div class="property-description">
<p>Condition string. Possible values:</p>

<ul>
<li><code>Creating</code> - Cluster is being created.</li>
<li><code>Created</code> - Cluster creation has been finished.</li>
<li><code>Updating</code> - Cluster is undergoing an update, e. g. for a release upgrade.</li>
<li><code>Updated</code> - Cluster update is finished.</li>
</ul>

</div>

</div>
</div>


<h2 id="example-v1alpha2">Example CR</h2>

<p>TODO: Show example CR</p>

</div>
