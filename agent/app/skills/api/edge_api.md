# edge_api

Manage edge hosts and registration tokens for edge cluster deployments.

## Registration Tokens

**create_token** — Create an edge registration token
`{"mode": "create_token", "token_name": "edge-token-prod"}`
*Returns token_uid and token_value. The token_value is used in edge host user-data.*

**get_token** — Get a token by name
`{"mode": "get_token", "token_name": "edge-token-prod"}`

**delete_token** — Delete a token
`{"mode": "delete_token", "uid": "token-uid"}`

## Edge Hosts

**list_edge_hosts** — List all registered edge hosts
`{"mode": "list_edge_hosts"}`

**get_edge_host** — Get edge host details
`{"mode": "get_edge_host", "uid": "host-uid"}`
`{"mode": "get_edge_host", "name": "edge-host-01"}`

**get_edge_hosts_by_tags** — Find edge hosts by tags
`{"mode": "get_edge_hosts_by_tags", "tags": {"env": "production", "site": "edge-01"}}`

**delete_edge_host** — Delete an edge host
`{"mode": "delete_edge_host", "uid": "host-uid"}`
