# Deploying EKS Cluster

End-to-end journey for deploying an EKS cluster

**API skills used:** `eks_cluster_api`, `cloud_account_api`, `profile_api`, `cluster_common_api`

## Shared Sub-Journeys

These reusable patterns are used in the steps below:

### Ensuring Cloud Account
Ensure a cloud account exists — list existing, let user pick, or create new with a form

- **List existing cloud accounts**
  Present accounts as a pick list. If user picks an existing account, use it. If no accounts exist, proceed to create one.
- **If no accounts found — create one using the account builder**
  Call with {"build_account": "<provider>"}. This shows a single form with the right credential fields for the provider (AWS: access key + secret, Azure: tenant/client/subscription, GCP: JSON key). The account is created automatically.
  *This is a shared sub-journey — referenced by deploy journeys that need a cloud account*
  *build_account handles: form building, credential collection, and account creation — all in one step*
  *For EKS/AKS/GKE journeys, the provider is mapped automatically (EKS→AWS, AKS→Azure, GKE→GCP)*
  *If accounts exist and user picks one, SKIP the build step*

### Ensuring Cluster Profile
Ensure a cluster profile exists — list existing, let user pick, or create new with a single form

- **List existing infra profiles for the target cloud type**
  Use search_profiles with interactive mode. Filter by the target environment and profileType 'cluster'. Example: {"mode": "search_profiles", "filter": {"filter": {"environment": ["eks"], "profileType": ["cluster"]}}, "interactive": true}. If profiles exist, let user pick and SKIP the creation step below.
- **If no profiles found — create one using the profile builder**
  Call with {"build_profile": "<cloud_type>"}. This shows a single form with pre-filtered pack dropdowns and smart defaults. The user reviews, adjusts if needed, and submits. The profile is created and published automatically.
  *This is a shared sub-journey — referenced by deploy journeys that need a cluster profile*
  *If profiles exist and user picks one, SKIP the build step*
  *build_profile handles: pack fetching, form building, defaults, creation, and publishing — all in one step*

---

## How to Choose a Variant

**Ask:** What kind of network environment?

- **airgapped**: user mentions "airgapped", "disconnected", "air-gapped", "isolated", "no internet", "air gap"
- **connected**: user mentions "private subnet", "NAT", "connected", "standard", "default", "public"
- **Default**: connected

---

## Connected

Standard EKS deployment with internet access

### Steps

1. **Ensure AWS cloud account exists**
   *Uses shared sub-journey: ensure_cloud_account*
   Context: Provider is AWS. The user needs an AWS cloud account for EKS deployment.
   List AWS accounts with interactive mode. If accounts exist, let user pick. If none exist, collect AWS credentials (name, access_key_id, secret_access_key) via a form, then create the account. The selected or created account's UID is needed for cluster creation.

2. **Ensure EKS-compatible cluster profile exists**
   *Uses shared sub-journey: ensure_cluster_profile*
   Context: Cloud type is EKS. The user needs an infra profile with cloud_type=eks.
   Search profiles filtered by environment=eks and profileType=cluster with interactive mode. If profiles exist, let user pick. If none exist, tell the user to create one in the Palette UI (profile creation requires pack selection which is not yet automated). The selected profile's UID is needed for cluster creation.

3. **Collect cluster configuration**
   `eks_cluster_api` with `{"collect": {"_title": "EKS Cluster Configuration", "cluster_name": {"type": "string", "label": "Cluster Name", "required": true}, "region": {"type": "string", "label": "AWS Region", "required": true, "default": "us-east-1"}, "instance_type": {"type": "string", "label": "Node Instance Type", "default": "t3.large"}, "node_count": {"type": "number", "label": "Node Count", "default": "3"}}}`
   Show a form to collect cluster details.

4. **Create the EKS cluster**
   Tool: `eks_cluster_api`
   Build the full create_eks input from: the cloud account UID (from step 1), the profile UID (from step 2), and the collected config (from step 3). Call with mode create_eks.

### After Completion

- Cluster provisioning is asynchronous — takes 10-15 minutes
- Monitor status: cluster_common_api with {"mode": "get_cluster_summary", "uid": "<cluster_uid>"}
- Get kubeconfig after Running: cluster_common_api with {"mode": "get_client_kubeconfig", "uid": "<cluster_uid>"}

---

## Airgapped

EKS deployment in an isolated network without internet

### Steps

1. **Ensure AWS cloud account exists**
   *Uses shared sub-journey: ensure_cloud_account*
   Context: Provider is AWS.
   Same as connected variant.

2. **Ensure EKS-compatible cluster profile exists**
   *Uses shared sub-journey: ensure_cluster_profile*
   Context: Cloud type is EKS.
   Same as connected variant.

3. **Ensure private ECR registry exists for airgapped images**
   *Uses shared sub-journey: ensure_registry*
   Context: Registry type is ECR. Required for airgapped deployment.
   List OCI registries with interactive mode. If none exist, collect ECR details (name, access_key, secret_key, region) and create. The registry is needed for image-swap configuration.

4. **Configure image-swap on the profile**
   Image-swap pack must be added to the cluster profile to redirect image pulls to the private ECR mirror. This is currently a manual step in the Palette UI.
   *(manual step — instruct the user)*

5. **Collect cluster and airgap configuration**
   `eks_cluster_api` with `{"collect": {"_title": "EKS Airgapped Cluster Configuration", "cluster_name": {"type": "string", "label": "Cluster Name", "required": true}, "region": {"type": "string", "label": "AWS Region", "required": true, "default": "us-east-1"}, "instance_type": {"type": "string", "label": "Node Instance Type", "default": "t3.large"}, "node_count": {"type": "number", "label": "Node Count", "default": "3"}, "http_proxy": {"type": "string", "label": "HTTP Proxy URL"}, "https_proxy": {"type": "string", "label": "HTTPS Proxy URL"}, "no_proxy": {"type": "string", "label": "No Proxy List"}}}`

6. **Create the EKS cluster with airgap configuration**
   Tool: `eks_cluster_api`
   Build the create_eks input including proxy config from collected data.

### After Completion

- Verify all pods pull images from ECR mirror
- Confirm proxy allows traffic to EKS API endpoint and ECR registry
- Monitor status: cluster_common_api with {"mode": "get_cluster_summary", "uid": "<cluster_uid>"}

---

## Notes

- EKS uses AWS managed control plane — Palette provisions and manages worker nodes
- Cross-account: use STS assume-role in the cloud account
- Profile must have cloud_type: eks
- When tools are called with 'interactive': true, they present a pick list and wait for user selection
- When tools are called with 'collect': {fields}, they show a form and wait for user input
- Steps with 'use: ensure_*' reference shared sub-journeys that handle 'pick or create' logic

