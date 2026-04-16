# api_key_api

Create, list, and delete API keys

**create_api_key** — Create a new API key
`{"mode": "create_api_key", "name": "ci-deploy-key", "expiry_hours": 720}`
*The API key value is only shown once at creation — save it securely*

**delete_api_key** — Delete an API key
`{"mode": "delete_api_key", "uid": "key-uid-here"}`

**list_api_keys** — List all API keys
`{"mode": "list_api_keys"}`

