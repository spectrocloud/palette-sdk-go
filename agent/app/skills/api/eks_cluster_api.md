# eks_cluster_api

Create EKS clusters

**create_eks** — Create a new EKS cluster
`{"mode": "create_eks", "cluster": {"metadata": {"name": "prod-eks"}, "spec": {"cloudAccountUid": "...", "profiles": [{"uid": "..."}], "cloudConfig": {"region": "us-east-1"}}}}`
*Requires a valid AWS cloud account UID*
*Requires a valid EKS-compatible cluster profile UID*
*Region must be a valid AWS region*
*This is the atomic API — for guided deployment, see the deploy_eks_cluster journey*

