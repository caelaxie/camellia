## Why

`camellia` currently evaluates every project-defined declaration in loaded package files, which makes adoption harder in repositories that contain generated code, fixtures, or legacy paths that should be left alone. Adding a linter-specific exclude list makes the rule practical to roll out without weakening the naming policy for the rest of the codebase.

## What Changes

- Add `camellia` plugin settings for an `exclude` list under `linters.settings.custom.camellia.settings`.
- Allow users to exclude specific files, glob patterns, and directory trees from `camellia` diagnostics.
- Interpret exclude patterns relative to the nearest module root and normalize paths consistently before matching.
- Treat bare directory entries as recursive exclusions.
- Fail fast when an exclude pattern is malformed.

## Capabilities

### New Capabilities

- None.

### Modified Capabilities

- `camellia`: Extend linting requirements so configured files and directories are skipped by `camellia` while all other abbreviation-casing behavior remains unchanged.

## Impact

- Affected code: plugin settings decoding, analyzer construction, file-path filtering, and analyzer tests/fixtures.
- API/config surface: adds `linters.settings.custom.camellia.settings.exclude`.
- Dependencies: likely adds a glob-matching dependency to support `**` semantics cleanly.
- Validation: requires regression coverage for excluded files/directories and malformed patterns, plus end-to-end lint verification.
