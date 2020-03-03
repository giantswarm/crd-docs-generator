---
title: Release CRD Schema Reference
linktitle: Release
technical_name: releases.release.giantswarm.io
description: >
  Custom Resource/Custom Resource Definition schema reference page
  for the Release resource (releases.release.giantswarm.io),
  as part of the Giant Swarm Control Plane Kubernetes API documentation.
date: 2020-03-03
weight: 100
---

# Release

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">releases.release.giantswarm.io</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">release.giantswarm.io</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">release</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">releases</dd>
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

<div class="property depth-1" id=".spec.changelog">
<div class="property-header">
<h3 class="property-path">.spec.changelog</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-2" id=".spec.changelog[*]">
<div class="property-header">
<h3 class="property-path">.spec.changelog[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.changelog[*].component">
<div class="property-header">
<h3 class="property-path">.spec.changelog[*].component</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-3" id=".spec.changelog[*].description">
<div class="property-header">
<h3 class="property-path">.spec.changelog[*].description</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-3" id=".spec.changelog[*].kind">
<div class="property-header">
<h3 class="property-path">.spec.changelog[*].kind</h3>
</div>
<div class="property-body">
<div class="property-meta">


<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-1" id=".spec.components">
<div class="property-header">
<h3 class="property-path">.spec.components</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-2" id=".spec.components[*]">
<div class="property-header">
<h3 class="property-path">.spec.components[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.components[*].name">
<div class="property-header">
<h3 class="property-path">.spec.components[*].name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.components[*].version">
<div class="property-header">
<h3 class="property-path">.spec.components[*].version</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-1" id=".spec.parentVersion">
<div class="property-header">
<h3 class="property-path">.spec.parentVersion</h3>
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

<div class="property depth-0" id=".status">
<div class="property-header">
<h3 class="property-path">.status</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-1" id=".status.cycle">
<div class="property-header">
<h3 class="property-path">.status.cycle</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-2" id=".status.cycle.disabledDate">
<div class="property-header">
<h3 class="property-path">.status.cycle.disabledDate</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".status.cycle.enabledDate">
<div class="property-header">
<h3 class="property-path">.status.cycle.enabledDate</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".status.cycle.phase">
<div class="property-header">
<h3 class="property-path">.status.cycle.phase</h3>
</div>
<div class="property-body">
<div class="property-meta">


<span class="property-required">Required</span>

</div>

</div>
</div>


<h2 id="example-v1alpha1">Example CR</h2>

<p>TODO: Show example CR</p>

</div>
