# Deploying Edge Cluster Agent

Deploy an edge cluster using agent mode — launches AWS VMs with stylus agent, registers hosts, creates cluster

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

### Preparing AWS Edge Host Agent
Launch AWS EC2 instances as edge hosts with stylus agent installed via user-data

- **Select VPC for edge hosts**
- **Select subnet**
  Use list_subnets with the selected VPC ID: {"mode": "list_subnets", "vpc_id": "<vpc_id>", "interactive": true}
- **Select SSH key pair**
- **Launch EC2 instances with agent user-data**
  For each host, call launch_instance with the Ubuntu AMI, selected subnet, key pair, and the agent install user-data script as user-data. The user-data installs the stylus agent which auto-registers the host with Palette.
  *Default AMI: Ubuntu 22.04 LTS*
  *Default instance type: t3.xlarge*
  *The stylus agent is installed via user-data bash script on first boot*
  *Hosts auto-register with Palette using the edge registration token*
  *Tags on the EC2 instance help identify hosts later*

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

Agent mode edge cluster on AWS

### Steps

1. **Create edge registration token**
   *Uses shared sub-journey: ensure_edge_registration_token*
   Get or create a registration token for this deployment

2. **Ensure edge cluster profile exists**
   *Uses shared sub-journey: ensure_edge_cluster_profile*
   Context: Cloud type is edge-native, agent mode. OS layer is standard (no provider image URI). K8s should be K3S or PXK-E.

3. **Collect deployment configuration**
   Collect: number of hosts (min 1), how many CP nodes, how many workers, VIP IP address, AWS region, site name/tags

4. **Prepare AWS edge hosts with agent**
   *Uses shared sub-journey: prepare_aws_edge_host_agent*
   Context: Launch N EC2 instances with Ubuntu AMI and stylus agent user-data. Use the registration token from step 1.

5. **Wait for hosts to register**
   *Uses shared sub-journey: register_edge_hosts*
   Context: Poll until all N hosts appear as registered in Palette

6. **Create edge cluster**
   Tool: `edge_cluster_api`
   Create the edge native cluster with: profile UID, VIP, edge host UIDs assigned to CP and worker pools based on user's counts.

### After Completion

- Edge cluster provisioning takes 5-10 minutes after hosts are assigned
- Monitor status with cluster_common_api
- For single-node clusters, the same host serves as both CP and worker

---

## Notes

- Agent mode: end users manage the OS lifecycle
- The stylus agent is installed on first boot via user-data bash script
- Hosts auto-register with Palette using the edge registration token
- No cloud account is needed for edge native clusters
- A VIP (Virtual IP) is required for the cluster API endpoint

