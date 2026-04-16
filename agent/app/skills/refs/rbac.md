# RBAC & Access Control — Deep Reference

## Role-Based Access Control Model

Palette uses a hierarchical RBAC model:

```
Tenant
├── Tenant Roles (org-wide permissions)
│   └── Applied to: Users or Teams
└── Projects
    └── Project Roles (project-scoped permissions)
        └── Applied to: Users or Teams
```

## Built-in Roles

Palette ships with built-in roles at both tenant and project level:

### Tenant-Level Roles
- **Tenant Admin** — Full access to all tenant resources
- **Tenant Viewer** — Read-only access to tenant resources

### Project-Level Roles
- **Project Admin** — Full access within a project
- **Project Editor** — Create/update resources, no delete
- **Project Viewer** — Read-only access within a project
- **Cluster Admin** — Full cluster management within a project

Custom roles can be created with granular permissions.

## User Management

- Users are created at the **tenant** level
- Users can be assigned roles at tenant level, project level, or both
- A user with no role assignments has no access
- User creation: `onboard_user` tool (creates user + assigns roles atomically)
- User lookup/delete: `manage_users` tool

## Team Management

- Teams are collections of Users
- Roles assigned to a Team apply to all members
- Teams can be **synced from OIDC/SAML** identity providers
  - When SSO is configured, group memberships from the IdP map to Palette Teams
  - This enables automatic role assignment based on IdP groups
- Team operations: `manage_team` tool

## OIDC / SAML Integration

- Configured at the **platform level** via `configure_platform` tool
- OIDC: issuer URL, client ID, client secret, scopes, callback URL
- SAML: IdP metadata (URL or XML), entity ID, assertion consumer URL
- Once configured, users can log in via SSO
- IdP group → Palette Team mapping enables automated access provisioning

## API Key Authentication

- API keys are created per user at the tenant level
- Used for SDK and programmatic access (what the tool server uses)
- API key inherits the permissions of the user who created it
- A user can have multiple API keys
- API keys are tenant-scoped and work across all projects the user has access to

## Scope Implications

| Operation | Scope | Who can do it |
|-----------|-------|--------------|
| Create user | Tenant | Tenant Admin |
| Create team | Tenant | Tenant Admin |
| Assign tenant role | Tenant | Tenant Admin |
| Assign project role | Project | Tenant Admin or Project Admin |
| Create cluster | Project | Project Admin or Project Editor |
| View cluster | Project | Any project role |
| Configure SSO | Tenant | Tenant Admin |
| Manage API keys | Tenant | Key owner or Tenant Admin |
