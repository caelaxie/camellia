# camellia

## Custom golangci-lint

Build the custom linter binary from the repo-local module plugin:

```bash
env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" \
  go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4 custom
```

Run the generated binary with a repo-local cache:

```bash
env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" \
  GOLANGCI_LINT_CACHE="$(pwd)/.golangci-cache" ./custom-gcl run ./...
```
