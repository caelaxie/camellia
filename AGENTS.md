# Repository Instructions

## Naming

- For Go identifiers declared in project code, abbreviations use camel-case, not all-caps runs.
- There are no acronym exceptions. Use `Api`, `Http`, `Id`, and `Url`, not `API`, `HTTP`, `ID`, or `URL`.

## Scope

- `abbrcase` enforces this rule for identifiers declared in project files. Keep code, tests, and docs aligned with that behavior.

## Validation

- Build the repo-local custom `golangci-lint` binary with:
  `env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4 custom`
- Run lint with:
  `env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" GOLANGCI_LINT_CACHE="$(pwd)/.golangci-cache" ./custom-gcl run ./...`
- Run tests with:
  `go test ./...`
