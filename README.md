# Golangci-Lint Linter

[![ci](https://github.com/manuelarte/golangci-lint-linter/actions/workflows/ci.yml/badge.svg)](https://github.com/manuelarte/golangci-lint-linter/actions/workflows/ci.yml)
[![PR checks](https://github.com/manuelarte/golangci-lint-linter/actions/workflows/pr-checks.yml/badge.svg)](https://github.com/manuelarte/golangci-lint-linter/actions/workflows/pr-checks.yml)
![version](https://img.shields.io/github/v/release/manuelarte/golangci-lint-linter)

Linter for the configuration files for the well known linter [`golangci-lint`][golangci-lint].

## ⬇️ Getting Started

Not yet implemented, but the idea is to publish this as a cli application.

## 📜 Rules

The list of linting rules can be found in [RULES.md](./RULES.md).

## 🚀 How To Run It Locally

The project is using [uv][uv], to run it:

```bash
uv run golangci-lint-linter [FILE]
```

- `FILE`: the path to the `.golangci.yml` to be linter.

[golangci-lint]: https://golangci-lint.run/
[uv]: https://docs.astral.sh/uv/
