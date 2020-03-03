---
title: Chart CRD Schema Reference
linktitle: Chart
technical_name: charts.application.giantswarm.io
description: >
  Custom Resource/Custom Resource Definition schema reference page
  for the Chart resource (charts.application.giantswarm.io),
  as part of the Giant Swarm Control Plane Kubernetes API documentation.
date: 2020-03-03
weight: 100
---

# Chart

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">charts.application.giantswarm.io</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">application.giantswarm.io</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">chart</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">charts</dd>
<dt class="scope">Scope:</dt>
<dd class="scope">Namespaced</dd>
<dt class="versions">Versions:</dt>
<dd class="versions"><a class="version" href="#v1alpha1" title="Show schema for version v1alpha1">v1alpha1</a></dd>
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

<div class="property depth-1" id=".spec.config">
<div class="property-header">
<h3 class="property-path">.spec.config</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.config.configMap">
<div class="property-header">
<h3 class="property-path">.spec.config.configMap</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.config.configMap.name">
<div class="property-header">
<h3 class="property-path">.spec.config.configMap.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-3" id=".spec.config.configMap.namespace">
<div class="property-header">
<h3 class="property-path">.spec.config.configMap.namespace</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-3" id=".spec.config.configMap.resourceVersion">
<div class="property-header">
<h3 class="property-path">.spec.config.configMap.resourceVersion</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.config.secret">
<div class="property-header">
<h3 class="property-path">.spec.config.secret</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.config.secret.name">
<div class="property-header">
<h3 class="property-path">.spec.config.secret.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-3" id=".spec.config.secret.namespace">
<div class="property-header">
<h3 class="property-path">.spec.config.secret.namespace</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-3" id=".spec.config.secret.resourceVersion">
<div class="property-header">
<h3 class="property-path">.spec.config.secret.resourceVersion</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-1" id=".spec.name">
<div class="property-header">
<h3 class="property-path">.spec.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-1" id=".spec.namespace">
<div class="property-header">
<h3 class="property-path">.spec.namespace</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-1" id=".spec.tarballURL">
<div class="property-header">
<h3 class="property-path">.spec.tarballURL</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-1" id=".spec.version">
<div class="property-header">
<h3 class="property-path">.spec.version</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>


<h2 id="example-v1alpha1">Example CR</h2>

<p>TODO: Show example CR</p>

</div>
