# Agent Instructions

Project status: stable

## Tooling

- **MUST use the Go toolchain** (`go test`, `go vet`, `go mod tidy`) for build and test.
- **NEVER** introduce alternative package managers, linters, or formatters — `gofmt` is enough.

## Getting started

- **ALWAYS** run `go mod tidy` after pulling new changes.
- **NEVER** start a long-running `go run` or dev server without asking — most likely nothing is running.

## Version Control

- **NEVER** commit without explicit user approval after they have tested and verified.
- **STRICTLY** follow **Conventional Commits** for commit messages.
  - **ALWAYS** add a `type` and a short `description` to the commit message.
  - **NEVER** use scopes in commit messages.
  - **NEVER** use uppercase letters in commit messages.
  - **NEVER** add extra lines to the commit messages.
- **STRICTLY** follow branch naming strategy described below.
  - **ALWAYS** use hyphens `-`, do not use underscores `_` or camel case `camelCase`.
  - **NEVER** add any prefix to the branch name (e.g., `feat/` or `fix/`).
  - **NEVER** use uppercase letters in branch names.
