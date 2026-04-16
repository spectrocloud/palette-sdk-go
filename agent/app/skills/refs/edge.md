# Edge Deployments — Deep Reference

## Overview

Palette supports deploying Kubernetes clusters to remote edge locations using physical or virtual bare-metal machines. Edge locations are typically remote sites like retail stores, factory floors, cell towers, or branch offices.

## Edge Hosts

- Any physical or virtual bare-metal machine at an edge location
- Registered with Palette as an "Edge Host"
- Can be managed via Appliance Mode or Agent Mode
- Edge Hosts are associated with an edge site/location

## Appliance Mode

- Uses **Kairos** — an immutable OS toolkit by Spectro Cloud
- Pre-built appliance images are created with:
  - Immutable OS (based on Ubuntu or openSUSE)
  - Kubernetes distribution (PXK-E or K3S)
  - Palette agent baked in
- Installation process:
  1. Build an edge installer ISO using Palette's Edge Forge tooling
  2. Flash/install the ISO on target hardware
  3. Machine boots and auto-registers with Palette
  4. Palette provisions the cluster on registered Edge Hosts
- **Benefits:** Secure (immutable, tamper-resistant), predictable, air-gap friendly
- **Updates:** Delivered as new OS image layers — atomic, rollback-capable

## Agent Mode

- Palette agent installed on **customer-managed host OS**
- Supports existing Linux distributions (Ubuntu, RHEL, etc.)
- Installation process:
  1. Install the Palette agent on target machine
  2. Agent contacts Palette and registers the Edge Host
  3. Palette provisions the cluster on registered hosts
- **Benefits:** Flexible, works with existing OS and infrastructure
- **Trade-offs:** OS management is the customer's responsibility

## Edge-Specific Considerations

- **Air-gapped deployments:** Appliance Mode supports fully disconnected operation. Content is pre-loaded into the appliance image.
- **Low bandwidth:** Edge clusters can operate with intermittent connectivity to Palette
- **Scale:** Palette supports thousands of edge sites from a single management plane
- **K8s distributions:** K3S (lightweight, preferred for edge) or PXK-E (full-featured edge variant)
- **Single-node clusters:** Supported for resource-constrained edge sites
- **Multi-node clusters:** Supported for higher availability at edge
