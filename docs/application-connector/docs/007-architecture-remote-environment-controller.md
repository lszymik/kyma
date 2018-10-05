---
title: Remote Environment Controller
type: Details
---

The Remote Environment Controller is responsible for provisioning and deprovisioning Event and Proxy services.       
It watches Remote Environment CRD and reacts appropriately to the event of adding or removing CRD instances.

## Implementation details

Remote Environment Controller's repository includes [this](https://kyma-project.io/) chart containing Event and Proxy Service. The chart is installed or uninstalled with Helm client.