# camellia

Camellia provides an `abbrcase` module plugin for `golangci-lint`. It flags all-caps abbreviation runs in project-defined Go identifiers, for example `UserID` -> `UserId`, and ignores imported symbols.

## Integrate Camellia Into Another Project

Use Camellia through golangci-lint's module plugin system.

### Prerequisites

- Go
- `git`
- `golangci-lint` `v2.11.4`

### 1. Add a custom golangci build config

Create a `.custom-gcl.yml` file in the root of the project that will consume Camellia.

#### Option A: Remote module integration

```yaml
version: v2.11.4
plugins:
  - module: github.com/caelaxie/camellia
    import: github.com/caelaxie/camellia/pkg/abbrcase/plugin
    version: <go-pseudo-version>
```

Camellia does not currently publish Git tags, so pin a Go pseudo-version for the commit you want.

#### Option B: Local checkout integration

```yaml
version: v2.11.4
plugins:
  - module: github.com/caelaxie/camellia
    import: github.com/caelaxie/camellia/pkg/abbrcase/plugin
    path: ../camellia
```

### 2. Enable the `abbrcase` linter in `.golangci.yml`

If you already have a `.golangci.yml`, merge this into it.

Additive snippet:

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

Fresh-file example only:

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

### 3. Build the custom golangci-lint binary

Run this from the consumer repository:

```bash
golangci-lint custom
```

If you change the local Camellia checkout, rebuild `custom-gcl` so the binary picks up the updated plugin code.

### 4. Run the linter

```bash
./custom-gcl run ./...
```

### 5. Verify the integration

Run the custom binary on code containing an identifier such as `UserID`:

```bash
./custom-gcl run ./...
```

You should get a diagnostic that includes the suggested rename, for example `UserId`.

### Troubleshooting

- `abbrcase` must be configured with `type: module` in `.golangci.yml`.
- The plugin import path must be `github.com/caelaxie/camellia/pkg/abbrcase/plugin`.
- `./custom-gcl` must be rebuilt after changing `.custom-gcl.yml` or local Camellia plugin code.

## Maintainer Workflow For This Repo

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

## References

- [golangci-lint Module Plugin System](https://golangci-lint.run/docs/plugins/module-plugins/)
- [.golangci.yml](.golangci.yml)
