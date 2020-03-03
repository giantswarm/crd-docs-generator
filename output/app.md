---
title: App CRD Schema Reference
linktitle: App
technical_name: apps.application.giantswarm.io
description: >
  Custom Resource/Custom Resource Definition schema reference page
  for the App resource (apps.application.giantswarm.io),
  as part of the Giant Swarm Control Plane Kubernetes API documentation.
date: 2020-03-03
weight: 100
---

# App

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">apps.application.giantswarm.io</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">application.giantswarm.io</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">app</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">apps</dd>
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

<div class="property depth-1" id=".spec.catalog">
<div class="property-header">
<h3 class="property-path">.spec.catalog</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

<div class="property-description">
<p>Name of the catalog this app is found in.</p>

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

<div class="property-description">
<p>Optional information for looking up a ConfigMap resource. This can be used to configure the installed app.</p>

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

<div class="property-description">
<p>Name of the ConfigMap resource.</p>

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

<div class="property-description">
<p>Namespace the ConfigMap resource can be found in.</p>

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

<div class="property-description">
<p>Optional information for looking up a Secret resource, which can be used by the installed app.</p>

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

<div class="property-description">
<p>Name of the Secret resource.</p>

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

<div class="property-description">
<p>Namespace the Secret resource can be found in.</p>

</div>

</div>
</div>

<div class="property depth-1" id=".spec.kubeConfig">
<div class="property-header">
<h3 class="property-path">.spec.kubeConfig</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.kubeConfig.context">
<div class="property-header">
<h3 class="property-path">.spec.kubeConfig.context</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.kubeConfig.context.name">
<div class="property-header">
<h3 class="property-path">.spec.kubeConfig.context.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.kubeConfig.inCluster">
<div class="property-header">
<h3 class="property-path">.spec.kubeConfig.inCluster</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">boolean</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.kubeConfig.secret">
<div class="property-header">
<h3 class="property-path">.spec.kubeConfig.secret</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.kubeConfig.secret.name">
<div class="property-header">
<h3 class="property-path">.spec.kubeConfig.secret.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-3" id=".spec.kubeConfig.secret.namespace">
<div class="property-header">
<h3 class="property-path">.spec.kubeConfig.secret.namespace</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

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

<div class="property-description">
<p>Name of the app in the catalog.</p>

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

<div class="property-description">
<p>Namespace to install the actuall app workloads to.</p>

</div>

</div>
</div>

<div class="property depth-1" id=".spec.userConfig">
<div class="property-header">
<h3 class="property-path">.spec.userConfig</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.userConfig.configMap">
<div class="property-header">
<h3 class="property-path">.spec.userConfig.configMap</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.userConfig.configMap.name">
<div class="property-header">
<h3 class="property-path">.spec.userConfig.configMap.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-3" id=".spec.userConfig.configMap.namespace">
<div class="property-header">
<h3 class="property-path">.spec.userConfig.configMap.namespace</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-2" id=".spec.userConfig.secret">
<div class="property-header">
<h3 class="property-path">.spec.userConfig.secret</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.userConfig.secret.name">
<div class="property-header">
<h3 class="property-path">.spec.userConfig.secret.name</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>

<span class="property-required">Required</span>

</div>

</div>
</div>

<div class="property depth-3" id=".spec.userConfig.secret.namespace">
<div class="property-header">
<h3 class="property-path">.spec.userConfig.secret.namespace</h3>
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

<div class="property-description">
<p>App version number to install from the catalog.</p>

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
<p>Just a fake entry to test a few things.</p>

</div>

</div>
</div>

<div class="property depth-1" id=".status.arrayProp">
<div class="property-header">
<h3 class="property-path">.status.arrayProp</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>


</div>

<div class="property-description">
<p>Here is an example property of type Array.</p>

</div>

</div>
</div>

<div class="property depth-2" id=".status.arrayProp[*]">
<div class="property-header">
<h3 class="property-path">.status.arrayProp[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

<div class="property-description">
<p>Example array items.</p>

</div>

</div>
</div>

<div class="property depth-3" id=".status.arrayProp[*].bar">
<div class="property-header">
<h3 class="property-path">.status.arrayProp[*].bar</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-3" id=".status.arrayProp[*].foo">
<div class="property-header">
<h3 class="property-path">.status.arrayProp[*].foo</h3>
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
