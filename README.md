# camellia

`camellia` is a `golangci-lint` module plugin that enforces camel-case abbreviations in Go identifiers declared by your project.

It flags names such as `APIError`, `UserID`, and `HTTPClient`, and suggests `ApiError`, `UserId`, and `HttpClient` instead.

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
