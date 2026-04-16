# Palette Agentic — Architecture Guide

## Overview

Palette Agentic is an AI-powered management assistant for the Spectro Cloud Palette platform. It combines a Go tool server, Python agent backend, and Next.js frontend to provide an interactive chatbot that can query, manage, and orchestrate Palette resources.

```
Browser (Next.js :3000) ──WebSocket──▸ Python (FastAPI :8080) ──HTTP──▸ Go (Tool Server :8090)
                                           │                              │
                                      LLM (Claude)                   Palette API
                                      Web Search                     AWS SDK
```

---

## Layer Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│  Frontend (Next.js)                                              │
│  Chat UI, forms, selects, tables, plan visualization, debug      │
│  Status indicator (spinner), thinking states                     │
├─────────────────────────────────────────────────────────────────┤
│  Agent Backend (Python/FastAPI)                                  │
│  LLM-based router, session management, journey orchestration     │
│  Web search (DuckDuckGo), error handling, status messages        │
├─────────────────────────────────────────────────────────────────┤
│  Tool Server (Go)                                                │
│  HTTP endpoints wrapping workflows → services → APIs             │
├─────────────────────────────────────────────────────────────────┤
│  Services (Go)                                                   │
│  Abstraction over Palette SDK + AWS SDK + Edge APIs              │
├─────────────────────────────────────────────────────────────────┤
│  Palette API / AWS API / External APIs                           │
└─────────────────────────────────────────────────────────────────┘
```

---

## Go Layers

### SDK (Auto-generated)

`api/` — Generated from OpenAPI spec via go-swagger. ~4000 files. **Do not edit manually.**
- `api/client/version1/` — Generated HTTP client operations
- `api/models/` — Generated request/response models

### Services

`services/` — Hand-written abstractions over the SDK client. Each service is an interface with a constructor that takes `*client.V1Client`.

| Service | Package | Scope | Purpose |
|---------|---------|-------|---------|
| Clusters | `services/clusters` | Project | Cluster CRUD, profiles, variables, backup |
| Cloud Accounts | `services/cloudaccounts` | Project/Tenant | Provider credentials (AWS, Azure, GCP, etc.) |
| Profiles | `services/profiles` | Project | Cluster profile CRUD, packs, variables |
| Access | `services/access` | Tenant | Users, teams, roles, API keys |
| Registries | `services/registries` | Project | Pack, Helm, OCI registries |
| Diagnostics | `services/diagnostics` | Project | Compliance scans, alerts, events |
| Platform | `services/platform` | Tenant | Projects, SSO, settings |
| VMO | `services/vmo` | Project | Virtual machine lifecycle (KubeVirt) |
| Edge | `services/edge` | Project | Edge hosts, registration tokens |
| AWS Infra | `services/awsinfra` | External | Direct AWS queries (VPCs, EC2, NLB, security groups) |

### Workflows

Two types:

**Generated** (`workflows/generated/`) — Auto-generated from `manifests/api/*.yaml` by the generator. Each workflow has a mode switch that validates inputs and calls the appropriate service method.

**Hand-written** (`workflows/`) — For tools that need custom logic:
- `aws_infra.go` — Resolves AWS credentials from env vars, dispatches to awsinfra service (VPCs, EC2, NLB, security groups)
- `edge.go` — Edge host and registration token operations. Token reuse: checks for existing active, non-expired token before creating new one

### Tool Server

`cmd/toolserver/` — HTTP server exposing all workflows as `POST /tools/{name}` endpoints.
- Initializes SDK clients (project-scoped + tenant-scoped)
- Creates all services
- Registers all tools (generated + hand-written)
- Routes: `GET /healthz`, `GET /tools`, `POST /tools/{name}`

---

## Manifest System

### API Manifests (`manifests/api/`)

Define atomic tool operations. Each manifest generates:
- Go workflow code (`workflows/generated/*.go`)
- Tool registration (`workflows/generated/register.go`)
- Skill markdown (`agent/app/skills/api/*.md`)

```yaml
kind: api
name: user_api
service: access
agent: access_admin
modes:
  list_users:
    method: GetUsers
    returns: "value_error"
    output: {list: users, table: [name, email, uid]}
    example: '{}'
```

**21 API manifests** covering all Palette resources + AWS infra.

### Journey Manifests (`manifests/journeys/`)

Define multi-step workflows. Generate **skill docs only** (no Go code). The skill docs guide the LLM or Python orchestrator.

```
manifests/journeys/
├── shared/                              # Reusable sub-journeys
│   ├── ensure_cloud_account.yaml
│   ├── ensure_cluster_profile.yaml
│   ├── ensure_edge_registration_token.yaml
│   ├── prepare_aws_edge_host_agent.yaml
│   ├── prepare_aws_edge_host_appliance.yaml
│   ├── register_edge_hosts.yaml
│   └── ensure_registry.yaml
├── eks/
│   └── deploy_eks_cluster.yaml
├── edge/
│   ├── deploy_edge_cluster_agent.yaml
│   └── deploy_edge_cluster_appliance.yaml
└── profiles/
    └── create_cluster_profile.yaml
```

### Defaults (`manifests/defaults/`)

Configuration values consumed by generated skills and Python tools at runtime.

```
manifests/defaults/
├── global.yaml              # Platform-wide (registry, naming, versions)
├── providers/               # Per-cloud-type (packs, cluster config)
│   ├── eks.yaml
│   ├── aws.yaml
│   ├── azure.yaml
│   ├── edge-native.yaml     # Edge packs (no CSI), OS version pinned, 100GB disk
│   ├── edge.yaml
│   └── ... (9 providers)
└── shared/                  # Cross-provider (security, networking, backup)
    ├── security.yaml
    ├── networking.yaml
    └── backup.yaml
```

### Generator

`tools/generate/` — Go CLI that reads manifests and produces code + skills.

```bash
make generate-tools    # Runs generator + compiles Go code
```

**Generates from `manifests/api/`:**
- `workflows/generated/*.go` — Go workflow with mode switch + service calls
- `workflows/generated/register.go` — Tool registration
- `agent/app/skills/api/*.md` — Compact skill docs for LLM context

**Generates from `manifests/journeys/`:**
- `agent/app/skills/journeys/*.md` — Journey guide with steps, variants, defaults

---

## Python Agent

### Router (`agent/app/crews/palette_crew.py`)

Uses a **single LLM call** to classify both intent and specialist agent (replaced hardcoded keyword matching):

```
User message → _classify_with_llm() → (intent, agent_key)
```

| Intent | Route | LLM Calls | Description |
|--------|-------|-----------|-------------|
| **direct** | Single specialist | 1 (pick tool) | Fetch and display: "list clusters", "show projects" |
| **interpret** | Single specialist + analysis | 2 (pick tool + analyze) | Reason about data: "are my clusters healthy?" |
| **journey** | Multi-step orchestration | 1+ | Complex flows: "deploy EKS cluster", "create profile" |

**Automatic interpret escalation:** Any user message that is a question (contains `?` or starts with what/which/how/why/is/are/do/does/can/any) automatically triggers the interpret flow, ensuring the LLM always analyzes results for questions rather than showing raw data.

**Routing hints in the classification prompt:**
- AWS VMs, EC2, VPCs, subnets → infra (aws_infra_api)
- KubeVirt VMs inside a cluster → app_workload (vm_api)
- Packs, profiles, registries, versions → app_workload (profile_api)
- Cluster health, diagnostics → operations
- Users, teams, roles → access_admin

**Journey matching:** Regex patterns match first (fast), with LLM fallback for typo tolerance.

**Fallback:** If LLM classification fails, keyword-based fallback runs.

### Agents and Tool Mapping

| Agent | Tools | Keywords |
|-------|-------|----------|
| **Infra Provisioner** | cluster_common_api, eks_cluster_api, aws_cluster_api, azure_cluster_api, aks_cluster_api, gcp_cluster_api, gke_cluster_api, vmware_cluster_api, maas_cluster_api, edge_cluster_api, cloud_account_api, aws_infra_api, edge_api | deploy, eks, aws, azure, vpc, subnet, edge, ec2, instance |
| **Platform Admin** | project_api, platform_settings_api | project, sso, session, fips, tenant |
| **App & Workload** | profile_api, registry_api, vm_api | profile, registry, pack, version, monitoring, logging, upstream, prometheus, grafana, helm, vm, kubevirt |
| **Operations** | cluster_common_api, diagnostics_api | inspect, backup, compliance, scan, health, diagnostic, monitor |
| **Access Admin** | user_api, team_api, role_api, api_key_api | user, team, role, api key, rbac, permission |

### LLM vs Python Execution

| What | Executed by | Why |
|------|------------|-----|
| Intent + agent classification | LLM (single call, 30 tokens) | Natural language understanding, handles typos/rephrasings |
| Tool selection ("which tool to call?") | LLM | Maps user intent to API |
| Tool input generation ("what JSON to pass?") | LLM | Maps natural language to API params |
| Tool execution (HTTP call to Go server) | Python | Deterministic, reliable |
| Result interpretation ("what does this mean?") | LLM | Reasoning over data |
| Web search (upstream versions, current info) | Python (DuckDuckGo) | LLM knowledge may be stale |
| Error recovery ("what went wrong?") | LLM | Contextual understanding |
| Recommendations ("which option?") | LLM | Domain knowledge |
| Journey step ordering | Python | Must be deterministic and reliable |
| Form rendering and user input collection | Python | Interactive, needs WebSocket sync |
| Profile/account creation flows | Python | Multi-step, needs exact API formats |
| Journey summary | LLM | Natural language generation |

### Web Search

`_web_search()` uses DuckDuckGo (no API key needed) for questions about upstream software versions, releases, and comparisons. Triggered automatically when user asks about "latest version", "upstream", "release", etc.

- 10-second timeout to prevent hangs
- Results fed to LLM as additional context
- Query appends "software release github" to disambiguate from pop-culture results

### Python-Driven Journeys

6 implemented journeys:

| Journey | Handler | Steps |
|---------|---------|-------|
| `create_cloud_account` | `_run_create_account_journey` | Provider picker → credentials form → create |
| `create_cluster_profile` | `_run_create_profile_journey` | Cloud type detection (LLM) → pack selection → create. Edge profiles: no CSI layer, system.uri patched to "NA" for agent mode |
| `deploy_eks_cluster` | `_run_deploy_eks_journey` | Account → profile → config (name, region, instance, nodes, SSH key) → confirm → create with full machinePoolConfig |
| `prepare_and_deploy_edge` | chains prepare → deploy | Combined: launch AWS VMs + deploy edge cluster |
| `prepare_edge_hosts_agent` | `_run_prepare_edge_hosts_agent` | Token (reuse or create) → VPC/subnet/key/SG → config → pre-create hosts in Palette → launch EC2 with user-data |
| `deploy_edge_cluster_agent` | `_run_deploy_edge_cluster` | Select hosts (from artifacts or Palette) → profile → NLB (if AWS) → create cluster |

**Journey features:**
- Plan visualization with step-by-step progress (PlanMessage + PlanUpdateMessage)
- Confirmation dialogs before destructive actions
- Agentic error handling (LLM interprets errors, suggests recovery)
- LLM-powered recommendations for selections
- Session artifacts for data passing between chained journeys
- Token reuse (checks for existing valid token before creating)

### Skills Architecture

```
agent/app/skills/
├── concepts.md                # Shared glossary — loaded into ALL agents
├── refs/                      # Deep domain knowledge (loaded per-agent)
│   ├── provisioning.md
│   ├── cloud_types.md
│   ├── edge.md
│   ├── edge_ports.md
│   ├── vmo.md
│   ├── profiles_ref.md
│   └── rbac.md
├── api/                       # Per-tool reference (generated + hand-written)
│   ├── cluster_common_api.md  # Includes client_filter examples
│   ├── profile_api.md         # Includes search_packs with addOnType, name filters
│   ├── aws_infra_api.md       # Hand-written
│   ├── edge_api.md            # Hand-written
│   └── ... (21+ total)
└── journeys/                  # Generated from manifests/journeys/
    ├── deploying_eks_cluster.md
    ├── deploying_edge_cluster_agent.md
    ├── deploying_edge_cluster_appliance.md
    └── creating_cluster_profile.md
```

### Interactive Tools

Tools support special input flags for rich UI interactions:

| Flag | Behavior | Example |
|------|----------|---------|
| `"interactive": true` | Show results as pick list, wait for selection | `{"mode": "list_accounts", "interactive": true}` |
| `"collect": {fields}` | Show form, wait for submission | `{"collect": {"name": {"type": "string"}}}` |
| `"build_profile": "eks"` | Fetch packs + show combined form | Profile builder |
| `"build_account": "aws"` | Show auth mode picker + credential form | Account builder |
| `"client_filter": {k:v}` | Filter results client-side (case-insensitive substring) | `{"client_filter": {"cloudType": "edge-native", "state": "UnHealthy"}}` |

### Table Rendering

The `PaletteTool._run()` auto-detects list responses and renders them as tables:
- `_find_list()` — finds the list in nested response data
- `_extract_scalars()` → `_collect_scalars()` — flattens nested objects to key-value rows
- Skips noisy fields: annotations, labels, timestamps, permissions, projectMeta, cloudAccountMeta, isAvailable, cost, fips
- Skips `false` booleans to reduce noise
- Extracts `latestVersion` from pack `registries` arrays
- Protects top-level fields from being overwritten by nested duplicates (e.g., metadata.name vs projectMeta.name)
- `client_filter` applies after flattening, before display — filtered results also used for LLM interpretation

### Data Compaction for LLM

When `interpret_result` sends tool data to the LLM, it uses `_compact_tool_result()`:
- Extracts only scalar fields from each item (skips nested objects/lists)
- Skips misleading fields (isAvailable, cost, fips, etc.)
- Allows up to 8000 chars (vs original 3000 raw JSON truncation)
- Explicit instruction: "Analyze ALL items, not just the first few"

### Session & Conversation

- WebSocket session per browser connection
- Split architecture: `_input_listener` (background task reads WS messages) + main handler (processes chat)
- Conversation history maintained (last 20 turns) for context
- Pending input synchronization via `threading.Event` for form/select/confirm responses
- Session artifacts for inter-journey data passing (edge_host_ids, ec2_instances, etc.)

---

## Frontend (Next.js)

### Message Types

| Type | Component | Purpose |
|------|-----------|---------|
| `text` | TextMessage | Markdown-rendered text |
| `table` | TableMessage | Sortable data table with auto-detected columns |
| `form` | FormMessage | Dynamic form with dropdowns, inputs, defaults |
| `select` | SelectMessage | Single or multi-select (radio/checkbox) |
| `confirm` | ConfirmMessage | Yes/No confirmation dialog |
| `plan` | PlanMessage | Visual step-by-step flowchart with status dots |
| `plan_update` | (updates PlanMessage) | Updates step status in existing plan |
| `status` | (controls spinner) | Shows/hides "Analyzing results..." spinner |
| `error` | ErrorMessage | Red error alert |

### Status Indicators

Two independent indicators:
- **`isThinking`** — shown from when user sends message until first server response
- **`statusText`** — shown when server sends `StatusMessage(done=false)`, cleared on `StatusMessage(done=true)`. Used between tool result and LLM analysis.

Both show the same spinner component but are controlled independently, so the "Analyzing results..." spinner persists even after the table arrives.

### Debug Panel

Toggle-able right panel showing real-time execution trace:
- Route decision (intent, agent)
- LLM calls (purpose, timing)
- Tool calls (name, input, result, duration)
- Journey steps

### Palette Branding

Colors, fonts, and assets from the Spectro Cloud optic repository:
- Brand teal `#1F7A78`, sidebar dark `#012121`
- Poppins font family
- Palette logo SVG

---

## Running

### Prerequisites

```bash
# .env file at repo root
export PALETTE_HOST=api.spectrocloud.com
export PALETTE_API_KEY=<your-key>
export PALETTE_PROJECT_UID=<your-project-uid>
export AWS_ACCESS_KEY_ID=<your-aws-key>        # For AWS infra queries
export AWS_SECRET_ACCESS_KEY=<your-aws-secret>
export AWS_REGION=us-east-1
export ANTHROPIC_API_KEY=<your-anthropic-key>
```

### Start Services (3 terminals)

```bash
# Terminal 1: Go tool server (MUST source .env for PALETTE_PROJECT_UID)
source .env && make toolserver

# Terminal 2: Python agent
source .env && cd agent && uvicorn app.main:app --port 8080 --reload

# Terminal 3: Next.js frontend
cd ui && npm run dev
```

Open http://localhost:3000

### Generate After Manifest Changes

```bash
make generate-tools    # Reads manifests → generates Go + skills → compiles
```

---

## Adding Capabilities

### New API operation (on existing resource)

1. Add mode to `manifests/api/{resource}_api.yaml`
2. Run `make generate-tools`
3. Restart tool server

### New resource

1. Create service in `services/{name}/`
2. Create `manifests/api/{name}_api.yaml` (or hand-write workflow + skill)
3. Add to `tools/register.go` and `cmd/toolserver/main.go`
4. Add to agent config in `palette_crew.py`
5. Run `make generate-tools`

### New journey

1. Create shared sub-journeys in `manifests/journeys/shared/`
2. Create journey manifest in `manifests/journeys/{category}/`
3. Add journey pattern to `JOURNEY_PATTERNS` in `palette_crew.py`
4. For Python-driven: add handler in `palette_crew.py`
5. Run `make generate-tools`

### New defaults

1. Edit `manifests/defaults/providers/{cloud_type}.yaml` or `shared/*.yaml`
2. Run `make generate-tools` (embeds in journey skills)
3. Python picks up changes on restart (via `app/defaults.py`)

---

## Known Limitations

1. **Single tool call per query** — `direct` and `interpret` flows call one tool, once. Multi-step data gathering (e.g., "get K8s version for each cluster") requires either a Go-side aggregation mode or the planned agentic tool loop.

2. **K8s version not in cluster list** — The `list_clusters` / `search_clusters` API returns summary data only. K8s version is only in the full `get_cluster` response (in `spec.clusterProfileTemplates[].packs[]`).

3. **LLM classification latency** — Each query makes 2 LLM calls minimum (classify + pick tool). Could be optimized by combining into one call or using a faster model for classification.

4. **Table columns are auto-generated** — No control over which columns appear or their order. The flattener extracts all scalar fields.

5. **No multi-project queries** — Each tool server instance is scoped to one project. Cross-project queries require switching context.

---

## File Map

```
palette-sdk-go/
├── api/                          # Auto-generated SDK (DO NOT EDIT)
├── client/                       # SDK client methods
├── services/                     # Service layer (10 packages)
│   ├── clusters/
│   ├── cloudaccounts/
│   ├── profiles/
│   ├── access/
│   ├── registries/
│   ├── diagnostics/
│   ├── platform/
│   ├── vmo/
│   ├── edge/                     # Edge hosts + registration tokens
│   └── awsinfra/                 # AWS VPCs, EC2, NLB, security groups
├── workflows/                    # Workflow layer
│   ├── workflow.go               # Shared types (Result, Phase, FieldSpec)
│   ├── aws_infra.go              # Hand-written (credential resolution, NLB, SG)
│   ├── edge.go                   # Hand-written (token reuse, batch host create)
│   └── generated/                # Auto-generated from manifests
├── tools/                        # Tool registry + registration
│   ├── tool.go                   # Registry, Tool types
│   ├── register.go               # Registers generated + hand-written tools
│   └── generate/                 # YAML → Go + Markdown generator
├── manifests/                    # Source of truth for tools and journeys
│   ├── api/                      # API manifests (21 files)
│   ├── journeys/                 # Journey manifests
│   │   ├── shared/               # Reusable sub-journeys
│   │   ├── eks/
│   │   ├── edge/
│   │   └── profiles/
│   └── defaults/                 # Configuration defaults
│       ├── global.yaml
│       ├── providers/            # Per cloud type
│       └── shared/               # Cross-provider
├── cmd/toolserver/               # Go HTTP server (3 files)
├── agent/                        # Python backend
│   ├── app/
│   │   ├── main.py               # FastAPI + WebSocket (split input listener)
│   │   ├── config.py             # Settings (model, URLs, log level)
│   │   ├── protocol.py           # Message types (12 types)
│   │   ├── session.py            # WebSocket session + pending input + artifacts
│   │   ├── defaults.py           # Hierarchical defaults loader with caching
│   │   ├── crews/
│   │   │   └── palette_crew.py   # LLM router, agents, 6 journey runners, web search
│   │   ├── tools/
│   │   │   ├── base.py           # PaletteTool (HTTP bridge, table rendering, client_filter, data compaction)
│   │   │   ├── profile_builder.py # Pack fetch + form + values patch (system.uri)
│   │   │   ├── account_builder.py # Multi-provider credential forms
│   │   │   └── edge_userdata.py   # Agent/appliance mode user-data generation
│   │   ├── journeys/
│   │   │   └── agent_layer.py    # Agentic error handling + recommendations + summaries
│   │   └── skills/
│   │       ├── concepts.md
│   │       ├── refs/             # 7 reference docs
│   │       ├── api/              # 21+ tool skill docs
│   │       └── journeys/         # 4 journey skill docs
│   └── pyproject.toml
├── ui/                           # Next.js frontend
│   ├── app/
│   ├── components/
│   │   ├── chat/                 # ChatContainer, ChatInput, MessageList (with ThinkingIndicator)
│   │   ├── messages/             # Text, Table, Form, Select, Confirm, Plan
│   │   ├── sidebar/
│   │   └── debug/
│   ├── hooks/
│   │   └── useChat.ts            # WS client, isThinking, statusText, message state
│   ├── lib/
│   │   ├── types.ts              # Protocol types
│   │   └── websocket.ts          # WS client with reconnect
│   └── public/
├── docker-compose.yaml
├── .env.example
├── Makefile
└── AGENTS.md                     # This file
```
