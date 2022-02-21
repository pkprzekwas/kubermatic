# **KCP**

**Author**: Patryk Przekwas (@pkprzekwas)

**Status**: Draft proposal; research.

## Goals

The goal of this document is an evaluation of [kcp](https://github.com/kcp-dev/kcp). It investigates the fitness of the tool for Kubermatic, focusing on three main areas:
* minimal api-server
* multi-tenancy
* multi-clustering

## Non-Goals

This document will not suggest concrete solutions. `kcp` is too premature, there were no official releases and documentation hasn't been updated for a while (9 months). Nonetheless, it's worth keeping an eye on the project as it's trending pretty rapidly and development is moving forward.

## Motivation and Background

`kcp` is a minimal Kubernetes api-server. It doesn't know about `pods`, `deployments`, `services`, or even `nodes`. The tool was created as a base for all applications leveraging the reconciliation pattern. Some resources like `namespaces` or `serviceAccounts` have been preserved for multi-tenancy purposes.

The project recently gained some interest in the Kubernetes community. Engineers from RedHat seem to be pushing it forward. There were other similar projects in the past (e.g. [badidea](https://github.com/thetirefire/badidea)) proving that there is some need for such a solution.

`kcp` is being sold as a [Minimal API-Server](https://github.com/kcp-dev/kcp/blob/main/docs/investigations/minimal-api-server.md) but goes further than that. Authors introduced more components to it, namely: `cluster-controller`, `syncer` and `deployment-splitter`. All tools combined create an opinionated framework to resolve multi-tenancy and multi-clustering problems in today's Kubernetes ecosystem. The `Clusters` CRD lets users create so-called [Logical Clusters](https://github.com/kcp-dev/kcp/blob/main/docs/investigations/logical-clusters.md), which seems to be the main abstraction of `kcp`. They hold information to reach and authenticate to another Kubernetes cluster's api-server. The `cluster-controller` watches for any `Cluster` resources to connect to them and install `syncer`. It also watches the cluster's API resources to discover new types and updates to existing types. The `cluster-controller` uses this information to negotiate possible incompatible CRD type definitions, to determine whether an incoming resource can be sent to a cluster's `syncer`. The `syncer` maintains a connection between a `kcp` instance and a Kubernetes cluster's api-server. After initial type negotiation, the `syncer` watches for resources of all types that are scheduled to that cluster. At last but not least, `deployment-splitter` is an example of a very simple multi-cluster resource scheduler. It watches for `Deployment` resources in the `kcp` and determines how many of that Deployment's replicas should be scheduled to which of the available Cluster resources. It's not meant for real production scenarios. Instead, it presents an idea of having [Transparent Multi-Clusters](https://github.com/kcp-dev/kcp/blob/main/docs/investigations/transparent-multi-cluster.md). In short, the key area of investigation for `kcp` is exploring transparency of workloads to clusters.

## Implementation

As mentioned before, it's too early to consider any serious integrations with `KKP`. However, I want to list here some likely-looking options:
- installing `kcp` on `KKP` as an addon to `user-clusters` letting end-user manage multi-cluster workflows using `kcp`
- using minimal api-server as a base for testing all node independent CRD controllers
- using `kcp` [api-server as a library](https://github.com/kcp-dev/kcp/blob/main/DEVELOPMENT.md#using-kcp-as-a-library) letting us deploy server and controllers as a single binary
- multi-tenancy model proposed by `kcp`, based on logical clusters works only with all of the previously mentioned components, hence it cannot be used to solve any of the current `KKP` problems

## Alternatives considered

While investigating `kcp` in terms of its multi-tenancy capabilities, I found a reference to an interesting work describing current efforts in the official K8s multi-tenancy SIG - [A Multi Tenant Framework for Cloud Container Services](https://github.com/kubernetes-sigs/multi-tenancy/blob/master/incubator/virtualcluster/doc/vc-icdcs.pdf).

## Task & effort:

For now, I suggest waiting till all ideas clarify. It looks like the authors want to solve too many problems at the same time. I wish they decoupled them into smaller chunks e.g. release minimal api-server separately so that community could fork it and try some other approaches.

I suggest checking the status of the project every 3-4 months to get a sense of the direction it is heading in.

