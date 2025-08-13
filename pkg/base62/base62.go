package base62

import (
	"math"
	"strings"
)

// 避免恶意使用可以打乱

var (
	base62Str string
)

// MustInit 初始化base62
func MustInit(bs string) {
	if len(bs) == 0 {
		panic("need base string")
	}
	base62Str = bs
}

// Int2String 0->62进制string
func Int2String(seq uint64) string {
	if seq == 0 {
		return string(base62Str[0])
	}
	var result []byte
	for seq > 0 {
		result = append(result, base62Str[seq%62])
		seq /= 62
	}

	return string(reverse(result))
}

// String2Int 62string->10进制
func String2Int(s string) (seq uint64) {
	bl := []byte(s)
	bl = reverse(bl)
	for idx, b := range bl {
		base := math.Pow(62, float64(idx))
		seq += uint64(strings.Index(base62Str, string(b))) * uint64(base)
	}
	return seq
}

func reverse(s []byte) []byte {
	for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
		s[i], s[j] = s[j], s[i]
	}
	return s
}
