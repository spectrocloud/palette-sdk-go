# Virtual Machine Orchestrator (VMO) — Deep Reference

## Overview

VMO is Palette's VMware alternative, built on KubeVirt. It allows running and managing virtual machines alongside containers on the same Kubernetes cluster.

## Architecture

- **Foundation:** KubeVirt runs as an add-on on Kubernetes clusters
- **Supported base clusters:** MaaS-based (bare metal) or Edge clusters
- **VM lifecycle:** Fully managed through Kubernetes CRDs (VirtualMachine, VirtualMachineInstance)
- **Storage:** Uses Kubernetes PVCs and CSI drivers for VM disks
- **Networking:** Integrates with Kubernetes CNI, supports multus for multiple networks

## VMware Migration (Forklift)

- Migration from VMware vSphere to KubeVirt is supported using the Forklift project
- Forklift converts VMware VMDK disks to KubeVirt-compatible formats
- Supports warm migration (pre-copy) and cold migration (offline)
- Migration workflow:
  1. Connect Forklift to source vSphere environment
  2. Select VMs to migrate
  3. Map source networks/storage to target Kubernetes resources
  4. Execute migration plan
  5. Validate migrated VMs on KubeVirt

## VM Lifecycle Operations

| Operation | Description |
|-----------|-------------|
| create | Create a new VM from a spec (CPU, memory, disk, image) |
| get | Retrieve VM details and status |
| delete | Remove a VM and its resources |
| start | Power on a stopped VM |
| stop | Gracefully shut down a running VM |
| pause | Freeze a running VM (preserves state in memory) |
| resume | Unfreeze a paused VM |
| restart | Stop then start a VM |
| migrate | Live-migrate a VM to a different node |
| clone | Create a copy of an existing VM |

## Reference Architectures

Spectro Cloud publishes reference architectures in the Profiles Hub for common VMO scenarios:
- Development workstations
- Legacy application hosting
- Database servers
- Windows workloads on Kubernetes

## Requirements

- Cluster must have KubeVirt add-on profile applied
- Nodes must support hardware virtualization (nested virt or bare-metal)
- Sufficient CPU and memory resources for VMs
- Storage class that supports ReadWriteOnce (RWO) or ReadWriteMany (RWX) PVCs
