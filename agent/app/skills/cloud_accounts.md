# Cloud Accounts

## setup_cloud_account (two-phase)

### Phase 1a — List available providers:
```json
{"phase": "plan"}
```
Returns field: `provider` (string, required, enum: aws, azure, gcp, vsphere, maas, cloudstack).

### Phase 1b — Get required fields for a provider:
```json
{"phase": "plan", "provider": "aws"}
```

Returns `data.fields[]` — a `spec` object is required. Structure per provider:

| Provider | Spec description |
|----------|-----------------|
| `aws` | Object with metadata.name and spec containing accessKeyId, secretAccessKey, and optionally arn for STS |
| `azure` | Object with metadata.name and spec containing tenantId, clientId, clientSecret, subscriptionId |
| `gcp` | Object with metadata.name and spec containing jsonCredentials (service account key JSON) |
| `vsphere` | Object with metadata.name and spec containing vcenterServer, username, password |
| `maas` | Object with metadata.name and spec containing apiEndpoint, apiKey |
| `cloudstack` | Object with metadata.name and spec containing apiUrl, apiKey, secretKey |

### Phase 2 — Create account:
```json
{"phase": "apply", "provider": "aws", "spec": {"metadata": {"name": "my-aws"}, "spec": {"accessKeyId": "...", "secretAccessKey": "..."}}}
```
Returns: `data.account_uid`, `data.provider`.

**Example — user says "what do I need to set up an AWS account?":**
Call: `setup_cloud_account` with `{"phase": "plan", "provider": "aws"}`

**Example — user says "what cloud providers are supported?":**
Call: `setup_cloud_account` with `{"phase": "plan"}`

## LIMITATIONS
- **NO tool to list, get, update, or delete existing cloud accounts.** This tool only CREATES accounts.
- If a user asks "do I have cloud accounts?", "list my accounts", or "show my AWS accounts", tell them this is not yet available and suggest the Palette UI.
