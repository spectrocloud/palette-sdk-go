# Cluster Profiles — Deep Reference

## What is a Cluster Profile?

A Cluster Profile is a **blueprint** that defines the complete software stack for a Kubernetes cluster. It is a layered model where each layer represents a specific infrastructure or application concern.

## Profile Layers (Stack Order)

```
┌─────────────────────────┐
│   Add-On Layers         │  ← Monitoring, logging, ingress, apps (Helm charts, manifests)
├─────────────────────────┤
│   CSI (Storage)         │  ← AWS EBS, Azure Disk, vSphere CSI, Longhorn, etc.
├─────────────────────────┤
│   CNI (Networking)      │  ← Calico, Cilium, Flannel, etc.
├─────────────────────────┤
│   Kubernetes            │  ← PXK, K3S, RKE2, PXK-E
├─────────────────────────┤
│   OS                    │  ← Ubuntu, RHEL, openSUSE
└─────────────────────────┘
```

Each layer is a **Pack** (or Helm chart for add-ons).

## Profile Types

### Infra Profile
- Contains: OS + Kubernetes + CNI + CSI layers
- **Cloud-type-specific** — an AWS infra profile cannot be used on Azure
- One infra profile is required for every cluster

### Add-On Profile
- Contains: Application layers only (above CSI)
- **Cloud-type agnostic** — can be applied to any cluster
- Zero or more add-on profiles can be stacked on a cluster
- Built from Helm charts, Packs, or raw manifests

### Full Profile
- Contains: Infra + Add-On layers in a single profile
- Convenience type for standardized full-stack deployments

## Profile Lifecycle

1. **Create** — Define layers by selecting Packs
2. **Publish** — Make the profile available for cluster creation (draft → published)
3. **Version** — Profiles are versioned. Clusters pin to a specific version.
4. **Update** — Modify layers. Creates a new version.
5. **Import/Export** — Profiles can be exported as content strings and imported into other projects/tenants
6. **Delete** — Remove a profile (only if not in use by any cluster)

## Pack Selection

When building a profile, each layer requires selecting a Pack:

| Layer | Available Packs (examples) |
|-------|---------------------------|
| OS | Ubuntu 22.04, Ubuntu 24.04, RHEL 8, RHEL 9, openSUSE |
| Kubernetes | PXK 1.28, PXK 1.29, K3S 1.28, RKE2 1.28 |
| CNI | Calico, Cilium, Flannel, Antrea |
| CSI | AWS EBS CSI, Azure Disk CSI, vSphere CSI, Longhorn, Rook-Ceph |
| Add-On | Prometheus, Grafana, Nginx Ingress, Cert-Manager, ArgoCD, etc. |

## Profile Variables

- Profiles support **variables** — parameterized values that can be overridden at cluster creation time
- Example: a profile might define `replicas` variable, and each cluster can set its own replica count
- Variables are defined in the profile and resolved at deployment time
