package rand

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/hex"
	"math/rand"
	"strconv"
	"time"
)

// RandomString 随机字符串
func RandomString(flag string, length int) string {
	var pool string

	switch flag {
	case "basic":
		seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		return strconv.Itoa(seededRand.Int())

	case "alnum", "numeric", "nozero", "alpha", "hex":
		switch flag {
		case "alpha":
			pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
			break
		case "alnum":
			pool = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
			break
		case "numeric":
			pool = "0123456789"
			break
		case "nozero":
			pool = "123456789"
			break
		case "hex":
			pool = "0123456789abcdefABCDEF"
			break
		}

		var result []byte
		bytes := []byte(pool)
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		for i := 0; i < length; i++ {
			result = append(result, bytes[r.Intn(len(bytes))])
		}
		return string(result)

	case "md5":
		seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		h := md5.New()
		h.Write([]byte(strconv.Itoa(seededRand.Int())))
		return hex.EncodeToString(h.Sum(nil))

	case "sha1":
		seededRand := rand.New(rand.NewSource(time.Now().UnixNano()))
		h := sha1.New()
		h.Write([]byte(strconv.Itoa(seededRand.Int())))
		return hex.EncodeToString(h.Sum(nil))
	}

	return ""
}
