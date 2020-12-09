# Batler

[![Build Status](https://travis-ci.com/vitorbaraujo/batler.svg?branch=master)](https://travis-ci.com/vitorbaraujo/batler)
[![codecov](https://codecov.io/gh/vitorbaraujo/batler/branch/master/graph/badge.svg?token=R4NPK8XCRW)](https://codecov.io/gh/vitorbaraujo/batler)

A Xcode test CLI for continuous integration. It aims to ease running tests and retrieving coverage reports for iOS applications.

## Installation

```sh
brew tap vitorbaraujo/formulae
brew install batler
```

## Usage

Before running the CLI, you first need to [configure](#configuration) the CLI.

### Run tests
```sh
batler test [-p <your_ios_project_path>] [-h]
```

### Displaying coverage

```sh
batler coverage [-p <your_ios_project_path>] [--html] [--output_dir <path>]
```

## Configuration

All the CLI configuration can be done through a config file named `.batler.yml`. This file needs to be located at your project root.

Here is an example configuration file:

```yaml
# File: .batler.yml
workspace: TestApp/MyProject.xcworkspace
project: TestApp/MyProject.xcodeproj
scheme: TestApp
build_dir: Build
xcode_version: 11.6 # will use '/Applications/Xcode-11.6.app/Contents/Developer'
# xcode_developer_dir: /my/custom/xcode/installation
clean: true
destination: platform=iOS Simulator,name=iPhone 8,OS=13.6
coverage:
  exclude:
  - TestApp/*
```

The config file contains the following parameters:

### workspace

The relative path to the xcworkspace directory.

### project

The relative path to the xcodeproj  directory.

### scheme

The scheme used to run tests.

### build_dir

The path where the test artifacts will be stored.

### xcode_version

The Xcode version to use when running tests. It expects the Xcode installation will be located at `/Applications/Xcode-{version}.app/Contents/Developer`

If you set this parameter, you cannot use `xcode_developer_dir`.

If you do not set neither `xcode_version` not `xcode_developer_dir`, the CLI will fetch the Xcode path from `xcode-select -p`.

### xcode_developer_dir

The Xcode developer directory path to run tests.

If you set this parameter, you cannot use `xcode_version`.

If you do not set neither `xcode_version` not `xcode_developer_dir`, the CLI will fetch the Xcode path from `xcode-select -p`.

### coverage

Set of parameters for configuring the coverage report.

#### coverage.exclude

A list of wildcard paths to exclude from the coverage report

### clean

If set to true, it will clean the project before building and testing when running `batler test`

### destination

A Xcode iOS simulator destination.

This parameter must be formatted as such: `platform=iOS Simulator,name={device_name},OS={os_version}`

If passed, the destination should be exist (check `xcrun simctl list devices` to check available simulators).

If you do not set this parameter, the CLI will create an arbitrary simulator given available runtimes and devicetypes (using `xcrun simctl`).

## Future improvements

- Add logging (+ verbose option)
- Add xcpretty integration
- Add xcodeproj option when running `batler test` (currently, the project config is only used for the `coverage` option)
- Add custom xcodebuild parameters when running `batler test`
- Add custom slather parameters when running `batler coverage`
- Add flag to pass custom config file path