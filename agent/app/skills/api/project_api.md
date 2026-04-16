# project_api

Create, get, list, update, and delete Palette projects

**get_project_uid** — Get project UID by name
`{"mode": "get_project_uid", "name": "production"}`

**create_project** — Create a new project
`{"mode": "create_project", "project": {"metadata": {"name": "staging"}, "spec": {}}}`

**update_project** — Update an existing project
`{"mode": "update_project", "uid": "project-uid", "project": {"metadata": {"name": "staging-v2"}, "spec": {}}}`

**delete_project** — Delete a project
`{"mode": "delete_project", "uid": "project-uid-here"}`

**list_projects** — List all projects in the tenant
`{"mode": "list_projects"}`

**get_project** — Get project details by UID
`{"mode": "get_project", "uid": "project-uid-here"}`

