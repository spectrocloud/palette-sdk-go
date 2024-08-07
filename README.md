# palette-sdk-go

Palette SDK for the Go programming language.

## Usage

A project-scoped client can be instantiated as follows:

```go
pc := client.New(
    client.WithPaletteURI(host),
    client.WithAPIKey(apiKey),
    client.WithScopeProject(projectUid),
)
```

Switch from a project-scoped client to a tenant-scoped client:

```go
client.WithScopeTenant()(pc)
```
Note that the above will only succeed if original client's credentials are associated with a user who is authorized as a tenant administrator.

Switch from a tenant-scoped client to a project-scoped client - or from one project to another:

```go
client.WithScopeProject(projectUid)(pc)
```

### Next Steps
- Refer to the [examples](examples/) to get started quickly.
- Refer to [client.go](client/client.go) for all possible client configuration options.
- Refer to [terraform-provider-spectrocloud](https://github.com/spectrocloud/terraform-provider-spectrocloud) for additional usage examples.

## Contributing

All contributions are welcome! Feel free to reach out on the [Spectro Cloud community Slack](https://spectrocloudcommunity.slack.com/join/shared_invite/zt-g8gfzrhf-cKavsGD_myOh30K24pImLA#/shared-invite/email).

Make sure `pre-commit` is [installed](https://pre-commit.com#install).

Install the `pre-commit` scripts:

```console
pre-commit install --hook-type commit-msg
pre-commit install --hook-type pre-commit
```
