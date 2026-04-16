# Palette — Core Glossary

**Palette:** Enterprise Kubernetes management platform by Spectro Cloud. Provisions and manages clusters across clouds, data centers, and edge. Available as Palette (commercial) or Palette VerteX (FIPS-enabled).

**Tenant:** Isolated partition within a Palette installation. All resources live inside a Tenant.

**Project:** Logical partition within a Tenant for organizing infrastructure. Contains Clusters, Cluster Profiles, Cloud Accounts (project-scoped), and Registries. Most API operations are project-scoped.

**User:** Platform user within a Tenant. Can be local or synced via OIDC.

**Team:** Collection of Users. Allows bulk role assignment. Can sync from OIDC provider.

**Role:** Set of permissions. Applied at tenant level (org-wide) or project level.

**API Key:** Token for programmatic access. Tenant-scoped, one user can have multiple keys.

**Cloud Type:** Target deployment environment — AWS, Azure, GCP (IaaS); EKS, AKS, GKE (managed K8s); VMware, MaaS (data center); Edge Native.

**Cloud Account:** Credentials for a specific cloud provider. Scoped to Tenant or Project. Required for cluster provisioning.

**Pack:** Spectro Cloud's proprietary format for one layer of a full-stack system. Building block for Cluster Profiles. Hosted in Pack Registries.

**Registry:** Repository for Packs, Helm charts, or Zarf packages. Types: Pack Registry, Helm Registry, Zarf Registry. Project-scoped.

**Cluster Profile:** Blueprint for clusters. Built from Packs. Types: Infra (OS+K8s+CNI+CSI), Add-On (apps/charts/manifests), Full (both). Infra profiles are cloud-type-specific.

**Cluster:** Kubernetes cluster deployed to a supported environment. Built from one Infra profile (required) + Add-On profiles, or one Full profile.

**Scope Hierarchy:** Platform → Tenant → Project → Resources (Clusters, Profiles, etc.)
