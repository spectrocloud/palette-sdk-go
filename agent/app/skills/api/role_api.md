# role_api

Create, get, update, and delete roles

**get_role_by_id** — Get role by UID
`{"mode": "get_role_by_id", "uid": "role-uid-here"}`

**create_role** — Create a custom role
`{"mode": "create_role", "role": {"metadata": {"name": "custom-viewer"}, "spec": {"permissions": ["cluster.get", "cluster.list"]}}}`

**update_role** — Update an existing role
`{"mode": "update_role", "uid": "role-uid", "role": {"metadata": {"name": "custom-viewer"}, "spec": {"permissions": ["cluster.get"]}}}`

**delete_role** — Delete a role
`{"mode": "delete_role", "uid": "role-uid-here"}`

**get_role** — Get role by name
`{"mode": "get_role", "name": "Tenant Admin"}`

