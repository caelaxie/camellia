## Why

The project needs a consistent Go naming rule for abbreviations so identifiers do not mix all-caps acronyms with camel-case naming. Adding this as a project-only `golangci-lint` rule prevents drift in new code and gives developers actionable rename suggestions instead of forcing them to infer the preferred spelling themselves.

## What Changes

- Add a custom Go analyzer that checks all identifier declarations defined in project code, including package-level declarations, type members, locals, parameters, results, type parameters, and test code, for all-caps abbreviation runs and enforces camel-case abbreviations such as `ApiError`, `HttpClient`, and `UserId`.
- Apply the rule with no acronym exceptions, so common forms such as `ID`, `URL`, `HTTP`, and `API` are also rewritten to camel-case forms like `Id`, `Url`, `Http`, and `Api`.
- Limit enforcement to identifiers defined in project code and ignore symbols originating from libraries or imported dependencies.
- Emit diagnostics with the corrected replacement name in the diagnostic message so each violation includes a concrete rename recommendation without requiring automatic code rewriting.
- Integrate the analyzer into the repository's `golangci-lint` workflow and cover it with regression tests.

## Capabilities

### New Capabilities
- `go-abbreviation-casing-lint`: Enforce camel-case abbreviation naming for all Go identifier declarations in project code, with no acronym exceptions, and provide the corrected spelling in lint diagnostics.

### Modified Capabilities

None.

## Impact

Affected areas include the new analyzer package, `golangci-lint` integration/configuration, and regression tests for all declaration kinds in project code, strict no-exception acronym handling, and diagnostic-message suggestion generation. Developer lint output will change by including the corrected replacement name in each abbreviation-casing violation.
