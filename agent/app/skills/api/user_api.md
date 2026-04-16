# user_api

List, look up, delete, and assign roles to platform users

**assign_tenant_role** — Assign tenant-level roles to an existing user
`{"mode": "assign_tenant_role", "email": "jane@example.com", "tenant_roles": {"roleUids": ["role-uid-1"]}}`
Resolves by `uid`, `email`, `name`.
*User can be identified by uid, email, or name — resolved automatically*
*Role UIDs are required — ask the user if unknown*

**assign_project_role** — Assign project-level roles to an existing user
`{"mode": "assign_project_role", "email": "jane@example.com", "project_roles": {"projects": [{"projectUid": "proj-uid", "roles": [{"uid": "role-uid"}]}]}}`
Resolves by `uid`, `email`, `name`.
*User can be identified by uid, email, or name — resolved automatically*

**list_users** — List all users in the tenant
`{"mode": "list_users"}`

**get_user_by_id** — Find a user by their UID
`{"mode": "get_user_by_id", "uid": "user-uid-here"}`

**get_user_by_email** — Find a user by their email address
`{"mode": "get_user_by_email", "email": "jane@example.com"}`

**get_user_by_name** — Find a user by their display name
`{"mode": "get_user_by_name", "name": "Jane Doe"}`

**delete_user** — Delete a user by UID
`{"mode": "delete_user", "uid": "user-uid-here"}`

