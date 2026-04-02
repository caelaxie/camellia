## Purpose

Define the repository requirement for enforcing camel-case abbreviation naming in project-defined Go identifiers through the `camellia` golangci-lint workflow.

## Requirements

### Requirement: Lint all project identifier declarations
The system SHALL analyze all Go identifier declarations defined in project code, including package-level declarations, struct fields, interface methods, local variables, constants, parameters, named results, type parameters, and declarations in `_test.go` files.

#### Scenario: Project declaration is analyzed
- **WHEN** `golangci-lint` runs on a project package containing a declared identifier such as a type, field, local variable, parameter, result, or type parameter
- **THEN** the `camellia` linter evaluates that declaration for abbreviation-casing violations

### Requirement: Enforce camel-case abbreviations with no exceptions
The system SHALL reject any project-defined identifier declaration that contains an all-uppercase abbreviation run and SHALL require the abbreviation to be written in camel-case form, with no exceptions for common acronyms such as `ID`, `URL`, `HTTP`, or `API`.

#### Scenario: Common acronym is rejected
- **WHEN** a project-defined identifier declaration uses a name such as `UserID`, `APIError`, or `HTTPClient`
- **THEN** the linter reports a violation for that declaration

### Requirement: Ignore imported and library symbols
The system SHALL only report violations for identifiers declared in project code and SHALL NOT report violations for symbols originating from imported packages, standard library packages, or third-party libraries.

#### Scenario: Imported symbol is not reported
- **WHEN** project code references an imported symbol such as `http.Client` or `externalpkg.APIError`
- **THEN** the linter does not report a violation for the imported symbol name

### Requirement: Include corrected spelling in diagnostics
The system SHALL include the corrected camel-case replacement name in each diagnostic message for an abbreviation-casing violation.

#### Scenario: Diagnostic includes suggested name
- **WHEN** the linter reports a violation for a declaration such as `APIError`
- **THEN** the diagnostic message includes the corrected spelling `ApiError`

### Requirement: Run through repository golangci-lint workflow
The system SHALL be integrated into the repository's `camellia` golangci-lint workflow so the abbreviation-casing rule runs as part of project linting.

#### Scenario: Repository lint run executes custom rule
- **WHEN** a developer runs the repository's configured `camellia` workflow
- **THEN** the `camellia` linter runs and emits diagnostics for project-defined violations
