package rand

import (
	"math/rand"
	"time"
)

/*
RandomString 生成指定长度随机字符串
*/
func RandomString(t string, l int) string {
	pool := "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	switch t {

	case "alpha":
		pool = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	case "alnum":
		pool = "0123456789abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

	case "numeric":
		pool = "0123456789"

	case "nozero":
		pool = "123456789"

	case "hex":
		pool = "0123456789abcdefABCDEF"
	}

	bytes := []byte(pool)
	result := []byte{}

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	for i := 0; i < l; i++ {
		result = append(result, bytes[r.Intn(len(bytes))])
	}

	return string(result)
}
