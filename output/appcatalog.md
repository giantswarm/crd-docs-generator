---
title: AppCatalog CRD Schema Reference
linktitle: AppCatalog
technical_name: appcatalogs.application.giantswarm.io
description: >
  Custom Resource/Custom Resource Definition schema reference page
  for the AppCatalog resource (appcatalogs.application.giantswarm.io),
  as part of the Giant Swarm Control Plane Kubernetes API documentation.
date: 2020-03-03
weight: 100
---

# AppCatalog

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">appcatalogs.application.giantswarm.io</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">application.giantswarm.io</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">appcatalog</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">appcatalogs</dd>
<dt class="scope">Scope:</dt>
<dd class="scope">Cluster</dd>
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

<div class="property depth-1" id=".spec.description">
<div class="property-header">
<h3 class="property-path">.spec.description</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-1" id=".spec.logoURL">
<div class="property-header">
<h3 class="property-path">.spec.logoURL</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-1" id=".spec.storage">
<div class="property-header">
<h3 class="property-path">.spec.storage</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-2" id=".spec.storage.URL">
<div class="property-header">
<h3 class="property-path">.spec.storage.URL</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-2" id=".spec.storage.type">
<div class="property-header">
<h3 class="property-path">.spec.storage.type</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-1" id=".spec.title">
<div class="property-header">
<h3 class="property-path">.spec.title</h3>
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
