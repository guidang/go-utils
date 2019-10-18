package crypto

import (
	"crypto/hmac"
	"crypto/sha256"
)

// HmacSha256 HmacSha256
func HmacSha256(data, key []byte) []byte {
	s := hmac.New(sha256.New, key)
	s.Write(data)
	return s.Sum([]byte(nil))
}
