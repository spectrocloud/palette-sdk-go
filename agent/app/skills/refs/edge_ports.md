# Edge Cluster Required Ports (AWS Security Group)

## Outbound

| Port | Protocol | Category | Purpose |
|------|----------|----------|---------|
| 443 | TCP | Palette Management | HTTPS + gRPC to Palette management plane, registry, and repo; also used for Remote Shell (SSH) |

## Inbound

| Port | Protocol | Category | Purpose |
|------|----------|----------|---------|
| 22 | TCP | SSH | SSH access for management |
| 5080 | TCP | Agent / Local UI | Palette agent / Local UI access |
| 5082 | TCP | Agent / Local UI | Content bundle upload (airgap/local management) |
| 6443 | TCP | Kubernetes | Kubernetes API server |

## Inter-node (within the cluster)

| Port | Protocol | Category | Purpose |
|------|----------|----------|---------|
| 6443 | TCP | Kubernetes | Kubernetes API server |
| 2379-2380 | TCP | Kubernetes (Control Plane) | etcd server client API |
| 10250 | TCP | Kubernetes | Kubelet API (control plane & worker nodes) |
| 10259 | TCP | Kubernetes (Control Plane) | kube-scheduler |
| 10257 | TCP | Kubernetes (Control Plane) | kube-controller-manager |
| 30000-32767 | TCP | Kubernetes (Worker) | NodePort Services |
| 5473 | TCP | CNI - Calico | Calico Typha (Helm-based installation) |
| 9098 | TCP | CNI - Calico | Calico Typha health check |
| 8472 | UDP | CNI - Cilium | VXLAN overlay networking |
| 4240 | TCP | CNI - Cilium | Cilium health checks |
