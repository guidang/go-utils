package hexconv

import "encoding/hex"

// String HexString To String
func String(str string) string {
	return hex.EncodeToString([]byte(str))
}
