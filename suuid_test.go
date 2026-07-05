package suuid

import (
	"regexp"
	"testing"

	"github.com/gofrs/uuid"
)

var (
	base62Pattern = regexp.MustCompile(`^[0-9A-Za-z]+$`)
	uuidPattern   = regexp.MustCompile(`^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$`)
)

// dnsNS is the standard DNS namespace UUID used by the V3/V5 tests.
// gofrs/uuid ships it as a pre-defined constant.
var dnsNS = uuid.NamespaceDNS

func TestV1(t *testing.T) {
	id, err := V1()
	if err != nil {
		t.Fatalf("V1 failed: %v", err)
	}
	if id == "" {
		t.Fatal("V1 returned empty string")
	}
	if len(id) >= 36 {
		t.Errorf("expected shorter than 36 chars, got %d (%q)", len(id), id)
	}
	if !base62Pattern.MatchString(id) {
		t.Errorf("invalid base62 string: %q", id)
	}

	id2, err := V1()
	if err != nil {
		t.Fatalf("V1 (second call) failed: %v", err)
	}
	if id == id2 {
		t.Error("expected different ids on successive calls")
	}
}

func TestV3(t *testing.T) {
	id, err := V3("hello", dnsNS)
	if err != nil {
		t.Fatalf("V3 failed: %v", err)
	}
	if id == "" {
		t.Fatal("V3 returned empty string")
	}
	if len(id) >= 36 {
		t.Errorf("expected shorter than 36 chars, got %d (%q)", len(id), id)
	}
	if !base62Pattern.MatchString(id) {
		t.Errorf("invalid base62 string: %q", id)
	}

	id2, err := V3("hello", dnsNS)
	if err != nil {
		t.Fatalf("V3 (second call) failed: %v", err)
	}
	if id != id2 {
		t.Errorf("expected identical ids for identical inputs, got %q and %q", id, id2)
	}
}

func TestV4(t *testing.T) {
	id, err := V4()
	if err != nil {
		t.Fatalf("V4 failed: %v", err)
	}
	if id == "" {
		t.Fatal("V4 returned empty string")
	}
	if len(id) >= 36 {
		t.Errorf("expected shorter than 36 chars, got %d (%q)", len(id), id)
	}
	if !base62Pattern.MatchString(id) {
		t.Errorf("invalid base62 string: %q", id)
	}

	id2, err := V4()
	if err != nil {
		t.Fatalf("V4 (second call) failed: %v", err)
	}
	if id == id2 {
		t.Error("expected different ids on successive calls")
	}
}

func TestV5(t *testing.T) {
	id, err := V5("hello", dnsNS)
	if err != nil {
		t.Fatalf("V5 failed: %v", err)
	}
	if id == "" {
		t.Fatal("V5 returned empty string")
	}
	if len(id) >= 36 {
		t.Errorf("expected shorter than 36 chars, got %d (%q)", len(id), id)
	}
	if !base62Pattern.MatchString(id) {
		t.Errorf("invalid base62 string: %q", id)
	}

	id2, err := V5("hello", dnsNS)
	if err != nil {
		t.Fatalf("V5 (second call) failed: %v", err)
	}
	if id != id2 {
		t.Errorf("expected identical ids for identical inputs, got %q and %q", id, id2)
	}
}

func TestV6(t *testing.T) {
	id, err := V6()
	if err != nil {
		t.Fatalf("V6 failed: %v", err)
	}
	if id == "" {
		t.Fatal("V6 returned empty string")
	}
	if len(id) >= 36 {
		t.Errorf("expected shorter than 36 chars, got %d (%q)", len(id), id)
	}
	if !base62Pattern.MatchString(id) {
		t.Errorf("invalid base62 string: %q", id)
	}

	id2, err := V6()
	if err != nil {
		t.Fatalf("V6 (second call) failed: %v", err)
	}
	if id == id2 {
		t.Error("expected different ids on successive calls")
	}
}

func TestV7(t *testing.T) {
	id, err := V7()
	if err != nil {
		t.Fatalf("V7 failed: %v", err)
	}
	if id == "" {
		t.Fatal("V7 returned empty string")
	}
	if len(id) >= 36 {
		t.Errorf("expected shorter than 36 chars, got %d (%q)", len(id), id)
	}
	if !base62Pattern.MatchString(id) {
		t.Errorf("invalid base62 string: %q", id)
	}

	id2, err := V7()
	if err != nil {
		t.Fatalf("V7 (second call) failed: %v", err)
	}
	if id == id2 {
		t.Error("expected different ids on successive calls")
	}
}

func TestEncode(t *testing.T) {
	tests := []string{
		"550e8400-e29b-41d4-a716-446655440000",
		"00000000-0000-0000-0000-000000000000",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
	}
	for _, u := range tests {
		encoded, err := Encode(u)
		if err != nil {
			t.Errorf("Encode(%q) failed: %v", u, err)
			continue
		}
		if encoded == "" {
			t.Errorf("Encode(%q) returned empty string", u)
		}
		if !base62Pattern.MatchString(encoded) {
			t.Errorf("Encode(%q) = %q, not a valid base62 string", u, encoded)
		}
		if len(encoded) >= len(u) {
			t.Errorf("Encode(%q) = %q, expected shorter than input", u, encoded)
		}
	}

	enc1, _ := Encode("00000000-0000-0000-0000-000000000000")
	enc2, _ := Encode("ffffffff-ffff-ffff-ffff-ffffffffffff")
	if enc1 == enc2 {
		t.Error("Encode produced identical output for distinct UUIDs")
	}
}

func TestEncodeInvalid(t *testing.T) {
	cases := []string{
		"not-a-uuid",
		"550e8400-e29b-41d4-a716-44665544000",   // too short
		"550e8400-e29b-41d4-a716-4466554400000", // too long
		"550e8400-e29b-41d4-a716-44665544000g",  // non-hex
		"",
	}
	for _, c := range cases {
		if _, err := Encode(c); err == nil {
			t.Errorf("Encode(%q) expected error, got nil", c)
		}
	}
}

func TestDecode(t *testing.T) {
	u := "550e8400-e29b-41d4-a716-446655440000"
	encoded, err := Encode(u)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	decoded, err := Decode(encoded)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if decoded != u {
		t.Errorf("round-trip mismatch: got %q, want %q", decoded, u)
	}
}

func TestDecodeRoundTrip(t *testing.T) {
	tests := []string{
		"00000000-0000-0000-0000-000000000000",
		"ffffffff-ffff-ffff-ffff-ffffffffffff",
		"123e4567-e89b-12d3-a456-426614174000",
		"019b4a5a-fa57-778a-a1e0-cc25c5765935",
	}
	for _, u := range tests {
		encoded, err := Encode(u)
		if err != nil {
			t.Errorf("Encode(%q) failed: %v", u, err)
			continue
		}
		decoded, err := Decode(encoded)
		if err != nil {
			t.Errorf("Decode(%q) failed: %v", encoded, err)
			continue
		}
		if decoded != u {
			t.Errorf("round-trip mismatch: got %q, want %q", decoded, u)
		}
	}
}

func TestDecodeInvalid(t *testing.T) {
	if _, err := Decode("invalid-chars!@#"); err == nil {
		t.Error("Decode expected error for invalid base62 chars, got nil")
	}
}

func TestIntegrationV1(t *testing.T) {
	su, err := V1()
	if err != nil {
		t.Fatalf("V1 failed: %v", err)
	}
	uu, err := Decode(su)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if !uuidPattern.MatchString(uu) {
		t.Errorf("decoded value %q is not a valid UUID", uu)
	}
	re, err := Encode(uu)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if re != su {
		t.Errorf("round-trip mismatch: got %q, want %q", re, su)
	}
}

func TestIntegrationV4(t *testing.T) {
	su, err := V4()
	if err != nil {
		t.Fatalf("V4 failed: %v", err)
	}
	uu, err := Decode(su)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if !uuidPattern.MatchString(uu) {
		t.Errorf("decoded value %q is not a valid UUID", uu)
	}
	re, err := Encode(uu)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if re != su {
		t.Errorf("round-trip mismatch: got %q, want %q", re, su)
	}
}

func TestIntegrationV7(t *testing.T) {
	su, err := V7()
	if err != nil {
		t.Fatalf("V7 failed: %v", err)
	}
	uu, err := Decode(su)
	if err != nil {
		t.Fatalf("Decode failed: %v", err)
	}
	if !uuidPattern.MatchString(uu) {
		t.Errorf("decoded value %q is not a valid UUID", uu)
	}
	re, err := Encode(uu)
	if err != nil {
		t.Fatalf("Encode failed: %v", err)
	}
	if re != su {
		t.Errorf("round-trip mismatch: got %q, want %q", re, su)
	}
}
