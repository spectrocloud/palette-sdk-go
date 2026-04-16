# azure_cluster_api

Create Azure IaaS clusters

**create_azure** — Create a new Azure IaaS cluster
`{"mode": "create_azure", "cluster": {"metadata": {"name": "prod-azure"}, "spec": {"cloudAccountUid": "...", "profiles": [{"uid": "..."}], "cloudConfig": {"region": "eastus", "subscriptionId": "..."}}}}`
*Requires a valid Azure cloud account UID*
*Requires an Azure-compatible infra cluster profile UID*
*Uses Controller Mode (CAPI-based orchestration)*

