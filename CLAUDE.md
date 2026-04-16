# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Overview

Go SDK for the Spectro Cloud Palette API. Module: `github.com/spectrocloud/palette-sdk-go` (Go 1.24.11).

The SDK has two layers:
- **`api/`** — Auto-generated code (models, client operations, transport). ~4000 files generated from OpenAPI spec via go-swagger. **Do not edit manually.**
- **`client/`** — Hand-written SDK logic that wraps the generated client with ergonomic CRUD methods, scope management, and error handling.

## Common Commands

```bash
make test          # Run all tests with coverage
make fmt           # go fmt
make vet           # go vet
make lint          # golangci-lint (v1.64.6)
make reviewable    # Full pre-commit check: fmt + vet + lint + mod tidy + submodule update
make generate      # Regenerate api/ from OpenAPI spec (clones hapi repo, runs swagger + controller-gen)
make check-diff    # Verify branch is clean after reviewable
```

Run a single test:
```bash
go test ./client/herr/ -run TestIsNotFound -v
```

## Architecture

### Client initialization (builder pattern)

```go
pc := client.New(
    client.WithPaletteURI(host),
    client.WithAPIKey(apiKey),
    client.WithScopeProject(projectUID),  // or WithScopeTenant()
)
```

`V1Client` holds the generated `clientv1.ClientService`, a context with scope headers, and auth/TLS config. All options are in `client/client.go`.

### Resource files in `client/`

Each file covers one API resource (e.g., `cluster.go`, `account.go`, `project.go`) and follows a consistent pattern: List, Get, GetByName, Create, Update, Delete methods that call the generated client and unwrap the response payload.

### Scope management

Operations are scoped to either "tenant" or "project". Scope is set via context headers using `ContextForScope(baseCtx, scope, projectUID)`. Most resource methods accept a `scope` string parameter.

### Error handling

- `client/herr/` — Helper predicates: `IsNotFound(err)`, `IsEdgeHostDeviceNotRegistered(err)`, etc.
- `client/apiutil/` — `Is404(err)`, `ToV1ErrorObj(err)` for transport-level errors.
- Transport errors use `api/apiutil/transport.TransportError`.

### Code generation (`api/generate.sh`)

1. Clones hapi repo for the latest OpenAPI spec
2. Flattens spec → generates models and client via go-swagger
3. Generates DeepCopy methods via controller-gen
4. Post-processes edge cases (empty map loops, V1Time alias)

Generated code lives entirely under `api/`. The `api/spec/` dir holds the OpenAPI specs, `api/swagger/` holds custom templates.

## Conventions

- Commit messages must follow [Conventional Commits](https://www.conventionalcommits.org/) (enforced by pre-commit hook).
- Linting excludes generated code (`api/apiutil`, `api/client/version1`) and test files. Config in `.golangci.yaml`.
- Pre-commit hooks: gitleaks (secrets scanning), conventional-pre-commit, golangci-lint.
- Tests use `github.com/stretchr/testify`.
- Use `apiutil.Ptr[T](v)` for pointer wrapping of literals.
