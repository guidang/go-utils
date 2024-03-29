package conv

import (
	"fmt"
	"math"
	"strconv"
)

// Round Round
func Round(val float64, roundOn float64, places int) (newVal float64) {
	var round float64
	pow := math.Pow(10, float64(places))
	digit := pow * val
	_, div := math.Modf(digit)
	if div >= roundOn {
		round = math.Ceil(digit)
	} else {
		round = math.Floor(digit)
	}
	newVal = round / pow
	return
}

// ByteFormat bytes 数据流转换成 KB / MB / GB / TB 格式
func ByteFormat(size float64) string {
	suffixes := []string{
		"B",
		"KB",
		"MB",
		"GB",
		"TB",
	}

	base := math.Log(size) / math.Log(1024)
	getSize := Round(math.Pow(1024, base-math.Floor(base)), .5, 2)
	getSuffix := suffixes[int(math.Floor(base))]

	return fmt.Sprintf(strconv.FormatFloat(getSize, 'f', -1, 64) + " " + string(getSuffix))
}
