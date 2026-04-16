# vmware_cluster_api

Create VMware vSphere clusters

**create_vsphere** — Create a new vSphere cluster
`{"mode": "create_vsphere", "cluster": {"metadata": {"name": "prod-vsphere"}, "spec": {"cloudAccountUid": "...", "profiles": [{"uid": "..."}], "cloudConfig": {"datacenter": "DC1", "folder": "/DC1/vm", "network": {"networkName": "VM Network"}}}}}`
*Requires a valid vSphere cloud account (vcenter server, username, password)*
*Uses Controller Mode*
*Can serve as foundation for VMO (KubeVirt) workloads*

