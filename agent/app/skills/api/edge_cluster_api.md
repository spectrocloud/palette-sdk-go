# edge_cluster_api

Create Edge Native clusters

**create_edge_native** — Create a new Edge Native cluster
`{"mode": "create_edge_native", "cluster": {"metadata": {"name": "edge-site-01"}, "spec": {"profiles": [{"uid": "..."}], "cloudConfig": {"edgeHosts": [{"hostUid": "..."}]}}}}`
*No cloud account required — uses pre-registered Edge Hosts*
*Uses Appliance Mode (Kairos immutable OS) or Agent Mode*
*Edge Hosts must be registered with Palette before cluster creation*
*Supports K3S and PXK-E Kubernetes distributions*

