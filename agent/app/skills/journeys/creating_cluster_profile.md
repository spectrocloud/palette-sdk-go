# Creating Cluster Profile

Create a new infra cluster profile with a single form — packs are auto-fetched and pre-filtered

## Defaults

**AWS IaaS defaults:**
- os: `ubuntu-aws`
- k8s: `kubernetes`
- cni: `cni-calico`
- csi: `csi-aws-ebs`
- ssh_key_name: ``
- region: `us-east-1`
- instance_type: `t3.large`
- node_count: `3`

**Azure IaaS defaults:**
- os: `ubuntu-azure`
- k8s: `kubernetes`
- cni: `cni-calico-azure`
- csi: `csi-azure`
- region: `eastus`
- instance_type: `Standard_D2s_v3`
- node_count: `3`

**Google Cloud IaaS defaults:**
- cni: `cni-calico`
- csi: `csi-gcp-driver`
- os: `ubuntu-gcp`
- k8s: `kubernetes`
- region: `us-central1`
- machine_type: `n1-standard-4`
- node_count: `3`

**VMware vSphere defaults:**
- k8s: `kubernetes`
- cni: `cni-calico`
- csi: `csi-vsphere-csi`
- os: `ubuntu-vsphere`
- memory_mb: `8192`
- disk_gb: `60`
- node_count: `3`
- datacenter: ``
- folder: ``
- resource_pool: ``
- datastore: ``
- network: ``
- num_cpus: `4`

**Edge Native defaults:**
- os: `edge-native-byoi`
- k8s: `edge-k3s`
- cni: `cni-calico`
- csi: `csi-longhorn`
- node_count: `1`
- single_node: `true`
- vip: ``

**MAAS Bare Metal defaults:**
- csi: `longhorn`
- os: `ubuntu-maas`
- k8s: `kubernetes`
- cni: `cni-calico`
- node_count: `3`

**Amazon EKS defaults:**
- os: `amazon-linux-eks`
- k8s: `kubernetes-eks`
- cni: `cni-aws-vpc-eks`
- csi: `csi-aws-ebs`
- region: `us-east-1`
- instance_type: `t3.large`
- node_count: `3`
- endpoint_access: `public`
- kubernetes_version: `1.28`

**Global:** preferred registry = `Palette Repo`

---

**API skills used:** `profile_api`

## How to Choose a Variant

**Ask:** What cloud type is this profile for?

- **eks**: user mentions "eks"
- **aws**: user mentions "aws"
- **azure**: user mentions "azure", "aks"
- **gcp**: user mentions "gcp", "gke"
- **vmware**: user mentions "vmware", "vsphere"
- **edge**: user mentions "edge"
- **maas**: user mentions "maas"
- **Default**: ask

---

## Gcp

Create a GCP infra profile

### Steps

1. **Build GCP profile**
   `profile_api` with `{"build_profile": "gcp"}`

---

## Vmware

Create a VMware infra profile

### Steps

1. **Build VMware profile**
   `profile_api` with `{"build_profile": "vmware"}`

---

## Edge

Create an Edge Native infra profile

### Steps

1. **Build Edge profile**
   `profile_api` with `{"build_profile": "edge-native"}`

---

## Maas

Create a MAAS infra profile

### Steps

1. **Build MAAS profile**
   `profile_api` with `{"build_profile": "maas"}`

---

## Ask

Cloud type not specified — ask first, then build

### Steps

1. **Ask the user which cloud type, then trigger the profile builder**
   `profile_api` with `{"collect": {"_title": "New Cluster Profile", "cloud_type": {"type": "select", "label": "Cloud Type", "required": true, "enum": ["eks", "aws", "azure", "aks", "gcp", "gke", "vmware", "edge-native", "maas"]}}}`
   After collecting the cloud type, call profile_api with build_profile set to the cloud type.

2. **Build the profile**
   Tool: `profile_api`
   Call with: {"build_profile": "<cloud_type from previous step>"}. This fetches available packs, shows a single form with pre-filtered dropdowns and smart defaults, and creates + publishes the profile.

---

## Eks

Create an EKS infra profile

### Steps

1. **Build EKS profile with auto-fetched packs and defaults**
   `profile_api` with `{"build_profile": "eks"}`
   This single call fetches all available EKS packs, presents a form with OS/K8s/CNI/CSI dropdowns (pre-filtered for EKS, with best defaults selected), collects the user's choices, creates the profile, and publishes it.

### After Completion

- Profile is ready for EKS cluster deployments
- Deploy a cluster with: 'deploy an EKS cluster'

---

## Aws

Create an AWS IaaS infra profile

### Steps

1. **Build AWS profile**
   `profile_api` with `{"build_profile": "aws"}`

---

## Azure

Create an Azure infra profile

### Steps

1. **Build Azure profile**
   `profile_api` with `{"build_profile": "azure"}`

---

## Notes

- build_profile triggers a smart form that auto-fetches packs filtered for the cloud type
- Defaults are pre-selected — user just reviews and submits
- Profile is automatically published after creation
- The profile builder shows one form with all 4 layers (OS, K8s, CNI, CSI) as dropdowns

