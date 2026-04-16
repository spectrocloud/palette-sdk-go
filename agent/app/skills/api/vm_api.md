# vm_api

Create, manage, and control virtual machines on clusters

**delete_vm** — Delete a virtual machine
`{"mode": "delete_vm", "cluster_uid": "cluster-uid", "namespace": "default", "name": "my-vm"}`

**start_vm** — Start a stopped VM
`{"mode": "start_vm", "cluster_uid": "cluster-uid", "vm_name": "my-vm", "vm_namespace": "default"}`

**stop_vm** — Stop a running VM
`{"mode": "stop_vm", "cluster_uid": "cluster-uid", "vm_name": "my-vm", "vm_namespace": "default"}`

**resume_vm** — Resume a paused VM
`{"mode": "resume_vm", "cluster_uid": "cluster-uid", "vm_name": "my-vm", "vm_namespace": "default"}`

**restart_vm** — Restart a VM
`{"mode": "restart_vm", "cluster_uid": "cluster-uid", "vm_name": "my-vm", "vm_namespace": "default"}`

**list_vms** — List all VMs on a cluster
`{"mode": "list_vms", "cluster": {"metadata": {"uid": "cluster-uid"}}}`
*Requires the full cluster object — use eks_cluster_api get_cluster first*

**get_vm** — Get VM details
`{"mode": "get_vm", "cluster_uid": "cluster-uid", "namespace": "default", "name": "my-vm"}`

**create_vm** — Create a new virtual machine
`{"mode": "create_vm", "cluster_uid": "cluster-uid", "vm": {"metadata": {"name": "my-vm", "namespace": "default"}, "spec": {"running": true, "template": {"spec": {"domain": {"cpu": {"cores": 2}, "memory": {"guest": "4Gi"}}}}}}}`

**pause_vm** — Pause a running VM
`{"mode": "pause_vm", "cluster_uid": "cluster-uid", "vm_name": "my-vm", "vm_namespace": "default"}`

**migrate_vm** — Live-migrate a VM to another node
`{"mode": "migrate_vm", "cluster_uid": "cluster-uid", "vm_name": "my-vm", "vm_namespace": "default"}`

**clone_vm** — Clone an existing VM
`{"mode": "clone_vm", "cluster_uid": "cluster-uid", "source_vm_name": "my-vm", "clone_vm_name": "my-vm-clone", "namespace": "default"}`

