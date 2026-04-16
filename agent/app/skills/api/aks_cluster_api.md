# aks_cluster_api

Create AKS (Azure managed Kubernetes) clusters

**create_aks** — Create a new AKS cluster
`{"mode": "create_aks", "cluster": {"metadata": {"name": "prod-aks"}, "spec": {"cloudAccountUid": "...", "profiles": [{"uid": "..."}], "cloudConfig": {"region": "eastus"}}}}`
*AKS uses Azure cloud account credentials*
*Uses Controller Mode — Azure manages the control plane, Palette manages worker nodes*

