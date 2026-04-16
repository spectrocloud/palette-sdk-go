# Cluster Provisioning — Deep Reference

## Orchestration Modes

Palette uses three orchestration modes to deploy clusters. The mode is determined by the cloud type and target environment.

### Controller Mode (Public Clouds & Data Centers)
- Default for: AWS, Azure, GCP, EKS, AKS, GKE, VMware, MaaS
- Uses Palette's CAPI-based (Cluster API) orchestrator
- Palette manages the full cluster lifecycle via Cluster API controllers
- Infrastructure is provisioned in the target cloud/data center
- Requires a Cloud Account with valid credentials for the target provider
- Palette communicates with the cloud provider API to create VMs/nodes, then bootstraps Kubernetes

### Appliance Mode (Edge — Immutable OS)
- Used for: Edge environments with Palette-shipped appliance images
- Built on Kairos immutable OS toolkit
- Pre-built appliance images are installed on physical or virtual bare-metal machines
- The OS is immutable — updates are delivered as new image layers, not in-place patches
- Machines are represented as "Edge Hosts" in Palette
- Suitable for air-gapped, remote, and factory-floor deployments

### Agent Mode (Edge — Customer-Managed Hosts)
- Used for: Edge environments where customer manages the host OS
- A lightweight Palette agent is installed on existing customer-managed machines
- The agent registers the host with Palette and participates in cluster orchestration
- More flexible than Appliance Mode — works with customer's existing OS and tooling
- Machines are also represented as "Edge Hosts"

## Provisioning Flow (Controller Mode)

1. User selects cloud type and provider
2. User selects or creates a Cloud Account (credentials)
3. User selects or creates a Cluster Profile (blueprint)
4. User provides cluster configuration (name, node pools, region, etc.)
5. Palette calls the cloud provider API to provision infrastructure
6. Palette bootstraps Kubernetes using the selected distribution (PXK, K3S, RKE2)
7. Palette installs add-on layers from the Cluster Profile
8. Cluster reaches "Running" state

## Supported K8s Distributions

| Distribution | Description | Use case |
|-------------|-------------|----------|
| PXK | Palette eXtended Kubernetes — Palette's upstream-aligned distribution | Default for cloud/data center |
| PXK-E | PXK Edge variant | Edge deployments |
| K3S | Lightweight Kubernetes by Rancher | Edge, resource-constrained environments |
| RKE2 | Rancher Kubernetes Engine 2 | Data center, security-focused |

## Provider-Specific Notes

**AWS (IaaS):** Provisions EC2 instances. Requires access_key + secret_key. Supports STS assume-role.

**EKS (Managed):** Uses Amazon's managed control plane. Provisions worker node groups. Same AWS credentials but creates EKS-specific resources.

**Azure (IaaS):** Provisions Azure VMs. Requires client_id, client_secret, tenant_id, subscription_id.

**AKS (Managed):** Azure's managed Kubernetes. Same Azure credentials.

**GCP (IaaS):** Provisions Compute Engine VMs. Requires service account JSON credentials.

**GKE (Managed):** Google's managed Kubernetes. Same GCP credentials.

**VMware (Data Center):** Provisions VMs on vSphere. Requires vcenter_server, username, password.

**MaaS (Bare Metal):** Provisions on bare-metal machines managed by Canonical MaaS. Requires api_endpoint, api_key.

**Edge Native:** No cloud account needed. Uses pre-registered Edge Hosts.
