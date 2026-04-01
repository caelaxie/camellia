# Repository Instructions

## Go Naming

- For Go identifiers declared in project code, abbreviations must use camel-case, not all-caps acronym runs.
- There are no acronym exceptions. Use `Api`, `Http`, `Id`, `Url`, not `API`, `HTTP`, `ID`, `URL`.

## Lint Workflow

- Build the repo-local custom golangci binary with:
  `env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4 custom`
- Run lint with:
  `env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" GOLANGCI_LINT_CACHE="$(pwd)/.golangci-cache" ./custom-gcl run ./...`
