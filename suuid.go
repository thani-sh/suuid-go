package suuid

import (
	"fmt"

	"github.com/gofrs/uuid"
)

// V1 generates a new SUUID based on UUID v1 (timestamp and MAC address).
func V1() (string, error) {
	id, err := uuid.NewV1()
	if err != nil {
		return "", err
	}
	return encodeBase62(id.Bytes()), nil
}

// V3 generates a new SUUID based on UUID v3 (namespace and name, MD5).
func V3(name string, ns uuid.UUID) (string, error) {
	return encodeBase62(uuid.NewV3(ns, name).Bytes()), nil
}

// V4 generates a new SUUID based on UUID v4 (random).
func V4() (string, error) {
	id, err := uuid.NewV4()
	if err != nil {
		return "", err
	}
	return encodeBase62(id.Bytes()), nil
}

// V5 generates a new SUUID based on UUID v5 (namespace and name, SHA-1).
func V5(name string, ns uuid.UUID) (string, error) {
	return encodeBase62(uuid.NewV5(ns, name).Bytes()), nil
}

// V6 generates a new SUUID based on UUID v6 (timestamp and MAC address, sortable).
func V6() (string, error) {
	id, err := uuid.NewV6()
	if err != nil {
		return "", err
	}
	return encodeBase62(id.Bytes()), nil
}

// V7 generates a new SUUID based on UUID v7 (timestamp-based, sortable).
func V7() (string, error) {
	id, err := uuid.NewV7()
	if err != nil {
		return "", err
	}
	return encodeBase62(id.Bytes()), nil
}

// Encode encodes a standard UUID (8-4-4-4-12 format) to a SUUID (base62 encoded).
func Encode(s string) (string, error) {
	id, err := uuid.FromString(s)
	if err != nil {
		return "", fmt.Errorf("invalid UUID: %w", err)
	}
	return encodeBase62(id.Bytes()), nil
}

// Decode decodes a SUUID (base62 encoded) back to standard UUID format (8-4-4-4-12).
func Decode(s string) (string, error) {
	b, err := decodeBase62(s)
	if err != nil {
		return "", err
	}
	return formatUUID(b), nil
}
