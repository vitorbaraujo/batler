# Batler

[![Build Status](https://travis-ci.com/vitorbaraujo/batler.svg?branch=master)](https://travis-ci.com/vitorbaraujo/batler)
[![codecov](https://codecov.io/gh/vitorbaraujo/batler/branch/master/graph/badge.svg?token=R4NPK8XCRW)](https://codecov.io/gh/vitorbaraujo/batler)

A Xcode test CLI for continuous integration.

## Installation

```sh
brew tap vitorbaraujo/formulae
brew install batler
```

## Usage

### Run tests
```sh
batler test [-p <project_path>] [-h]
```

### Displaying coverage

```sh
batler coverage [--project_path <project_path>] [--html] [--output_dir <path>]
```

## TODO

- Add logging (+ verbose option)
- Add xcpretty integration
