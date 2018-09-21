---
title: Application Connector components
type: Architecture
---

The Application Connector consists of the following components:

* **Ingress-Nginx controller** responsible for exposing Application Connectors to the external world in secured way.
* **Connector service** responsible for establishing connection to Kyma and exchaning the client certificates.
* **Metadata service** responsible for registration of external API and Event Catalogs.
* **Event service** responsible for delivering events to Kyma.
* **Proxy service** responsible for proxying calls to the registered solution's APIs.
* **Remote Environment controller** responsible for manipulation Remote Environments.


![Architecture Diagram](assets/001-application-connector.png)
