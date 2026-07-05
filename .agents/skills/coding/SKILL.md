---
name: coding
description: Coding guidelines, naming standards and Go-style comments
---

# Coding Standards Guide

This skill provides the core coding standards, Go best practices, and code style guidelines to adhere to when developing features or refactoring code in the suuid-go workspace.

## Naming Conventions

To keep the codebase uniform, clean, and highly readable, use these conventions:

- **Variables & Functions:** `lowerCamelCase` (e.g., `encodeBase62`, `userID`).
- **Exported Types, Functions & Constants:** `PascalCase` (e.g., `Encode`, `NamespaceDNS`, `UUID`).
- **Unexported Constants:** `lowerCamelCase` (e.g., `uuidHexLength`) or `SCREAMING_SNAKE_CASE` for package-level invariants where appropriate.
- **Files & Directories:** `kebab-case` for multi-word names (e.g., `base62.go` is fine, prefer `user-store.go` over `userStore.go`).

---

## Go-style Comments

You **MUST** always add comments above exported functions, types, constants, and variables. Instead of verbose block JSDoc comments with annotations like `@param` or `@returns`, use concise **Go-style comments**. The comment must be a complete sentence starting with the name of the documented item. **Examples:**

```go
// Encode encodes a standard UUID (8-4-4-4-12 format) to a SUUID (base62 encoded).
func Encode(s string) (string, error) {
    // ...
}

// DNSNamespace is the standard DNS namespace UUID used for V3/V5 hashing.
var DNSNamespace = uuid.MustParse("6ba7b810-9dad-11d1-80b4-00c04fd430c8")
```

Go's `go doc` and `golangci-lint` (revive / gocritic) tooling rely on this format. A properly named identifier with a short, complete sentence is almost always better than a longer block with `@param`/`@returns`.

---

## Go Best Practices

- **Return errors explicitly.** No panics in library code paths. Use `(T, error)` returns and propagate with `fmt.Errorf("...: %w", err)`.
- **Accept interfaces, return structs** where the API surface warrants it. For a small library like this one, package-level functions are fine and idiomatic.
- **Do not shadow the standard library.** Names like `uuid`, `encoding`, `time` belong to stdlib — qualify them or pick different identifiers.
- **Run `gofmt` and `go vet` before considering a change done.** These are non-negotiable.