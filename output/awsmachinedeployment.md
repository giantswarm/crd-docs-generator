---
title: AWSMachineDeployment CRD Schema Reference
linktitle: AWSMachineDeployment
technical_name: awsmachinedeployments.infrastructure.giantswarm.io
description: >
  Custom Resource/Custom Resource Definition schema reference page
  for the AWSMachineDeployment resource (awsmachinedeployments.infrastructure.giantswarm.io),
  as part of the Giant Swarm Control Plane Kubernetes API documentation.
date: 2020-03-03
weight: 100
---

# AWSMachineDeployment

<dl class="crd-meta">
<dt class="fullname">Full name:</dt>
<dd class="fullname">awsmachinedeployments.infrastructure.giantswarm.io</dd>
<dt class="groupname">Group:</dt>
<dd class="groupname">infrastructure.giantswarm.io</dd>
<dt class="singularname">Singular name:</dt>
<dd class="singularname">awsmachinedeployment</dd>
<dt class="pluralname">Plural name:</dt>
<dd class="pluralname">awsmachinedeployments</dd>
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

<div class="property depth-1" id=".spec.nodePool">
<div class="property-header">
<h3 class="property-path">.spec.nodePool</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.nodePool.description">
<div class="property-header">
<h3 class="property-path">.spec.nodePool.description</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.nodePool.machine">
<div class="property-header">
<h3 class="property-path">.spec.nodePool.machine</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.nodePool.machine.dockerVolumeSizeGB">
<div class="property-header">
<h3 class="property-path">.spec.nodePool.machine.dockerVolumeSizeGB</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.nodePool.machine.kubeletVolumeSizeGB">
<div class="property-header">
<h3 class="property-path">.spec.nodePool.machine.kubeletVolumeSizeGB</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.nodePool.scaling">
<div class="property-header">
<h3 class="property-path">.spec.nodePool.scaling</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.nodePool.scaling.max">
<div class="property-header">
<h3 class="property-path">.spec.nodePool.scaling.max</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.nodePool.scaling.min">
<div class="property-header">
<h3 class="property-path">.spec.nodePool.scaling.min</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">integer</span>


</div>

</div>
</div>

<div class="property depth-1" id=".spec.provider">
<div class="property-header">
<h3 class="property-path">.spec.provider</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">object</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.provider.availabilityZones">
<div class="property-header">
<h3 class="property-path">.spec.provider.availabilityZones</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">array</span>


</div>

</div>
</div>

<div class="property depth-3" id=".spec.provider.availabilityZones[*]">
<div class="property-header">
<h3 class="property-path">.spec.provider.availabilityZones[*]</h3>
</div>
<div class="property-body">
<div class="property-meta">
<span class="property-type">string</span>


</div>

</div>
</div>

<div class="property depth-2" id=".spec.provider.worker">
<div class="property-header">
<h3 class="property-path">.spec.provider.worker</h3>
</div>
<div class="property-body">
<div class="property-meta">



</div>

</div>
</div>

<div class="property depth-3" id=".spec.provider.worker.instanceType">
<div class="property-header">
<h3 class="property-path">.spec.provider.worker.instanceType</h3>
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
