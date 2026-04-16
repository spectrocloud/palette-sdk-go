# Deploying Edge Cluster Appliance

Deploy an edge cluster using appliance mode — launches AWS VMs from CanvOS AMI, registers hosts, creates cluster

**API skills used:** `edge_api`, `edge_cluster_api`, `aws_infra_api`, `profile_api`

## Shared Sub-Journeys

These reusable patterns are used in the steps below:

### Ensuring Edge Registration Token
Get or create an edge registration token for the tenant

- **Create a registration token**
  Creates a new token. Returns token_uid and token_value. The token_value is needed for edge host user-data.
  *Edge registration tokens are tenant-scoped*
  *The token_value is a secret — used in user-data to register edge hosts with Palette*
  *A single token can be reused for multiple edge hosts*

### Preparing AWS Edge Host Appliance
Launch AWS EC2 instances from a CanvOS appliance AMI for edge clusters

- **Collect appliance AMI ID**
  The user must provide the AMI ID of the CanvOS appliance image built via Edge Forge. This is customer-specific.
- **Select VPC for edge hosts**
- **Select subnet**
  Use list_subnets with the selected VPC
- **Select SSH key pair**
- **Launch EC2 instances with appliance AMI and user-data**
  Launch instances using the CanvOS AMI. The appliance user-data is in cloud-config YAML format with stylus site configuration.
  *Requires a pre-built CanvOS appliance AMI (built via Edge Forge process)*
  *The AMI contains the immutable Kairos OS + K8s distribution*
  *User-data is cloud-config YAML format (not bash script)*
  *Instance type should be t3.xlarge or larger for production*
  *Disk size should be 100GB+ for appliance mode*

### Registering Edge Hosts
Wait for edge hosts to register with Palette and verify their status

- **List registered edge hosts**
  Check if expected number of hosts have registered. If using tags, filter with get_edge_hosts_by_tags.
- **Verify all hosts are healthy**
  Each host should show status=registered and health=healthy. If hosts are missing, they may still be booting and installing the agent.
  *Edge hosts register automatically after the stylus agent starts*
  *Registration typically takes 2-5 minutes after VM boot*
  *If hosts don't appear, check: network connectivity, token validity, user-data correctness*
  *Use tags to filter hosts belonging to a specific deployment*

---

## Default

Appliance mode edge cluster on AWS

### Prerequisites

1. **CanvOS appliance AMI must be built and available in the target AWS region** (manual)
   If missing: Build the appliance ISO using the CanvOS Edge Forge process, then create an AMI from it.

### Steps

1. **Create edge registration token**
   *Uses shared sub-journey: ensure_edge_registration_token*

2. **Ensure edge cluster profile exists**
   *Uses shared sub-journey: ensure_edge_cluster_profile*
   Context: Cloud type is edge-native, appliance mode. The OS layer must have system:uri pointing to the provider image built by CanvOS.

3. **Collect deployment configuration**
   Collect: CanvOS AMI ID, number of hosts, CP/worker counts, VIP IP, AWS region, site name/tags

4. **Prepare AWS edge hosts with appliance AMI**
   *Uses shared sub-journey: prepare_aws_edge_host_appliance*
   Context: Launch N EC2 instances from the CanvOS AMI with appliance mode user-data (cloud-config YAML).

5. **Wait for hosts to register**
   *Uses shared sub-journey: register_edge_hosts*

6. **Create edge cluster**
   Tool: `edge_cluster_api`
   Create the edge native cluster with profile, VIP, and edge host assignments.

### After Completion

- Appliance mode clusters have immutable OS — updates are delivered as new image layers
- OS upgrades are managed by building new provider images via CanvOS
- The system:uri in the OS layer profile pack controls which provider image is used

---

## Notes

- Appliance mode: Kairos-based immutable OS, managed by Palette
- Requires pre-built CanvOS ISO/AMI via Edge Forge (github.com/spectrocloud/CanvOS)
- Provider images with desired K8s distro/version are built via CanvOS and hosted in a registry
- OS layer in cluster profile needs system:uri pointing to the provider image
- No cloud account needed for edge native clusters
- VIP required for cluster API endpoint

