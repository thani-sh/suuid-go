# Advanced Usage Guide

This document covers advanced usage, deep-dives into the library's capabilities, and serves as a comprehensive reference guide for `github.com/thani-sh/suuid-go`.

## Overview

`suuid-go` wraps [`github.com/gofrs/uuid`](https://github.com/gofrs/uuid) to generate short, URL-safe unique identifiers using base62 encoding. This results in significantly shorter strings (typically 22 characters) compared to the standard 36-character UUID format while maintaining the same uniqueness guarantees.

---

## Capabilities & Feature Walkthrough

### 1. Generating Short UUIDs (All Versions)

You can generate short UUIDs for UUID versions v1, v3, v4, v5, v6, and v7.

#### Random UUID (v4)

Generates a random short UUID.

```go
import "github.com/thani-sh/suuid-go"

id, _ := suuid.V4()
// "1zzimYPK5krp9yc34RVmym"
```

#### Timestamp-based UUID (v7 - Recommended)

Generates a time-ordered short UUID.

```go
id, _ := suuid.V7()
// "33kMoLhiPuX7auJxs62TZ"
```

#### Time & MAC address UUID (v1 & v6)

Generates a short UUID based on timestamp and MAC address.

- `V1()`: Standard layout.
- `V6()`: Field-compatible layout for sequential sorting.

```go
id1, _ := suuid.V1()
// "4xbLWrQUTRXaDGGqKItUuM"
id6, _ := suuid.V6()
// "wfWFsDJlkHuKkFelTJvDe"
```

#### Name-based UUIDs (v3 & v5)

Generates deterministic short UUIDs based on a namespace and a name.

- `V3()`: Uses MD5 hashing.
- `V5()`: Uses SHA-1 hashing.

```go
import "github.com/gofrs/uuid"
import "github.com/thani-sh/suuid-go"

ns := uuid.NamespaceDNS // 6ba7b810-9dad-11d1-80b4-00c04fd430c8

id3, _ := suuid.V3("hello", ns)
// "M1wYckVjHtf9kNJADE0zD"
id5, _ := suuid.V5("hello", ns)
// "4TsPdKNpquz4a4J5Fr9yHB"
```

The namespace parameter is a `uuid.UUID` (from `gofrs/uuid`). The package exposes `NamespaceDNS`, `NamespaceURL`, `NamespaceOID`, and `NamespaceX500` as ready-to-use values, or you can parse your own with `uuid.FromString`.

---

### 2. Bidirectional Conversion

You can encode existing standard UUIDs into base62 short UUIDs and decode them back.

#### Encoding UUIDs

Convert a standard 36-character UUID to a compact base62 string:

```go
suuid, _ := suuid.Encode("019b4a5a-fa57-778a-a1e0-cc25c5765935")
// "31xXF9ob9Zc8lbHAPvq8j"
```

#### Decoding SUUIDs

Convert a base62 short UUID back to the standard 36-character UUID format:

```go
uuid, _ := suuid.Decode("31xXF9ob9Zc8lbHAPvq8j")
// "019b4a5a-fa57-778a-a1e0-cc25c5765935"
```

_Note: `Decode` returns an error if the input contains invalid base62 characters. `Encode` returns an error if the input is not a valid 8-4-4-4-12 UUID string._

---

### 3. Error Handling

All public functions return `(string, error)`. The error is non-nil only when:

- `V1`, `V4`, `V6`, `V7`: the underlying random source or hardware query fails (extremely rare in practice).
- `V3`, `V5`: should not fail under normal use — `uuid.NewV3` / `uuid.NewV5` cannot return errors.
- `Encode`: the input is not a valid 8-4-4-4-12 UUID string.
- `Decode`: the input contains characters outside the base62 alphabet.

---

## API Reference

### `V1()`

Generates a new SUUID based on UUID v1 (timestamp and MAC address).

- **Returns:** `(string, error)` — A base62-encoded short UUID.

### `V3(name string, ns uuid.UUID)`

Generates a new SUUID based on UUID v3 (namespace and name, MD5).

- **Parameters:**
  - `name: string` — The name to hash.
  - `ns: uuid.UUID` — The namespace UUID (e.g. `uuid.NamespaceDNS`).
- **Returns:** `(string, error)` — A base62-encoded short UUID.

### `V4()`

Generates a new SUUID based on UUID v4 (random).

- **Returns:** `(string, error)` — A base62-encoded short UUID.

### `V5(name string, ns uuid.UUID)`

Generates a new SUUID based on UUID v5 (namespace and name, SHA-1).

- **Parameters:**
  - `name: string` — The name to hash.
  - `ns: uuid.UUID` — The namespace UUID (e.g. `uuid.NamespaceDNS`).
- **Returns:** `(string, error)` — A base62-encoded short UUID.

### `V6()`

Generates a new SUUID based on UUID v6 (timestamp and MAC address, sortable).

- **Returns:** `(string, error)` — A base62-encoded short UUID.

### `V7()`

Generates a new SUUID based on UUID v7 (timestamp-based, sortable).

- **Returns:** `(string, error)` — A base62-encoded short UUID.

### `Encode(uuid string)`

Encodes a standard UUID (8-4-4-4-12 format) to a SUUID (base62 encoded).

- **Parameters:**
  - `uuid: string` — A UUID string in standard format.
- **Returns:** `(string, error)` — A base62-encoded short UUID.

### `Decode(suuid string)`

Decodes a SUUID (base62 encoded) back to standard UUID format (8-4-4-4-12).

- **Parameters:**
  - `suuid: string` — A base62-encoded short UUID.
- **Returns:** `(string, error)` — A standard UUID string. Returns an error if the input contains invalid base62 characters.

---

## Why SUUIDs?

Standard UUIDs are 36 characters long (including dashes), which can be unwieldy in URLs, databases, and space-constrained contexts. SUUIDs solve this by encoding UUIDs as base62 strings, resulting in 22-character identifiers while maintaining:

- **Same uniqueness guarantees** as standard UUIDs.
- **URL-safe characters** (no special encoding needed).
- **Reversible encoding** (can convert back to standard UUID format).
- **Compact representation** (approximately 39% shorter).