# maas_cluster_api

Create MAAS bare-metal clusters

**create_maas** — Create a new MAAS bare-metal cluster
`{"mode": "create_maas", "cluster": {"metadata": {"name": "prod-maas"}, "spec": {"cloudAccountUid": "...", "profiles": [{"uid": "..."}]}}}`
*Requires a valid MAAS cloud account (api_endpoint, api_key)*
*Uses Controller Mode*
*Can serve as foundation for VMO (KubeVirt) workloads*

