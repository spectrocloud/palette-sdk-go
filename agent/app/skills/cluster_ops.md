# Cluster Operations

## inspect_cluster
Single-phase. Lists clusters or inspects one.

**List all clusters:**
```json
{}
```
Returns `data.clusters[]` — each has: uid, name, cloud_type, state.

**Inspect a specific cluster:**
```json
{"uid": "abc123"}
```
or:
```json
{"name": "my-cluster"}
```

**Inspect with specific sections only:**
```json
{"uid": "abc123", "include": ["summary", "variables"]}
```
Valid include values: `summary`, `variables`, `certs`, `backup`, `scan`. Omit for all.

---

## deploy_cluster (two-phase)

### Phase 1a — List available providers:
```json
{"phase": "plan"}
```
Returns `data.providers[]`: aws, azure, gcp, eks, aks, gke, vsphere, maas, cloudstack, edge_native.
Also returns field: `provider` (string, required, enum of providers).

### Phase 1b — Get required fields for a specific provider:
```json
{"phase": "plan", "provider": "eks"}
```
Returns `data.fields[]` with these common fields:
- `name` (string, required) — Cluster name
- `profile_name` (string, required) — Cluster profile name to apply
- `profile_version` (string, optional) — Profile version, latest if omitted
- `cloud_account_name` (string, required for most providers) — Cloud account name
- `cloud_account_scope` (string, optional, default "project", enum: tenant/project)
- Plus provider-specific fields in `data.fields[]`

Also returns:
- `data.available_cloud_accounts[]` — existing accounts for that provider
- `data.available_profiles[]` — existing profiles

### Phase 2 — Create the cluster:
```json
{
  "phase": "apply",
  "provider": "eks",
  "name": "prod-cluster",
  "profile_name": "eks-default",
  "cloud_account_name": "my-aws-account",
  "cloud_account_scope": "project"
}
```
Returns: `data.cluster_uid`, `data.cluster_name`, `data.provider`, `data.profile_uid`.

**Example — user says "deploy an EKS cluster":**
Call: `deploy_cluster` with `{"phase": "plan", "provider": "eks"}`

**Example — user says "what providers can I deploy to?":**
Call: `deploy_cluster` with `{"phase": "plan"}`

---

## teardown_cluster
Single-phase. Deletes a cluster.

**By UID:**
```json
{"uid": "abc123"}
```

**By name with force delete:**
```json
{"name": "my-cluster", "force": true}
```

Either `uid` or `name` required. `force` is optional (default false).
Returns: `data.cluster_uid`, `data.cluster_name`, `data.force`.

---

## update_cluster (two-phase)

### Phase 1a — List update modes:
```json
{"phase": "plan"}
```
Returns modes:
| Mode | Description |
|------|-------------|
| `profile` | Update or patch cluster profile values |
| `variables` | Update cluster profile variables |
| `metadata` | Update cluster metadata (name, labels, annotations) |
| `os_patch` | Configure OS patching schedule |

### Phase 1b — Get fields for a mode:
```json
{"phase": "plan", "mode": "profile"}
```

**Mode "profile" fields:**
- `uid` (string, required) — Cluster UID
- `profile_name` (string, required) — Profile name to apply
- `profile_version` (string, optional) — Profile version
- `patch` (boolean, optional, default false) — Use patch mode vs full update

**Mode "variables" fields:**
- `uid` (string, required) — Cluster UID
- `variables` (array, required) — Variable update entities with name and value

**Mode "metadata" fields:**
- `uid` (string, required) — Cluster UID
- `metadata` (object, required) — Object with name, labels, annotations

**Mode "os_patch" fields:**
- `uid` (string, required) — Cluster UID
- `os_patch` (object, required) — OS patch configuration entity

### Phase 2 — Apply:
```json
{"phase": "apply", "mode": "profile", "uid": "abc123", "profile_name": "new-profile"}
```

---

## LIMITATIONS
- No tool to list existing cloud accounts. Suggest the Palette UI.
- Cluster creation is asynchronous — returns a UID but provisioning takes time.
- Cluster names are resolved to UIDs internally for teardown/inspect.
