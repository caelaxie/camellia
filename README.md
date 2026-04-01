# camellia

`camellia` is a `golangci-lint` module plugin that enforces camel-case abbreviations in Go identifiers declared by your project.

It flags names such as `APIError`, `UserID`, and `HTTPClient`, and suggests `ApiError`, `UserId`, and `HttpClient` instead.

## Integrate Into Another Project

Use Camellia through golangci-lint's module plugin system.

### Prerequisites

- Go
- `git`
- `golangci-lint` `v2.11.4`

### 1. Create `.custom-gcl.yml`

Remote integration:

```yaml
version: v2.11.4
plugins:
  - module: github.com/caelaxie/camellia
    import: github.com/caelaxie/camellia/pkg/abbrcase/plugin
    version: v0.0.1
```

Pin the plugin with a normal Go module semver tag such as `v0.0.1`.

Local checkout integration:

```yaml
version: v2.11.4
plugins:
  - module: github.com/caelaxie/camellia
    import: github.com/caelaxie/camellia/pkg/abbrcase/plugin
    path: ../camellia
```

### 2. Enable `abbrcase` in `.golangci.yml`

Merge this into your existing config:

```yaml
linters:
  enable:
    - abbrcase
  settings:
    custom:
      abbrcase:
        type: module
        description: Enforce camel-case abbreviations for project-defined Go identifiers.
        original-url: github.com/caelaxie/camellia
```

Fresh-file example:

```yaml
version: "2"

linters:
  default: none
  enable:
    - abbrcase
  settings:
    custom:
      abbrcase:
        type: module
        description: Enforce camel-case abbreviations for project-defined Go identifiers.
        original-url: github.com/caelaxie/camellia
```

### 3. Build and run

Build the custom binary from the consumer repository:

```bash
golangci-lint custom
```

Run it:

```bash
./custom-gcl run ./...
```

If you change `.custom-gcl.yml` or a local Camellia checkout, rebuild `custom-gcl`.

### 4. Verify

Run the custom binary on code containing an identifier such as `UserID`. You should get a diagnostic that includes the suggested rename, for example `UserId`.

### Troubleshooting

- `abbrcase` must be configured with `type: module` in `.golangci.yml`.
- The plugin import path must be `github.com/caelaxie/camellia/pkg/abbrcase/plugin`.
- `./custom-gcl` must be rebuilt after changing `.custom-gcl.yml` or local Camellia plugin code.

## Rule

The bundled linter is `abbrcase`.

It reports project-defined declarations that use all-caps acronym runs:

```go
type APIError struct{}    // want ApiError
type HTTPClient struct{}  // want HttpClient

func ParseURL(userID string) {} // want ParseUrl, userId
```

External symbols are ignored, so dependencies like `http.Client` are not rewritten.

## Build

Build the repo-local `golangci-lint` binary with the module plugin:

```bash
env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" \
  go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4 custom
```

The build uses [`.custom-gcl.yml`](./.custom-gcl.yml), which registers `github.com/caelaxie/camellia/pkg/abbrcase/plugin`.

## Run

Run the generated binary with the repo-local cache:

```bash
env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" \
  GOLANGCI_LINT_CACHE="$(pwd)/.golangci-cache" ./custom-gcl run ./...
```

The repo's [`.golangci.yml`](./.golangci.yml) enables `abbrcase` like this:

```yaml
version: "2"

linters:
  default: none
  enable:
    - abbrcase
```

## Test

```bash
go test ./...
```

## Releases

The Camellia plugin module can be pinned with semver tags such as `v0.0.1`.

Tagged `custom-gcl` binaries are published separately from `custom-gcl/v*` tags.

The `custom-gcl` release version identifies the packaged binary distribution, not the Camellia Go module version used in a consumer's `.custom-gcl.yml`. Each release is built against the `golangci-lint` toolchain pinned in [`.custom-gcl.yml`](./.custom-gcl.yml), currently `v2.11.4`. If that pin changes, publish a new `custom-gcl` release line so the distributed binary stays aligned with the bundled toolchain.

Each GitHub Release publishes only:

- `custom-gcl_<release>_linux_amd64`
- `custom-gcl_<release>_darwin_amd64`
- `custom-gcl_<release>_darwin_arm64`
- `custom-gcl_<release>_checksums.txt`

The checksum manifest contains SHA-256 sums for every attached binary.

## References

- [golangci-lint Module Plugin System](https://golangci-lint.run/docs/plugins/module-plugins/)
- [.golangci.yml](.golangci.yml)
