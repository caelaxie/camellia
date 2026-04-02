# Repository Instructions

## Naming

- For Go identifiers declared in project code, abbreviations use camel-case, not all-caps runs.
- There are no acronym exceptions. Use `Api`, `Http`, `Id`, and `Url`, not `API`, `HTTP`, `ID`, or `URL`.

## Scope

- `camellia` enforces this rule for identifiers declared in project files. Keep code, tests, and docs aligned with that behavior.
- Preserve analyzer coverage across project declarations and `_test.go` files unless the user asks to change lint scope.
- Keep imported and third-party symbols out of `camellia` diagnostics; the fixture coverage in `internal/camellia/testdata/` relies on that behavior.
- Treat `internal/camellia/testdata/` as `analysistest` fixture input, not normal package code. Keep its expected diagnostics aligned with analyzer changes.

## Generated Files

- Treat `.gocache/`, `.gomodcache/`, `.gopath/`, `.golangci-cache/`, and `camellia` as generated artifacts. Do not inspect or edit them unless the task is specifically about build outputs or cache behavior.

## Documentation

- Do not update `README.md` unless the user explicitly asks for a README change.

## Versioning

- Treat `.custom-gcl.yml` as the source of truth for the pinned `golangci-lint` version.
- Keep that pin synchronized across validation commands, `.github/workflows/camellia.yml`, and `.github/workflows/camellia-release.yml`.
- Only update README version examples when the user explicitly asked for a README change.
- Keep Camellia Go module tags as normal semver tags like `v0.0.1`, and keep packaged binary release tags under `camellia/v*`. They are separate version lines and should not be conflated.

## Testing

- When changing abbreviation detection, update both `internal/camellia/normalize_test.go` and the `analysistest` fixtures under `internal/camellia/testdata/` together.
- If sandboxed `go test ./...` cannot write to the global Go cache, rerun with repo-local `GOPATH`, `GOMODCACHE`, and `GOCACHE` instead of changing test behavior.

## Validation

- Prefer repo-local Go caches for validation commands:
  `env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" ...`
- Preserve the repository lint policy in `.golangci.yml` unless the user asks to change it: analyze test files, keep `modules-download-mode: readonly`, and continue excluding `^internal/camellia/testdata/` from reported issues.
- Build the repo-local custom `golangci-lint` binary with:
  `env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" go run github.com/golangci/golangci-lint/v2/cmd/golangci-lint@v2.11.4 custom`
- Run lint with:
  `env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" GOLANGCI_LINT_CACHE="$(pwd)/.golangci-cache" ./camellia run ./...`
- Run tests with:
  `env GOPATH="$(pwd)/.gopath" GOMODCACHE="$(pwd)/.gomodcache" GOCACHE="$(pwd)/.gocache" go test ./...`
- If you change plugin registration, the custom build flow, or release packaging, also smoke-test the built binary with:
  `./camellia version`
  `./camellia linters -c .golangci.yml | grep -q '^camellia:'`
- Preserve the lint output contract unless the user asks otherwise: text output to stdout and SARIF output at `.golangci-cache/golangci-lint.sarif`.
