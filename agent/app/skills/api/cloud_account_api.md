# cloud_account_api

List, create, update, and delete cloud accounts for all providers

**delete_vsphere** — Delete a vSphere cloud account
`{"mode": "delete_vsphere", "uid": "account-uid-here"}`

**create_maas** — Create a MAAS cloud account
`{"mode": "create_maas", "account": {"metadata": {"name": "my-maas"}, "spec": {"apiEndpoint": "http://maas:5240/MAAS", "apiKey": "..."}}}`

**get_account** — Get cloud account summary by UID
`{"mode": "get_account", "uid": "account-uid-here"}`

**delete_aws** — Delete an AWS cloud account
`{"mode": "delete_aws", "uid": "account-uid-here"}`

**create_azure** — Create an Azure cloud account
`{"mode": "create_azure", "account": {"metadata": {"name": "my-azure"}, "spec": {"tenantId": "...", "clientId": "...", "clientSecret": "...", "subscriptionId": "..."}}}`

**list_accounts** — List all cloud accounts
`{"mode": "list_accounts"}`

**get_gcp_by_name** — Get GCP account by name
`{"mode": "get_gcp_by_name", "name": "my-gcp-account"}`

**create_vsphere** — Create a vSphere cloud account
`{"mode": "create_vsphere", "account": {"metadata": {"name": "my-vsphere"}, "spec": {"vcenterServer": "vcenter.example.com", "username": "admin", "password": "..."}}}`

**delete_maas** — Delete a MAAS cloud account
`{"mode": "delete_maas", "uid": "account-uid-here"}`

**create_aws** — Create an AWS cloud account
`{"mode": "create_aws", "account": {"metadata": {"name": "my-aws"}, "spec": {"accessKeyId": "AKIA...", "secretAccessKey": "..."}}}`

**get_azure_by_name** — Get Azure account by name
`{"mode": "get_azure_by_name", "name": "my-azure-account"}`

**delete_azure** — Delete an Azure cloud account
`{"mode": "delete_azure", "uid": "account-uid-here"}`

**delete_gcp** — Delete a GCP cloud account
`{"mode": "delete_gcp", "uid": "account-uid-here"}`

**get_aws_by_name** — Get AWS account by name
`{"mode": "get_aws_by_name", "name": "my-aws-account", "scope": "project"}`

**create_gcp** — Create a GCP cloud account
`{"mode": "create_gcp", "account": {"metadata": {"name": "my-gcp"}, "spec": {"jsonCredentials": "{...service account key...}"}}}`

