package crypto

import (
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// NewSha1 NewSha1
func NewSha1() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	h := sha1.New()
	h.Write([]byte(strconv.Itoa(seededRand.Int())))
	return hex.EncodeToString(h.Sum(nil))
}
