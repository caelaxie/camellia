## Context

This repository does not yet have Go source, lint configuration, or a custom analyzer framework. The change introduces a project-specific `golangci-lint` rule that enforces camel-case abbreviations for all project-defined identifier declarations, ignores imported/library symbols, and includes the corrected name in the diagnostic message.

The implementation needs to work within Go's static analysis model and fit cleanly into `golangci-lint` so the rule can be run in normal developer workflows. Because the rule applies broadly across declaration kinds, the design also needs a precise definition of where identifiers come from and how to avoid flagging dependency code.

## Goals / Non-Goals

**Goals:**
- Enforce a single abbreviation-casing rule across all project-defined identifier declarations.
- Ignore identifiers that originate from the standard library or third-party/imported packages.
- Include the corrected camel-case replacement name directly in each diagnostic message.
- Make the rule runnable through the repository's `golangci-lint` workflow and verifiable with regression tests.

**Non-Goals:**
- Automatically rewrite code or perform refactor-grade renames across references.
- Support acronym allowlists or exception handling for selected terms.
- Enforce naming rules outside Go code.

## Decisions

### Use a `go/analysis` analyzer as the rule engine
The linter will be implemented as a custom `go/analysis` analyzer. This matches `golangci-lint`'s preferred extension model and gives access to package syntax, type information, and source positions in a way that is testable with `analysistest`.

Alternatives considered:
- Regex or text scanning: rejected because it cannot reliably distinguish declarations from references or imported symbols.
- A standalone CLI checker: rejected because it would duplicate package loading and would not integrate cleanly with `golangci-lint`.

### Inspect declared objects instead of raw identifier tokens
The analyzer will derive candidate names from declaration objects recorded in `pass.TypesInfo.Defs`, then report only on objects whose positions belong to the package files being analyzed. This keeps the rule focused on declared project identifiers across package-level and function-level scopes while excluding imported references.

Alternatives considered:
- Walking every `ast.Ident`: rejected because it would require additional filtering to avoid references and selectors, increasing false-positive risk.
- Scanning exported declarations only: rejected because the requirement covers all identifier declarations in project code.

### Treat every uppercase abbreviation run as a violation
The rule will normalize any contiguous all-uppercase abbreviation run into camel-case form by preserving the first letter of the run in uppercase and lowercasing the rest. This applies without exceptions, so names such as `ID`, `URL`, `HTTP`, and `API` are handled the same way as longer uppercase runs.

Alternatives considered:
- Built-in acronym allowlist: rejected because the requested policy is explicitly strict with no exceptions.
- Configurable dictionary-based matching: rejected for v1 because it adds configuration surface without matching the requested behavior.

### Provide suggestion text in the diagnostic message only
Each diagnostic will include the corrected spelling in its message, for example reporting that `APIError` should be renamed to `ApiError`. The analyzer will not emit `analysis.SuggestedFix` edits or attempt multi-file rename automation.

Alternatives considered:
- `analysis.SuggestedFix`: rejected because the requested behavior is suggestion text only, not automatic code rewriting.
- Refactor-grade rename support: rejected because it is substantially more complex and out of scope for lint-only enforcement.

### Integrate through golangci-lint's module plugin system
The custom rule will be exposed through a golangci module plugin rather than the `.so` Go plugin system. The module plugin path is the recommended integration model and avoids runtime plugin buildmode/version fragility.

Alternatives considered:
- Go `.so` plugin: rejected because it is more brittle and requires stricter binary compatibility management.
- Forking `golangci-lint`: rejected because it creates unnecessary maintenance overhead for a single project rule.

### Validate behavior with analysistest fixtures and lint wiring checks
Regression coverage will use `analysistest` fixtures that include valid and invalid project declarations across declaration kinds, along with imported-library references that must remain unreported. End-to-end verification will also build/run the configured `golangci-lint` path to confirm the analyzer is discoverable and active.

Alternatives considered:
- Unit testing string normalization only: rejected because it would miss declaration-scope and imported-symbol filtering behavior.

## Risks / Trade-offs

- [Broad declaration coverage increases edge cases] -> Mitigation: cover locals, params, results, type parameters, fields, and test files in analysistest fixtures.
- [Strict no-exception policy may create more findings than typical Go style] -> Mitigation: make the rule behavior explicit in diagnostics, proposal, specs, and tests so enforcement is predictable.
- [Message-only suggestions do not provide one-click fixes] -> Mitigation: ensure diagnostics always include the exact corrected spelling so developers do not need to infer the rename.
- [golangci module-plugin wiring can fail if misconfigured] -> Mitigation: add repository config and an end-to-end verification step that builds and runs the custom linter path.

## Migration Plan

1. Add the analyzer package, plugin registration, and repository `golangci-lint` configuration.
2. Add regression tests covering valid names, invalid names, and imported-symbol exclusions.
3. Verify `go test ./...` and the repository `golangci-lint` workflow both execute the rule successfully.
4. Roll back by removing the custom linter registration and config if the rule proves unusable during initial adoption.

## Open Questions

None currently.
