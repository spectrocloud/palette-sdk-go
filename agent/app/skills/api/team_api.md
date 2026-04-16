# team_api

Create, get, update, delete teams, and assign roles

**assign_team_project_role** — Assign project roles to a team
`{"mode": "assign_team_project_role", "uid": "team-uid", "project_roles": {"projects": [{"projectUid": "proj-uid", "roles": [{"uid": "role-uid"}]}]}}`

**assign_team_tenant_role** — Assign tenant roles to a team
`{"mode": "assign_team_tenant_role", "uid": "team-uid", "tenant_roles": {"roleUids": ["role-uid-1"]}}`

**create_team** — Create a new team
`{"mode": "create_team", "team": {"metadata": {"name": "dev-team"}, "spec": {"users": ["user-uid-1"]}}}`

**get_team** — Get team details by UID
`{"mode": "get_team", "uid": "team-uid-here"}`

**get_team_by_name** — Get team details by name
`{"mode": "get_team_by_name", "name": "dev-team"}`

**update_team** — Update an existing team
`{"mode": "update_team", "uid": "team-uid", "team": {"metadata": {"name": "dev-team"}, "spec": {"users": ["user-1", "user-2"]}}}`

**delete_team** — Delete a team
`{"mode": "delete_team", "uid": "team-uid-here"}`

