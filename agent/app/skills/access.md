# Access & Identity

## manage_users
Single-phase. No plan/apply. Defaults to "list" if no mode.

**Modes and required fields:**

| Mode | Required fields | Description |
|------|----------------|-------------|
| `list` | *none* | List all users |
| `get_by_id` | `uid` | Find user by UID |
| `get_by_email` | `email` | Find user by email |
| `get_by_name` | `name` | Find user by name |
| `delete` | `uid` | Delete a user |

**List all users:**
```json
{"mode": "list"}
```
Returns: `data.users[]` — each has uid, name, email.

**Find user by email:**
```json
{"mode": "get_by_email", "email": "jane@example.com"}
```
Returns: `data.user` — full user object.

**Find user by name:**
```json
{"mode": "get_by_name", "name": "Jane Doe"}
```

**Delete a user:**
```json
{"mode": "delete", "uid": "user-uid-here"}
```

**Assign tenant role (by email):**
```json
{"mode": "assign_tenant_role", "email": "jane@example.com", "tenant_roles": {"roleUids": ["role-uid-1"]}}
```

**Assign tenant role (by UID):**
```json
{"mode": "assign_tenant_role", "uid": "user-uid", "tenant_roles": {"roleUids": ["role-uid-1"]}}
```

**Assign project role:**
```json
{"mode": "assign_project_role", "email": "jane@example.com", "project_roles": {"projects": [{"projectUid": "proj-uid", "roles": [{"uid": "role-uid"}]}]}}
```

User can be identified by `uid`, `email`, or `name` for role assignment — resolved automatically.

**Example — user says "list users" or "show all users":**
Call: `manage_users` with `{"mode": "list"}`

**Example — user says "find user john@acme.com":**
Call: `manage_users` with `{"mode": "get_by_email", "email": "john@acme.com"}`

**Example — user says "make jane@acme.com a tenant admin":**
Call: `manage_users` with `{"mode": "assign_tenant_role", "email": "jane@acme.com", "tenant_roles": {"roleUids": ["<tenant-admin-role-uid>"]}}`
Note: Role UID is required. If the agent doesn't know it, ask the user.

---

## onboard_user
Single-phase. Creates a user and assigns roles atomically.

```json
{
  "user": {
    "metadata": {"name": "jane-doe"},
    "spec": {"firstName": "Jane", "lastName": "Doe", "emailId": "jane@example.com"}
  },
  "tenant_roles": {"roleUids": ["role-uid-1"]},
  "project_roles": {"projects": [{"projectUid": "proj-uid", "roles": [{"uid": "role-uid"}]}]}
}
```
`user` is required. `tenant_roles` and `project_roles` are optional.
Returns: `data.user_uid`, plus `data.tenant_roles` and `data.project_roles` status.

**Example — user says "create user jane@example.com":**
Call: `onboard_user` with `{"user": {"spec": {"firstName": "Jane", "lastName": "Doe", "emailId": "jane@example.com"}}}`

---

## manage_team (two-phase)

### Phase 1a — List modes:
```json
{"phase": "plan"}
```
Returns modes:
| Mode | Description |
|------|-------------|
| `create` | Create a new team |
| `get` | Get team details by UID or name |
| `update` | Update an existing team |
| `delete` | Delete a team |
| `assign_project_role` | Assign project roles to a team |
| `assign_tenant_role` | Assign tenant roles to a team |

### Phase 1b — Get fields for a mode:

**Mode "create":** `team` (object, required) — V1Team with metadata.name and spec.users
**Mode "get":** `uid` or `name` (either one)
**Mode "update":** `uid` (required) + `team` (required)
**Mode "delete":** `uid` or `name` (either one)
**Mode "assign_project_role":** `uid` (required) + `project_roles` (required)
**Mode "assign_tenant_role":** `uid` (required) + `tenant_roles` (required)

### Phase 2 — Apply:

**Get a team:**
```json
{"phase": "apply", "mode": "get", "name": "my-team"}
```

**Create a team:**
```json
{"phase": "apply", "mode": "create", "team": {"metadata": {"name": "dev-team"}, "spec": {"users": ["user-uid-1"]}}}
```

**Delete a team:**
```json
{"phase": "apply", "mode": "delete", "name": "my-team"}
```

**Example — user says "get team dev-ops":**
Call: `manage_team` with `{"phase": "apply", "mode": "get", "name": "dev-ops"}`

## LIMITATIONS
- No tool to list ALL teams. Can only get a specific team by uid or name.
- To create users, use `onboard_user`, NOT `manage_users`.
- `manage_users` handles: list, lookup, delete, and role assignment — but NOT user creation.
- Role assignment requires knowing the role UID. If unknown, ask the user for it.
