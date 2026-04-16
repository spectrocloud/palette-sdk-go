# Platform & Projects

## manage_project (two-phase)

### Phase 1a — List modes:
```json
{"phase": "plan"}
```
Returns modes:
| Mode | Description |
|------|-------------|
| `create` | Create a new project |
| `get` | Get project details by UID or name |
| `list` | List all projects |
| `update` | Update an existing project |
| `delete` | Delete a project |

### Phase 1b — Get fields for a mode:

**Mode "create":** `project` (object, required) — V1ProjectEntity with metadata.name and spec
**Mode "get":** `uid` or `name` (either one, name is resolved to UID)
**Mode "list":** No fields required.
**Mode "update":** `uid` (required) + `project` (required)
**Mode "delete":** `uid` or `name` (either one)

### Phase 2 — Apply:

**List all projects:**
```json
{"phase": "apply", "mode": "list"}
```
Returns: `data.projects[]` — each has uid, name.

**Get a project by name:**
```json
{"phase": "apply", "mode": "get", "name": "production"}
```
Returns: `data.project` — full project object.

**Create a project:**
```json
{"phase": "apply", "mode": "create", "project": {"metadata": {"name": "staging"}, "spec": {}}}
```
Returns: `data.project_uid`.

**Delete a project:**
```json
{"phase": "apply", "mode": "delete", "name": "old-project"}
```

**Example — user says "list projects" or "how many projects do I have?":**
Call: `manage_project` with `{"phase": "apply", "mode": "list"}`

**Example — user says "show project production":**
Call: `manage_project` with `{"phase": "apply", "mode": "get", "name": "production"}`

---

## configure_platform (two-phase)

### Phase 1a — List modes:
```json
{"phase": "plan"}
```
Returns modes:
| Mode | Description |
|------|-------------|
| `sso_oidc` | Configure OIDC single sign-on |
| `sso_saml` | Configure SAML single sign-on |
| `session_timeout` | Configure session timeout settings |
| `login_banner` | Configure login banner message |
| `password_policy` | Configure password policy rules |
| `fips` | Configure FIPS compliance preference |

### Phase 1b — Get fields for a mode:
```json
{"phase": "plan", "mode": "session_timeout"}
```
All modes require:
- `tenant_uid` (string, required) — Tenant UID

Plus a `spec` object specific to the mode:

| Mode | Spec description |
|------|-----------------|
| `sso_oidc` | V1TenantOidcClientSpec with issuerUrl, clientId, clientSecret, etc. |
| `sso_saml` | V1TenantSamlRequestSpec with idpMetadata, entityId, etc. |
| `session_timeout` | V1AuthTokenSettings with accessTokenExpiry, refreshTokenExpiry |
| `login_banner` | V1LoginBannerSettings with message and isEnabled |
| `password_policy` | V1TenantPasswordPolicyEntity with minLength, requireUppercase, etc. |
| `fips` | V1FipsSettings with isFipsEnabled |

### Phase 2 — Apply:
```json
{"phase": "apply", "mode": "session_timeout", "tenant_uid": "tenant-uid", "spec": {"accessTokenExpiry": 60, "refreshTokenExpiry": 1440}}
```

**Example — user says "what platform settings can I configure?":**
Call: `configure_platform` with `{"phase": "plan"}`

**Example — user says "show me the session timeout settings":**
Call: `configure_platform` with `{"phase": "plan", "mode": "session_timeout"}`

## LIMITATIONS
- FIPS cannot be disabled once enabled.
- All apply operations require the `tenant_uid`. There is no tool to look up tenant_uid automatically — the user must provide it.
- This tool only updates settings, it does NOT read/get current values.
