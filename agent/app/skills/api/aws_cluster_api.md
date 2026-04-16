# aws_cluster_api

Create AWS IaaS clusters

**create_aws** — Create a new AWS IaaS cluster
`{"mode": "create_aws", "cluster": {"metadata": {"name": "prod-aws"}, "spec": {"cloudAccountUid": "...", "profiles": [{"uid": "..."}], "cloudConfig": {"region": "us-east-1", "sshKeyName": "my-key"}}}}`
*Requires a valid AWS cloud account UID*
*Requires an AWS-compatible infra cluster profile UID*
*Uses Controller Mode (CAPI-based orchestration)*

