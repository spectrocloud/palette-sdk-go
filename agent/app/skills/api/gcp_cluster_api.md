# gcp_cluster_api

Create GCP IaaS clusters

**create_gcp** — Create a new GCP IaaS cluster
`{"mode": "create_gcp", "cluster": {"metadata": {"name": "prod-gcp"}, "spec": {"cloudAccountUid": "...", "profiles": [{"uid": "..."}], "cloudConfig": {"project": "my-gcp-project", "region": "us-central1"}}}}`
*Requires a valid GCP cloud account UID (service account JSON credentials)*
*Uses Controller Mode*

