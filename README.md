# suuid

Generate short, URL-safe UUIDs using base62 encoding.

## Getting Started

```bash
go get github.com/thani-sh/suuid-go
```

```go
import "github.com/thani-sh/suuid-go"

// Generate a compact base62-encoded UUID
id, err := suuid.V4()
// "1zzimYPK5krp9yc34RVmym"

// Convert the short UUID back to a standard UUID
uuid, err := suuid.Decode(id)
// "019b4a5a-fa57-778a-a1e0-cc25c5765935"
```