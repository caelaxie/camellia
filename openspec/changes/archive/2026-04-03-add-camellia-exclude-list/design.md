## Context

`camellia` currently runs as a zero-configuration `go/analysis` plugin that reports abbreviation-casing violations for declarations defined in package files loaded by `golangci-lint`. That behavior is intentionally strict, but it leaves no room to suppress findings from generated code, fixtures, or legacy paths that teams may need to keep out of scope during adoption.

This change adds a linter-specific exclude list rather than relying on global `golangci-lint` exclusions. The repository already uses `camellia` as a module plugin and keeps its behavior scoped to project-defined declarations, so the new design needs to preserve that boundary while adding path-based filtering that is predictable in tests and in real lint runs.

## Goals / Non-Goals

**Goals:**
- Add a `camellia` setting that excludes configured files and directory trees from diagnostics.
- Keep exclusion semantics local to `camellia` so other linters continue to see the same files.
- Support recursive glob matching, including `**`, with clear failure behavior for malformed patterns.
- Preserve existing analyzer behavior when no exclude list is configured.

**Non-Goals:**
- Change the naming rule itself or add acronym allowlists.
- Modify global `golangci-lint` path exclusion behavior.
- Provide automatic code fixes or rename tooling.
- Support multiple path anchors or ambiguous matching modes for the same setting.

## Decisions

### Decode excludes through plugin settings and build a configured analyzer

The plugin will extend its settings struct to accept `exclude []string` from `linters.settings.custom.camellia.settings`. `New` will decode that config, validate and compile the matcher once, and construct an analyzer instance with the matcher attached instead of relying only on a single hard-coded analyzer singleton.

This keeps configuration errors in plugin initialization, where they are easier to understand, and avoids recompiling patterns during per-package analysis.

Alternatives considered:
- Read raw settings inside `Run`: rejected because analyzer execution is too late for clean config validation and would repeat setup work.
- Keep the global analyzer singleton and mutate shared state: rejected because it makes tests and future extensions more brittle.

### Anchor exclude patterns to the nearest module root

Exclude entries will be interpreted relative to the nearest `go.mod` for each analyzed file. The analyzer will derive a stable module-relative path, normalize it with `/` separators, and match excludes against that normalized path.

This anchor is available from within the analyzer without depending on `golangci-lint` config loader internals, and it aligns naturally with how this repository structures Go code and tests.

Alternatives considered:
- Anchor patterns to `.golangci.yml`: rejected because the analyzer does not reliably know that path at runtime.
- Accept both module-relative and config-relative paths: rejected because it creates ambiguity and documentation overhead.

### Use `doublestar` for explicit recursive glob semantics

The implementation will use `github.com/bmatcuk/doublestar/v4` to match exclude patterns. Input patterns and candidate file paths will be slash-normalized before matching so behavior is consistent across platforms.

This makes `**` support explicit, avoids custom recursive matching logic, and keeps the matching rules easy to document and test.

Alternatives considered:
- Standard library matching only: rejected because `filepath.Match` does not support the required recursive-glob behavior cleanly.
- Custom recursive matcher: rejected because it adds maintenance cost for behavior a small dependency already provides.

### Treat bare directory entries as recursive subtree excludes

If a configured exclude entry has no glob metacharacters and matches a directory-style prefix, the matcher will treat it as excluding that directory and everything below it, equivalent to appending `/**`. Exact file paths will still exclude only the file itself.

This keeps common configuration concise while preserving predictable behavior for single-file excludes.

Alternatives considered:
- Require explicit `/**` for directories: rejected because it makes the common case noisier with little benefit.
- Split config into separate file and directory lists: rejected because it adds surface area without solving a real ambiguity in this repo.

### Fail fast on malformed patterns

Invalid exclude patterns will cause plugin initialization to return an error rather than being ignored silently.

This prevents partial enforcement caused by typoed patterns and gives immediate feedback during lint setup.

Alternatives considered:
- Ignore malformed patterns: rejected because it hides configuration mistakes.
- Warn and continue: rejected because the plugin interface does not provide a clean warning channel and silent partial behavior is still risky.

## Risks / Trade-offs

- [Module-root discovery may behave differently in unusual workspace layouts] -> Mitigation: keep matching module-relative, cover nested/fixture-style paths in tests, and fail closed when a module root cannot be derived.
- [A new dependency expands the maintenance surface] -> Mitigation: use a small, focused matcher library only for glob evaluation and keep the rest of the logic local.
- [Bare directory shorthand could surprise users expecting exact-path-only semantics] -> Mitigation: document that plain directory entries are recursive and cover both bare-directory and explicit `/**` cases in tests.
- [Path normalization bugs could produce inconsistent matching across platforms] -> Mitigation: normalize both patterns and candidate paths to slash-separated module-relative strings before matching.

## Migration Plan

1. Extend plugin settings and analyzer construction to accept compiled exclude matchers.
2. Add path normalization and exclusion checks before reporting declaration diagnostics.
3. Add regression coverage for zero-config behavior, excluded files/directories, and malformed patterns.
4. Validate `go test ./...` and the custom `golangci-lint` workflow with a config that proves only `camellia` honors the new excludes.
5. Roll back by removing the `exclude` setting and matcher wiring if the feature causes unexpected lint suppression.

## Open Questions

None currently.
