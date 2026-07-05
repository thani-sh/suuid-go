// Package suuid generates short, URL-safe UUIDs using base62 encoding.
package suuid

import (
	"errors"
	"fmt"
	"math/big"
	"strings"
)

const base62Alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

// uuidByteLength is the size in bytes of a UUID (128 bits / 8 bits per byte).
const uuidByteLength = 16

// uuidHexLength is the number of hex digits in a UUID (128 bits / 4 bits per hex char).
const uuidHexLength = 32

var base62Base = big.NewInt(62)

// encodeBase62 encodes a 16-byte big-endian UUID to a base62 string.
func encodeBase62(b []byte) string {
	num := new(big.Int).SetBytes(b)
	if num.Sign() == 0 {
		return "0"
	}

	digits := make([]byte, 0, 22)
	zero := big.NewInt(0)
	rem := new(big.Int)
	for num.Cmp(zero) > 0 {
		num.QuoRem(num, base62Base, rem)
		digits = append(digits, base62Alphabet[rem.Int64()])
	}
	// digits are produced least-significant-first; reverse them.
	for i, j := 0, len(digits)-1; i < j; i, j = i+1, j-1 {
		digits[i], digits[j] = digits[j], digits[i]
	}
	return string(digits)
}

// decodeBase62 decodes a base62 string back to a 16-byte big-endian UUID.
func decodeBase62(s string) ([]byte, error) {
	num := new(big.Int)
	for i := 0; i < len(s); i++ {
		c := s[i]
		idx := strings.IndexByte(base62Alphabet, c)
		if idx < 0 {
			return nil, fmt.Errorf("invalid base62 character: %q", c)
		}
		num.Mul(num, base62Base)
		num.Add(num, big.NewInt(int64(idx)))
	}

	b := num.Bytes()
	if len(b) > uuidByteLength {
		return nil, errors.New("base62 value exceeds UUID range")
	}
	if len(b) < uuidByteLength {
		padded := make([]byte, uuidByteLength)
		copy(padded[uuidByteLength-len(b):], b)
		b = padded
	}
	return b, nil
}

// formatUUID formats a 16-byte big-endian UUID as the standard 8-4-4-4-12 string.
func formatUUID(b []byte) string {
	hex := fmt.Sprintf("%032x", new(big.Int).SetBytes(b))
	return hex[0:8] + "-" + hex[8:12] + "-" + hex[12:16] + "-" + hex[16:20] + "-" + hex[20:32]
}
