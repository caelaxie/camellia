## MODIFIED Requirements

### Requirement: Lint all project identifier declarations
The system SHALL analyze all Go identifier declarations defined in project code, including package-level declarations, struct fields, interface methods, local variables, constants, parameters, named results, type parameters, and declarations in `_test.go` files, except for declarations in files or directories excluded by `camellia` configuration.

#### Scenario: Project declaration is analyzed
- **WHEN** `golangci-lint` runs on a project package containing a declared identifier such as a type, field, local variable, parameter, result, or type parameter in a non-excluded file
- **THEN** the `camellia` linter evaluates that declaration for abbreviation-casing violations

#### Scenario: Excluded declaration is not analyzed
- **WHEN** `golangci-lint` runs on a project package containing a declared identifier in a file or directory matched by `camellia`'s configured exclude list
- **THEN** the `camellia` linter does not evaluate that declaration for abbreviation-casing violations

## ADDED Requirements

### Requirement: Support camellia-specific path exclusions
The system SHALL allow `camellia` to be configured with an `exclude` list under `linters.settings.custom.camellia.settings`, and it SHALL apply those exclusions only to `camellia` diagnostics.

#### Scenario: Excluded file is skipped
- **WHEN** `camellia` is configured with an exclude entry that matches a specific file containing identifier declarations
- **THEN** `camellia` skips reporting diagnostics from that file

#### Scenario: Bare directory entry excludes subtree
- **WHEN** `camellia` is configured with a bare directory exclude entry
- **THEN** `camellia` treats that entry as excluding the directory and all files beneath it

#### Scenario: Recursive glob excludes subtree
- **WHEN** `camellia` is configured with an exclude entry that uses a recursive glob pattern
- **THEN** `camellia` skips reporting diagnostics from matching files beneath the targeted subtree

### Requirement: Validate exclude patterns during plugin initialization
The system SHALL validate `camellia` exclude patterns before analysis begins and SHALL fail initialization if any configured pattern is malformed.

#### Scenario: Malformed exclude pattern is rejected
- **WHEN** `camellia` is configured with a malformed exclude pattern
- **THEN** plugin initialization fails with a configuration error instead of silently ignoring the invalid pattern
