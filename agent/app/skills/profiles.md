# Cluster Profiles

## manage_profile (two-phase)

### Phase 1a — List available modes:
```json
{"phase": "plan"}
```
Returns modes:
| Mode | Description |
|------|-------------|
| `create` | Create a new cluster profile and optionally publish it |
| `import` | Import a cluster profile from content string |
| `update` | Update an existing cluster profile |
| `delete` | Delete a cluster profile by UID |
| `list` | List all available cluster profiles |

### Phase 1b — Get fields for a mode:
```json
{"phase": "plan", "mode": "create"}
```

**Mode "create" fields:**
- `profile` (object, required) — V1ClusterProfileEntity with metadata and spec
- `publish` (boolean, optional, default false) — Publish profile after creation

**Mode "import" fields:**
- `import_content` (string, required) — Profile content to import

**Mode "update" fields:**
- `profile_update` (object, required) — V1ClusterProfileUpdateEntity

**Mode "delete" fields:**
- `uid` (string, required) — Profile UID to delete

**Mode "list" fields:** None required.

### Phase 2 — Apply:

**List all profiles:**
```json
{"phase": "apply", "mode": "list"}
```
Returns: `data.profiles[]` — each has uid, name.

**Create a profile:**
```json
{"phase": "apply", "mode": "create", "profile": {"metadata": {"name": "my-profile"}, "spec": {...}}, "publish": true}
```
Returns: `data.profile_uid`, `data.published`.

**Delete a profile:**
```json
{"phase": "apply", "mode": "delete", "uid": "profile-uid-here"}
```

**Import a profile:**
```json
{"phase": "apply", "mode": "import", "import_content": "...exported string..."}
```
Returns: `data.profile_uid`.

**Example — user says "list profiles" or "show my profiles":**
Call: `manage_profile` with `{"phase": "apply", "mode": "list"}`

**Example — user says "delete profile X":**
Call: `manage_profile` with `{"phase": "apply", "mode": "delete", "uid": "X"}`

## LIMITATIONS
- Profile update requires the full V1ClusterProfileUpdateEntity object.
- Import requires a previously exported profile content string.
- No standalone "publish" mode — use `publish: true` with create, or update the profile.
- No "clone" mode — to clone, export then import.
