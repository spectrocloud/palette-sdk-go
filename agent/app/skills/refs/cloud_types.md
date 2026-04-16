# Cloud Types & Accounts — Deep Reference

## Cloud Type Categories

| Category | Cloud Types | Account Required | Orchestration |
|----------|------------|-----------------|---------------|
| Public Cloud (IaaS) | aws, azure, gcp | Yes | Controller Mode |
| Managed Kubernetes | eks, aks, gke | Yes (same as IaaS parent) | Controller Mode |
| Data Center | vsphere, maas | Yes | Controller Mode |
| Edge | edge_native | No | Appliance or Agent Mode |

## Cloud Account Credential Formats

### AWS
```json
{
  "metadata": {"name": "my-aws-account"},
  "spec": {
    "accessKeyId": "AKIA...",
    "secretAccessKey": "...",
    "arn": "arn:aws:iam::role/..."  // optional, for STS assume-role
  }
}
```

### Azure
```json
{
  "metadata": {"name": "my-azure-account"},
  "spec": {
    "tenantId": "...",
    "clientId": "...",
    "clientSecret": "...",
    "subscriptionId": "..."
  }
}
```

### GCP
```json
{
  "metadata": {"name": "my-gcp-account"},
  "spec": {
    "jsonCredentials": "{...service account key JSON...}"
  }
}
```

### vSphere
```json
{
  "metadata": {"name": "my-vsphere-account"},
  "spec": {
    "vcenterServer": "vcenter.example.com",
    "username": "admin@vsphere.local",
    "password": "..."
  }
}
```

### MaaS
```json
{
  "metadata": {"name": "my-maas-account"},
  "spec": {
    "apiEndpoint": "http://maas.example.com:5240/MAAS",
    "apiKey": "..."
  }
}
```

### CloudStack
```json
{
  "metadata": {"name": "my-cloudstack-account"},
  "spec": {
    "apiUrl": "http://cloudstack.example.com:8080/client/api",
    "apiKey": "...",
    "secretKey": "..."
  }
}
```

## Scope Rules
- Cloud Accounts can be **tenant-scoped** (visible to all projects) or **project-scoped** (visible only within one project)
- When deploying a cluster, use `cloud_account_scope: "tenant"` or `"project"` (default)
- EKS uses the same AWS account, AKS uses the same Azure account, GKE uses the same GCP account
