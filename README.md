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
    import: github.com/caelaxie/camellia
    version: v0.0.1
```

Pin the plugin with a normal Go module semver tag such as `v0.0.1`.

Local checkout integration:

```yaml
version: v2.11.4
plugins:
  - module: github.com/caelaxie/camellia
    import: github.com/caelaxie/camellia
    path: ../camellia
```

### 2. Enable `camellia` in `.golangci.yml`

Merge this into your existing config:

```yaml
linters:
  enable:
    - camellia
  settings:
    custom:
      camellia:
        type: module
        description: Enforce camel-case abbreviations for project-defined Go identifiers.
        original-url: github.com/caelaxie/camellia
        settings:
          exclude:
            - internal/testdata
            - generated/**/*.go
```

Fresh-file example:

```yaml
version: "2"

linters:
  default: none
  enable:
    - camellia
  settings:
    custom:
      camellia:
        type: module
        description: Enforce camel-case abbreviations for project-defined Go identifiers.
        original-url: github.com/caelaxie/camellia
        settings:
          exclude:
            - internal/testdata
            - generated/**/*.go
```

`exclude` is optional and applies only to `camellia`. Entries are resolved relative to the nearest module root.

- Use a file path such as `internal/legacy/id_aliases.go` to exclude one file.
- Use a bare directory such as `internal/testdata` to exclude that directory and everything under it.
- Use a glob such as `generated/**/*.go` for recursive pattern matches.

### 3. Build and run

Build the custom binary from the consumer repository:

```bash
golangci-lint custom
```

Run it:

```bash
./camellia run ./...
```

If you change `.custom-gcl.yml` or a local Camellia checkout, rebuild `camellia`.

### 4. Verify

Run the custom binary on code containing an identifier such as `UserID`. You should get a diagnostic that includes the suggested rename, for example `UserId`.

### Troubleshooting

- `camellia` must be configured with `type: module` in `.golangci.yml`.
- The plugin import path must be `github.com/caelaxie/camellia`.
- `./camellia` must be rebuilt after changing `.custom-gcl.yml` or local Camellia plugin code.

## Rule

The bundled linter is `camellia`.

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

The build uses [`.custom-gcl.yml`](./.custom-gcl.yml), which registers `github.com/caelaxie/camellia`.

## Run

Run the generated binary with the repo-local cache:

```bash
env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" \
  GOLANGCI_LINT_CACHE="$(pwd)/.golangci-cache" ./camellia run ./...
```

The repo's [`.golangci.yml`](./.golangci.yml) enables `camellia` like this:

```yaml
version: "2"

linters:
  default: none
  enable:
    - camellia
```

## Test

```bash
go test ./...
```

## References

- [golangci-lint Module Plugin System](https://golangci-lint.run/docs/plugins/module-plugins/)
- [.golangci.yml](.golangci.yml)
