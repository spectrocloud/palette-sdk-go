# registry_api

List, create, update, and delete container registries (Helm, OCI, ECR)

**list_helm_registries** — List all Helm registries
`{"mode": "list_helm_registries", "scope": "project"}`

**delete_ecr_registry** — Delete an ECR registry
`{"mode": "delete_ecr_registry", "uid": "registry-uid-here"}`

**list_oci_registries** — List all OCI registries
`{"mode": "list_oci_registries"}`

**get_helm_registry** — Get Helm registry by UID
`{"mode": "get_helm_registry", "uid": "registry-uid-here"}`

**get_helm_registry_by_name** — Get Helm registry by name
`{"mode": "get_helm_registry_by_name", "name": "my-helm-registry"}`

**create_helm_registry** — Create a Helm registry
`{"mode": "create_helm_registry", "registry": {"metadata": {"name": "my-helm"}, "spec": {"endpoint": "https://charts.example.com", "auth": {"username": "...", "password": "..."}}}}`

**delete_helm_registry** — Delete a Helm registry
`{"mode": "delete_helm_registry", "uid": "registry-uid-here"}`

**create_oci_basic_registry** — Create an OCI basic registry
`{"mode": "create_oci_basic_registry", "registry": {"metadata": {"name": "my-oci"}, "spec": {"endpoint": "https://registry.example.com"}}}`

**delete_oci_basic_registry** — Delete an OCI basic registry
`{"mode": "delete_oci_basic_registry", "uid": "registry-uid-here"}`

**create_ecr_registry** — Create an ECR registry
`{"mode": "create_ecr_registry", "registry": {"metadata": {"name": "my-ecr"}, "spec": {"accessKey": "...", "secretKey": "...", "region": "us-east-1"}}}`

**list_pack_registries** — List all pack registries
`{"mode": "list_pack_registries"}`

