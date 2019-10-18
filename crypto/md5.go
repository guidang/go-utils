package crypto

import (
	"crypto/md5"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// NewMD5 NewMD5
func NewMD5() string {
	seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
	h := md5.New()
	h.Write([]byte(strconv.Itoa(seededRand.Int())))
	return hex.EncodeToString(h.Sum(nil))
}

// MD5 字符串生成 MD5
func MD5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}
