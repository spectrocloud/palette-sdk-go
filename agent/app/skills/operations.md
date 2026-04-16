# Operations & Registries

## configure_backup
Single-phase. Auto-detects create vs update.

**Check current backup status:**
```json
{"cluster_uid": "cluster-uid-here"}
```
Returns: `data.has_backup` (boolean). If true, `data.backup` has the current config.

**Create or update backup config:**
```json
{"cluster_uid": "cluster-uid-here", "config": {"backupLocationName": "s3-backup", "schedule": {"scheduledRunTime": "0 2 * * *"}, "namespaces": ["default"]}}
```
Returns: `data.action` — "created" or "updated".

`cluster_uid` is always required. If `config` is omitted, just returns current status.

**Example — user says "check backup for cluster abc123":**
Call: `configure_backup` with `{"cluster_uid": "abc123"}`

---

## run_compliance_scan
Single-phase. Auto-detects create vs update.

**Check current scan status:**
```json
{"cluster_uid": "cluster-uid-here"}
```
Returns: `data.has_scan` (boolean). If true, `data.scan` has the current config.

**Create or update scan config:**
```json
{"cluster_uid": "cluster-uid-here", "config": {"schedule": {"scheduledRunTime": "0 3 * * 0"}, "driverSpec": {"kubeBench": {"enabled": true}, "kubeHunter": {"enabled": true}}}}
```
Returns: `data.action` — "created" or "updated".

**Example — user says "is compliance scanning enabled for cluster abc123?":**
Call: `run_compliance_scan` with `{"cluster_uid": "abc123"}`

---

## manage_vm (two-phase)

### Phase 1a — List modes:
```json
{"phase": "plan"}
```
Returns modes:
| Mode | Description |
|------|-------------|
| `create` | Create a new virtual machine |
| `get` | Get virtual machine details |
| `delete` | Delete a virtual machine |
| `start` | Start a stopped VM |
| `stop` | Stop a running VM |
| `pause` | Pause a running VM |
| `resume` | Resume a paused VM |
| `restart` | Restart a VM |
| `migrate` | Migrate VM to another node |
| `clone` | Clone an existing VM |

### Phase 1b — Get fields for a mode:
All modes require `cluster_uid` (string, required).

**Mode "create":** `vm` (object, required) — V1ClusterVirtualMachine spec
**Modes "get/delete/start/stop/pause/resume/restart/migrate":** `vm_name` + `namespace` (both required)
**Mode "clone":** `vm_name` + `namespace` + `clone_vm_name` (all required)

### Phase 2 — Apply:

**Get a VM:**
```json
{"phase": "apply", "mode": "get", "cluster_uid": "abc123", "vm_name": "my-vm", "namespace": "default"}
```

**Start a VM:**
```json
{"phase": "apply", "mode": "start", "cluster_uid": "abc123", "vm_name": "my-vm", "namespace": "default"}
```

**Clone a VM:**
```json
{"phase": "apply", "mode": "clone", "cluster_uid": "abc123", "vm_name": "source-vm", "namespace": "default", "clone_vm_name": "cloned-vm"}
```

**Example — user says "start VM web-server on cluster abc123":**
Call: `manage_vm` with `{"phase": "apply", "mode": "start", "cluster_uid": "abc123", "vm_name": "web-server", "namespace": "default"}`

---

## setup_registry (two-phase)

### Phase 1a — List registry types:
```json
{"phase": "plan"}
```
Returns field: `registry_type` (string, required, enum: helm, oci_basic, ecr).

### Phase 1b — Get fields for a type:
```json
{"phase": "plan", "registry_type": "helm"}
```

| Type | Spec description |
|------|-----------------|
| `helm` | V1HelmRegistryEntity with metadata.name, spec.endpoint, and auth config |
| `oci_basic` | V1BasicOciRegistry with metadata.name, spec.endpoint, and auth config |
| `ecr` | V1EcrRegistry with metadata.name, spec containing AWS credentials and region |

### Phase 2 — Create:
```json
{"phase": "apply", "registry_type": "helm", "spec": {"metadata": {"name": "my-helm"}, "spec": {"endpoint": "https://charts.example.com", "auth": {"username": "...", "password": "..."}}}}
```
Returns: `data.registry_uid`, `data.registry_type`.

**Example — user says "what types of registries can I create?":**
Call: `setup_registry` with `{"phase": "plan"}`

## LIMITATIONS
- No tool to list existing registries. Suggest the Palette UI.
- No tool to update or delete registries. This tool only creates.
- VM operations require knowing cluster_uid, vm_name, and namespace — use inspect_cluster first to find the cluster UID.
- Backup and scan tools require cluster_uid — use inspect_cluster to get it.
