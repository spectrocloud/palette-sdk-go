# gke_cluster_api

Create GKE (Google managed Kubernetes) clusters

**create_gke** — Create a new GKE cluster
`{"mode": "create_gke", "cluster": {"metadata": {"name": "prod-gke"}, "spec": {"cloudAccountUid": "...", "profiles": [{"uid": "..."}], "cloudConfig": {"project": "my-gcp-project", "region": "us-central1"}}}}`
*GKE uses GCP cloud account credentials*
*Google manages the control plane, Palette manages worker node pools*

