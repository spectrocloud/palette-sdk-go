# cluster_common_api

Provider-agnostic cluster operations: list, get, delete, kubeconfig, variables, profiles, backup, certs, metadata, OS patch

**get_cluster_by_name** — Get cluster details by name
`{"mode": "get_cluster_by_name", "name": "my-cluster"}`

**update_profile_values** — Update cluster profile values (full replacement)
`{"mode": "update_profile_values", "uid": "cluster-uid", "profiles": {"profiles": [{"uid": "profile-uid"}]}}`

**get_backup_config** — Get cluster backup configuration
`{"mode": "get_backup_config", "uid": "cluster-uid-here"}`

**approve_repave** — Approve a pending cluster repave operation
`{"mode": "approve_repave", "uid": "cluster-uid-here"}`

**get_rbac_config** — Get cluster RBAC configuration
`{"mode": "get_rbac_config", "uid": "cluster-uid-here"}`

**list_clusters** — List all clusters across all providers. Use client_filter to narrow results.
*Note: list_clusters returns summary data only (name, cloudType, health state, agent version). It does NOT include Kubernetes version. To get K8s version, use get_cluster for each cluster — the K8s version is in spec.clusterProfileTemplates[].packs[] where the pack layer is "k8s".*
`{"mode": "list_clusters"}`
*client_filter does case-insensitive substring matching on flattened fields:*
*All clusters:* `{"mode": "list_clusters"}`
*Edge clusters:* `{"mode": "list_clusters", "client_filter": {"cloudType": "edge-native"}}`
*EKS clusters:* `{"mode": "list_clusters", "client_filter": {"cloudType": "eks"}}`
*Unhealthy clusters:* `{"mode": "list_clusters", "client_filter": {"state": "UnHealthy"}}`
*Filter by name:* `{"mode": "list_clusters", "client_filter": {"name": "my-cluster"}}`
*Combine filters:* `{"mode": "list_clusters", "client_filter": {"cloudType": "edge-native", "state": "UnHealthy"}}`

**get_cluster** — Get cluster details by UID
`{"mode": "get_cluster", "uid": "cluster-uid-here"}`

**search_clusters** — Search clusters with filter and sort
`{"mode": "search_clusters", "filter": {"conjunction": "and", "filterGroups": []}}`

**get_admin_kubeconfig** — Get admin kubeconfig (elevated privileges)
`{"mode": "get_admin_kubeconfig", "uid": "cluster-uid-here"}`
*Admin kubeconfig has elevated privileges — use with caution*

**delete_cluster** — Delete a cluster by UID
`{"mode": "delete_cluster", "uid": "cluster-uid-here"}`

**get_variables** — Get cluster profile variables
`{"mode": "get_variables", "uid": "cluster-uid-here"}`

**update_variables** — Update cluster profile variable values
`{"mode": "update_variables", "uid": "cluster-uid", "variables": [{"name": "replica_count", "value": "3"}]}`

**create_backup_config** — Create backup configuration for a cluster
`{"mode": "create_backup_config", "uid": "cluster-uid", "config": {"backupLocationName": "s3-backup", "schedule": {"scheduledRunTime": "0 2 * * *"}}}`

**update_backup_config** — Update backup configuration for a cluster
`{"mode": "update_backup_config", "uid": "cluster-uid", "config": {"backupLocationName": "s3-backup", "schedule": {"scheduledRunTime": "0 3 * * *"}}}`

**get_certificates** — Get Kubernetes certificates for a cluster
`{"mode": "get_certificates", "uid": "cluster-uid-here"}`

**renew_certificates** — Initiate certificate renewal for a cluster
`{"mode": "renew_certificates", "uid": "cluster-uid-here"}`

**update_metadata** — Update cluster metadata (labels, annotations)
`{"mode": "update_metadata", "uid": "cluster-uid", "metadata": {"labels": {"env": "production"}, "annotations": {"team": "platform"}}}`

**get_cluster_summary** — Get cluster summary including health, status, and node counts
`{"mode": "get_cluster_summary", "uid": "cluster-uid-here"}`

**force_delete_cluster** — Force delete a cluster (use when normal delete fails or cluster is stuck)
`{"mode": "force_delete_cluster", "uid": "cluster-uid-here"}`
*Use only when normal delete fails or cluster is stuck in deleting state*

**get_client_kubeconfig** — Get client kubeconfig for kubectl access
`{"mode": "get_client_kubeconfig", "uid": "cluster-uid-here"}`

**patch_profile_values** — Patch cluster profile values (incremental update)
`{"mode": "patch_profile_values", "uid": "cluster-uid", "profiles": {"profiles": [{"uid": "profile-uid"}]}}`
*Use patch for incremental changes, update_profile_values for full replacement*

**update_os_patch** — Configure OS patching schedule
`{"mode": "update_os_patch", "uid": "cluster-uid", "config": {"schedule": "0 3 * * 0"}}`
*Schedule is a cron expression*
*Patches are applied during the maintenance window, one node at a time*

**get_namespace_config** — Get cluster namespace configuration
`{"mode": "get_namespace_config", "uid": "cluster-uid-here"}`

