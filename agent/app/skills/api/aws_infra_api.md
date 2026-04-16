# aws_infra_api

Query and manage AWS infrastructure directly. Credentials from env vars (AWS_ACCESS_KEY_ID, AWS_SECRET_ACCESS_KEY).

## Network

**list_vpcs** — List VPCs
`{"mode": "list_vpcs", "region": "us-east-1"}`

**list_subnets** — List subnets in a VPC
`{"mode": "list_subnets", "vpc_id": "vpc-abc123"}`

**list_key_pairs** — List SSH key pairs
`{"mode": "list_key_pairs"}`

**list_security_groups** — List security groups
`{"mode": "list_security_groups", "vpc_id": "vpc-abc123"}`

**list_availability_zones** — List AZs
`{"mode": "list_availability_zones", "region": "us-east-1"}`

**list_instance_types** — List common EC2 instance types
`{"mode": "list_instance_types"}`

## EKS

**list_eks_clusters** — List native EKS clusters in AWS
`{"mode": "list_eks_clusters", "region": "us-east-1"}`

## EC2 Instances

**launch_instance** — Launch a new EC2 instance
`{"mode": "launch_instance", "spec": {"name": "my-vm", "imageId": "ami-0c02fb55956c7d316", "instanceType": "t3.medium", "keyName": "my-key", "subnetId": "subnet-abc123", "securityGroups": ["sg-abc123"], "diskSizeGb": 30}}`
*Required: imageId (AMI ID). Defaults: instanceType=t3.medium. Optional: keyName, subnetId, securityGroups, diskSizeGb, name.*

**list_instances** — List EC2 instances (optionally filtered)
`{"mode": "list_instances"}`
`{"mode": "list_instances", "filters": {"instance-state-name": "running"}}`
`{"mode": "list_instances", "filters": {"vpc-id": "vpc-abc123"}}`
*Common filters: instance-state-name (running/stopped/terminated), vpc-id, tag:Name*

**get_instance** — Get details of a specific instance
`{"mode": "get_instance", "instance_id": "i-abc123"}`

**start_instance** — Start a stopped instance
`{"mode": "start_instance", "instance_id": "i-abc123"}`

**stop_instance** — Stop a running instance
`{"mode": "stop_instance", "instance_id": "i-abc123"}`

**reboot_instance** — Reboot an instance
`{"mode": "reboot_instance", "instance_id": "i-abc123"}`

**terminate_instance** — Terminate (delete) an instance
`{"mode": "terminate_instance", "instance_id": "i-abc123"}`
*WARNING: This permanently destroys the instance*

*`region` defaults to us-east-1 (from AWS_REGION env var) if omitted.*
