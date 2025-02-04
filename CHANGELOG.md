# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]

## [0.1.0] - 2023-10-27

### Added

-   Initial release of `strinterp` library.
-   Support for basic variable expansion (`${var}`).
-   Support for formatting (`${var:format}`) with basic numeric and string formats.
-   Escaping using `$$`.
-   Error handling for unclosed expansions, undefined variables, and invalid formats.