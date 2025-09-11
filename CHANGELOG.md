# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.1.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [1.0.2]

### Added
- `--trace` flag to create a `trace.out`

### Changed
- Internal refactor to reduce peak memory usage from 766MiB to 509MiB
- Explicitly state `EUPL-1.2 or later` as per EUPL Guideline 2021

## [1.0.1]

### Changed
- Internal refactor for maintainability and code-deduplication

## [1.0.0]

### Added
- `sources.yaml` for list url and header metadata
- Allowlist that is used to remove blocked domains from the final list
- `.update` additive diff files, so an existing `.list` file can be updated
- Configurable cleanup
- Support for domain, hosts and abp list types
- Rewrite mode to create a hosts file with custom IP

[1.0.2]: https://github.com/AtjonTV/monster-adlist-utility/compare/v1.0.1...v1.0.2
[1.0.1]: https://github.com/AtjonTV/monster-adlist-utility/compare/v1.0.0...v1.0.1
[1.0.0]: https://github.com/AtjonTV/monster-adlist-utility/releases/tag/v1.0.0
