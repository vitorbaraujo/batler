# Changelog
All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

## [Unreleased]
### Added
- Add `coverage` command
  - It outputs the coverage after running `batler test`. For more information, run `batler coverage --help`.

## [0.2.0] - 2020-12-04
### Added
- Add `test` command
  - This replaces the behaviour when running `batler`. Now, in order to run tests,
    run `batler test`. For more information, run `batler test --help`.

## [0.1.0] - 2020-12-03
### Added
- Initial version. It contains:
  - Support for running `xcodebuild <clean|build|test>`
  - Some xcodebuild configurations: workspace, destination, scheme, derived data path and xcode version (or xcode hardcoded path)
  - Support for running tests using any available simulator


[Unreleased]: https://github.com/vitorbaraujo/batler/compare/0.2.0...HEAD
[0.2.0]: https://github.com/vitorbaraujo/batler/releases/tag/0.2.0
[0.1.0]: https://github.com/vitorbaraujo/batler/releases/tag/0.1.0
